package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/mattermost/mattermost-load-test-ng/loadtest/user/websocket"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pion/webrtc/v3/pkg/media/ivfreader"
	"github.com/pion/webrtc/v3/pkg/media/oggreader"
)

var (
	rtpAudioCodec = webrtc.RTPCodecCapability{
		MimeType:     "audio/opus",
		ClockRate:    48000,
		Channels:     2,
		SDPFmtpLine:  "minptime=10;useinbandfec=1",
		RTCPFeedback: nil,
	}
	rtpVideoCodecVP8 = webrtc.RTPCodecCapability{
		MimeType:    "video/VP8",
		ClockRate:   90000,
		Channels:    0,
		SDPFmtpLine: "",
		RTCPFeedback: []webrtc.RTCPFeedback{
			{Type: "goog-remb", Parameter: ""},
			{Type: "ccm", Parameter: "fir"},
			{Type: "nack", Parameter: ""},
			{Type: "nack", Parameter: "pli"},
		},
	}
)

type config struct {
	username      string
	password      string
	teamID        string
	channelID     string
	siteURL       string
	wsURL         string
	duration      time.Duration
	unmuted       bool
	screenSharing bool
}

func transmitScreen(ws *websocket.Client, pc *webrtc.PeerConnection, connectedCh <-chan struct{}) {
	track, err := webrtc.NewTrackLocalStaticSample(rtpVideoCodecVP8, "video", "screen-"+model.NewId())
	if err != nil {
		log.Fatalf(err.Error())
	}

	rtpSender, err := pc.AddTrack(track)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	go func() {
		// Open a IVF file and start reading using our IVFReader
		file, ivfErr := os.Open("./lt/samples/video.ivf")
		if ivfErr != nil {
			log.Fatalf(ivfErr.Error())
		}
		defer file.Close()

		ivf, header, ivfErr := ivfreader.NewWith(file)
		if ivfErr != nil {
			log.Fatalf(ivfErr.Error())
		}

		// Wait for connection established
		<-connectedCh

		info := map[string]string{
			"screenStreamID": track.StreamID(),
		}
		data, err := json.Marshal(&info)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err := ws.SendMessage("custom_com.mattermost.calls_screen_on", map[string]interface{}{
			"data": string(data),
		}); err != nil {
			log.Fatalf(err.Error())
		}
		defer func() {
			if err := ws.SendMessage("custom_com.mattermost.calls_screen_off", nil); err != nil {
				log.Fatalf(err.Error())
			}
		}()

		// Send our video file frame at a time. Pace our sending so we send it at the same speed it should be played back as.
		// This isn't required since the video is timestamped, but we will such much higher loss if we send all at once.
		//
		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		ticker := time.NewTicker(time.Millisecond * time.Duration((float32(header.TimebaseNumerator)/float32(header.TimebaseDenominator))*1000))
		for ; true; <-ticker.C {
			var frame []byte
			var ivfErr error
			frame, _, ivfErr = ivf.ParseNextFrame()
			if ivfErr == io.EOF || (ivfErr != nil && ivfErr.Error() == "incomplete frame data") {
				ivf.ResetReader(func(_ int64) io.Reader {
					_, _ = file.Seek(0, 0)
					ivf, header, ivfErr = ivfreader.NewWith(file)
					if ivfErr != nil {
						log.Fatalf(ivfErr.Error())
					}
					return file
				})
				frame, _, ivfErr = ivf.ParseNextFrame()
			}
			if ivfErr != nil {
				log.Fatalf(ivfErr.Error())
			}

			if err := track.WriteSample(media.Sample{Data: frame, Duration: time.Second}); err != nil {
				log.Printf("failed to write video sample: %s", err.Error())
			}
		}
	}()
}

