package db

import (
	"context"
	"log"
	"testing"

	"github.com/mattermost/mattermost-plugin-calls/server/testutils"

	serverMocks "github.com/mattermost/mattermost-plugin-calls/server/mocks/github.com/mattermost/mattermost-plugin-calls/server/interfaces"
	mlogMocks "github.com/mattermost/mattermost-plugin-calls/server/mocks/github.com/mattermost/mattermost/server/public/shared/mlog"

	"github.com/mattermost/mattermost/server/public/model"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func initMMSchema(t *testing.T, store *Store) {
	t.Helper()

	if store.driverName == model.DatabaseDriverPostgres {
		_, err := store.wDB.Exec(`
CREATE TABLE IF NOT EXISTS pluginkeyvaluestore (
    pluginid varchar(190) NOT NULL,
    pkey varchar(150) NOT NULL,
    pvalue bytea,
		expireat bigint DEFAULT 0,
    PRIMARY KEY (pluginid, pkey)
);
		`)
		require.NoError(t, err)
	} else {
		_, err := store.wDB.Exec(`
CREATE TABLE IF NOT EXISTS PluginKeyValueStore (
  PluginId varchar(190) NOT NULL,
  PKey varchar(150) NOT NULL,
  PValue mediumblob,
  ExpireAt bigint(20) DEFAULT 0,
  PRIMARY KEY (PluginId, PKey)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`)
		require.NoError(t, err)
	}
}

func newPostgresStore(t *testing.T) (*Store, func()) {
	t.Helper()

	mockMetrics := &serverMocks.MockMetrics{}
	mockLogger := &mlogMocks.MockLoggerIFace{}

	dsn, tearDown, err := testutils.RunPostgresContainerLocal(context.Background())
	require.NoError(t, err)

	var settings model.SqlSettings
	settings.SetDefaults(false)
	settings.DataSource = model.NewString(dsn)
	settings.DriverName = model.NewString(model.DatabaseDriverPostgres)

	conn, err := pq.NewConnector(dsn)
	require.NoError(t, err)
	require.NotNil(t, conn)

	mockLogger.On("Info", mock.Anything).Run(func(args mock.Arguments) {
		log.Printf(args.Get(0).(string))
	})
	mockLogger.On("Debug", mock.Anything).Run(func(args mock.Arguments) {
		log.Printf(args.Get(0).(string))
	})

	store, err := NewStore(settings, conn, nil, mockLogger, mockMetrics)
	require.NoError(t, err)
	require.NotNil(t, store)

	return store, func() {
		require.NoError(t, store.Close())
		tearDown()
	}
}

func newMySQLStore(t *testing.T) (*Store, func()) {
	t.Helper()

	mockMetrics := &serverMocks.MockMetrics{}
	mockLogger := &mlogMocks.MockLoggerIFace{}

	dsn, tearDown, err := testutils.RunMySQLContainerLocal(context.Background())
	require.NoError(t, err)

	var settings model.SqlSettings
	settings.SetDefaults(false)
	settings.DataSource = model.NewString(dsn)
	settings.DriverName = model.NewString(model.DatabaseDriverMysql)

	config, err := mysql.ParseDSN(dsn)
	require.NoError(t, err)

	conn, err := mysql.NewConnector(config)
	require.NoError(t, err)
	require.NotNil(t, conn)

	mockLogger.On("Info", mock.Anything).Run(func(args mock.Arguments) {
		log.Printf(args.Get(0).(string))
	})
	mockLogger.On("Debug", mock.Anything).Run(func(args mock.Arguments) {
		log.Printf(args.Get(0).(string))
	})

	store, err := NewStore(settings, conn, nil, mockLogger, mockMetrics)
	require.NoError(t, err)
	require.NotNil(t, store)

	return store, func() {
		require.NoError(t, store.Close())
		tearDown()
	}
}