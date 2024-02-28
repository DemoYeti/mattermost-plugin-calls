// Copyright (c) 2022-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-calls/server/db"
)

func (p *Plugin) initDB() error {
	serverCfg := p.API.GetConfig()
	if serverCfg == nil {
		return fmt.Errorf("server config should not be nil")
	}

	if serverCfg.SqlSettings.DriverName == nil {
		return fmt.Errorf("SqlSettings.DriverName should not be nil")
	}

	store, err := db.NewStore(*serverCfg.SqlSettings.DriverName, p.Driver, p.metrics)
	if err != nil {
		p.LogError(err.Error())
		return fmt.Errorf("failed to create db store: %w", err)
	}
	p.store = store

	return nil
}

// KVGet is an alternative to p.API.KVGet() that can also fetch from the writer DB node.
func (p *Plugin) KVGet(key string, fromWriter bool) ([]byte, error) {
	return p.store.KVGet(manifest.Id, key, fromWriter)
}