func transmitAudio(ws *websocket.Client, pc *webrtc.PeerConnection, connectedCh <-chan struct{}) {
	track, err := webrtc.NewTrackLocalStaticSample(rtpAudioCodec, "audio", "voice"+model.NewId())
	if err != nil {
		log.Fatalf(err.Error())
	}

	rtpSender, err := pc.AddTrack(track)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	go func() {
		// Open a OGG file and start reading using our OGGReader
		file, oggErr := os.Open("./lt/samples/audio.ogg")
		if oggErr != nil {
			log.Fatalf(oggErr.Error())
		}
		defer file.Close()

		// Open on oggfile in non-checksum mode.
		ogg, _, oggErr := oggreader.NewWith(file)
		if oggErr != nil {
			log.Fatalf(oggErr.Error())
		}

		// Wait for connection established
		<-connectedCh

		if err := ws.SendMessage("custom_com.mattermost.calls_unmute", nil); err != nil {
			log.Fatalf(err.Error())
		}
		defer func() {
			if err := ws.SendMessage("custom_com.mattermost.calls_mute", nil); err != nil {
				log.Fatalf(err.Error())
			}
		}()

		// Keep track of last granule, the difference is the amount of samples in the buffer
		var lastGranule uint64

		// It is important to use a time.Ticker instead of time.Sleep because
		// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
		// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
		oggPageDuration := time.Millisecond * 20
		ticker := time.NewTicker(oggPageDuration)
		for ; true; <-ticker.C {
			var oggErr error
			var pageData []byte
			var pageHeader *oggreader.OggPageHeader
			pageData, pageHeader, oggErr = ogg.ParseNextPage()
			if oggErr == io.EOF {
				ogg.ResetReader(func(_ int64) io.Reader {
					_, _ = file.Seek(0, 0)
					return file
				})
				pageData, pageHeader, oggErr = ogg.ParseNextPage()
			}
			if oggErr != nil {
				log.Fatalf(oggErr.Error())
			}

			// The amount of samples is the difference between the last and current timestamp
			sampleCount := float64(pageHeader.GranulePosition - lastGranule)
			lastGranule = pageHeader.GranulePosition
			sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond

			if err := track.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); err != nil {
				log.Printf("failed to write audio sample: %s", err.Error())
			}
		}
	}()
}

func initRTC(ws *websocket.Client, channelID, username string, unmuted, screenSharing bool) (*webrtc.PeerConnection, error) {
	log.Printf("%s: setting up RTC connection", username)

	peerConnConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{
					"stun:calls.test.mattermost.com:3478",
				},
			},
		},
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}

	pc, err := webrtc.NewPeerConnection(peerConnConfig)
	if err != nil {
		return nil, err
	}

	pc.OnICECandidate(func(c *webrtc.ICECandidate) {
		log.Printf("ice: %v", c)
	})

	connectedCh := make(chan struct{})
	pc.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		if connectionState == webrtc.ICEConnectionStateConnected {
			close(connectedCh)
		}

		if connectionState == webrtc.ICEConnectionStateDisconnected || connectionState == webrtc.ICEConnectionStateFailed {
			log.Printf("ice disconnect")
			ws.Close()
		}
	})

	pc.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		rtcpSendErr := pc.WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}})
		if rtcpSendErr != nil {
			log.Printf(rtcpSendErr.Error())
		}

		codecName := strings.Split(track.Codec().RTPCodecCapability.MimeType, "/")[1]
		log.Printf("Track has started, of type %d: %s \n", track.PayloadType(), codecName)

		buf := make([]byte, 1400)
		for {
			_, _, readErr := track.Read(buf)
			if readErr != nil {
				log.Printf("%v", readErr.Error())
				return
			}
		}
	})

	if unmuted {
		transmitAudio(ws, pc, connectedCh)
	}

	if screenSharing {
		transmitScreen(ws, pc, connectedCh)
	}

	sdp, err := pc.CreateOffer(nil)
	if err != nil {
		return nil, err
	}

	if err := pc.SetLocalDescription(sdp); err != nil {
		return nil, err
	}

	var sdpData bytes.Buffer
	w := zlib.NewWriter(&sdpData)
	if err := json.NewEncoder(w).Encode(sdp); err != nil {
		return nil, err
	}
	w.Close()

	data := map[string]interface{}{
		"data": sdpData.Bytes(),
	}
	if err := ws.SendBinaryMessage("custom_com.mattermost.calls_sdp", data); err != nil {
		return nil, err
	}

	return pc, nil
}

func handleSignal(ws *websocket.Client, pc *webrtc.PeerConnection, ev *model.WebSocketEvent, iceCh chan webrtc.ICECandidateInit, username string) {
	evData := ev.GetData()
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(evData["data"].(string)), &data); err != nil {
		log.Fatalf(err.Error())
	}

	t, _ := data["type"].(string)

	if t == "candidate" {
		log.Printf("%s: ice!", username)
		iceCh <- webrtc.ICECandidateInit{Candidate: data["candidate"].(map[string]interface{})["candidate"].(string)}
	} else if t == "answer" {
		log.Printf("%s: sdp answer!", username)
		if err := pc.SetRemoteDescription(webrtc.SessionDescription{
			Type: webrtc.SDPTypeAnswer,
			SDP:  data["sdp"].(string),
		}); err != nil {
			log.Printf(err.Error())
		}

		go func() {
			for ice := range iceCh {
				if err := pc.AddICECandidate(ice); err != nil {
					log.Printf(err.Error())
				}
			}
		}()

	} else if t == "offer" {
		log.Printf("%s: sdp offer!", username)
		if err := pc.SetRemoteDescription(webrtc.SessionDescription{
			Type: webrtc.SDPTypeOffer,
			SDP:  data["sdp"].(string),
		}); err != nil {
			log.Printf(err.Error())
		}

		sdp, err := pc.CreateAnswer(nil)
		if err != nil {
			log.Printf(err.Error())
		}

		if err := pc.SetLocalDescription(sdp); err != nil {
			log.Printf(err.Error())
		}

		var sdpData bytes.Buffer
		w := zlib.NewWriter(&sdpData)
		if err := json.NewEncoder(w).Encode(sdp); err != nil {
			log.Fatalf(err.Error())
		}
		w.Close()

		data := map[string]interface{}{
			"data": sdpData.Bytes(),
		}
		if err := ws.SendBinaryMessage("custom_com.mattermost.calls_sdp", data); err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func eventHandler(ws *websocket.Client, channelID, username string, unmuted, screenSharing bool, doneCh chan struct{}) {
	var err error
	var pc *webrtc.PeerConnection
	iceCh := make(chan webrtc.ICECandidateInit, 10)
	defer close(iceCh)

	for {
		select {
		case ev, ok := <-ws.EventChannel:
			if !ok {
				return
			}
			switch ev.EventType() {
			case "hello":
				log.Printf("%s: joining call", username)
				data := map[string]interface{}{
					"channelID": channelID,
				}
				if err := ws.SendMessage("custom_com.mattermost.calls_join", data); err != nil {
					log.Fatalf(err.Error())
				}
			case "custom_com.mattermost.calls_join":
				log.Printf("%s: joined call", username)
				pc, err = initRTC(ws, channelID, username, unmuted, screenSharing)
				if err != nil {
					log.Fatalf(err.Error())
				}
				defer pc.Close()
			case "custom_com.mattermost.calls_signal":
				log.Printf("%s: received signal", username)
				handleSignal(ws, pc, ev, iceCh, username)
			default:
			}
		case <-doneCh:
			return
		}
	}
}

func connectUser(c config) error {
	log.Printf("%s: connecting user", c.username)

	var user *model.User
	client := model.NewAPIv4Client(c.siteURL)
	// login (or create) user
	user, _, err := client.Login(c.username, c.password)
	appErr, ok := err.(*model.AppError)
	if err != nil && !ok {
		return err
	}

	if ok && appErr != nil && appErr.Id != "api.user.login.invalid_credentials_email_username" {
		return err
	} else if ok && appErr != nil && appErr.Id == "api.user.login.invalid_credentials_email_username" {
		log.Printf("%s: registering user", c.username)
		user, _, err = client.CreateUser(&model.User{
			Username: c.username,
			Password: c.password,
			Email:    c.username + "@example.com",
		})
		if err != nil {
			return err
		}
		_, _, err = client.Login(c.username, c.password)
		if err != nil {
			return err
		}
	}

	log.Printf("%s: logged in", c.username)

	// join team
	_, _, err = client.AddTeamMember(c.teamID, user.Id)
	if err != nil {
		return err
	}

	// join channel
	_, _, err = client.AddChannelMember(c.channelID, user.Id)
	if err != nil {
		return err
	}

	log.Printf("%s: connecting to websocket", c.username)

	ws, err := websocket.NewClient4(&websocket.ClientParams{
		WsURL:     c.wsURL,
		AuthToken: client.AuthToken,
	})
	if err != nil {
		return err
	}

	doneCh := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		eventHandler(ws, c.channelID, c.username, c.unmuted, c.screenSharing, doneCh)
	}()

	ticker := time.NewTicker(c.duration)
	defer ticker.Stop()
	<-ticker.C

	log.Printf("%s: disconnecting...", c.username)
	close(doneCh)
	wg.Wait()
	ws.Close()

	log.Printf("%s: disconnected", c.username)

	return nil
}

func main() {
	// TODO: consider using a config file instead.
	var teamID string
	var siteURL string
	var userPassword string
	var userPrefix string
	var duration string
	var joinDuration string
	var adminUsername string
	var adminPassword string
	var offset int
	var numUnmuted int
	var numScreenSharing int
	var numCalls int
	var numUsersPerCall int

	flag.StringVar(&teamID, "team", "", "team ID")
	flag.StringVar(&siteURL, "url", "http://localhost:8065", "MM SiteURL")
	flag.StringVar(&userPrefix, "user-prefix", "testuser-", "user prefix")
	flag.StringVar(&userPassword, "user-password", "testPass123$", "user password")
	flag.IntVar(&numUnmuted, "unmuted", 0, "number of unmuted users per call")
	flag.IntVar(&numScreenSharing, "screen-sharing", 0, "number of users screen-sharing")
	flag.IntVar(&offset, "offset", 0, "users offset")
	flag.IntVar(&numCalls, "calls", 1, "number of calls")
	flag.IntVar(&numUsersPerCall, "users-per-call", 1, "number of users per call")
	flag.StringVar(&duration, "duration", "1m", "duration")
	flag.StringVar(&joinDuration, "join-duration", "30s", "join duration")
	flag.StringVar(&adminUsername, "admin-username", "sysadmin", "admin username")
	flag.StringVar(&adminPassword, "admin-password", "Sys@dmin-sample1", "admin password")

	flag.Parse()

	if teamID == "" {
		log.Fatalf("team must be set")
	}

	if numCalls == 0 {
		log.Fatalf("calls should be > 0")
	}

	if numUsersPerCall == 0 {
		log.Fatalf("users-per-call should be > 0")
	}

	if siteURL == "" {
		log.Fatalf("siteURL must be set")
	}

	dur, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalf(err.Error())
	}

	joinDur, err := time.ParseDuration(joinDuration)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var wsURL string
	u, err := url.Parse(siteURL)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if u.Scheme == "https" {
		wsURL = "wss://" + u.Host
	} else {
		wsURL = "ws://" + u.Host
	}

	if numUnmuted > numUsersPerCall {
		log.Fatalf("unmuted cannot be greater than the number of users per call")
	}

	if numScreenSharing > numCalls {
		log.Fatalf("screen-sharing cannot be greater than the number of calls")
	}

	adminClient := model.NewAPIv4Client(siteURL)
	_, _, err = adminClient.Login(adminUsername, adminPassword)
	if err != nil {
		log.Fatalf("failed to login as admin: %s", err.Error())
	}

	channels, _, err := adminClient.SearchChannels(teamID, &model.ChannelSearch{
		Public:  true,
		PerPage: &numCalls,
	})
	if err != nil {
		log.Fatalf("failed to search channels: %s", err.Error())
	}

	if len(channels) < numCalls {
		channels = make([]*model.Channel, numCalls)
		for i := 0; i < numCalls; i++ {
			name := model.NewId()
			channel, _, err := adminClient.CreateChannel(&model.Channel{
				TeamId:      teamID,
				Name:        name,
				DisplayName: "test-" + name,
				Type:        model.ChannelTypeOpen,
			})
			if err != nil {
				log.Fatalf("failed to create channel: %s", err.Error())
			}
			channels[i] = channel
		}
	}

	var wg sync.WaitGroup
	wg.Add(numUsersPerCall * numCalls)
	for j := 0; j < numCalls; j++ {
		log.Printf("starting call in %s", channels[j].DisplayName)
		for i := 0; i < numUsersPerCall; i++ {
			go func(idx int, channelID string, unmuted, screenSharing bool) {
				defer wg.Done()
				time.Sleep(time.Duration(rand.Intn(int(joinDur.Seconds()))) * time.Second)
				username := fmt.Sprintf("%s%d", userPrefix, idx)
				cfg := config{
					username:      username,
					password:      userPassword,
					teamID:        teamID,
					channelID:     channelID,
					siteURL:       siteURL,
					wsURL:         wsURL,
					duration:      dur,
					unmuted:       unmuted,
					screenSharing: screenSharing,
				}
				if err := connectUser(cfg); err != nil {
					log.Printf("connectUser failed: %s", err.Error())
				}
			}((numUsersPerCall*j)+i+offset, channels[j].Id, i < numUnmuted, i == 0 && j < numScreenSharing)
		}
	}

	wg.Wait()

	fmt.Println("DONE")
}
