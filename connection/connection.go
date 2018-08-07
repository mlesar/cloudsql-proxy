package connection

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	goauth "golang.org/x/oauth2/google"
)

const sqlScope = "https://www.googleapis.com/auth/sqlservice.admin"

func Open(ctx context.Context, driver, connStr string, jsonKey []byte) (*sql.DB, error) {
	if connStr == "" {
		return nil, errors.New("connection string is not set")
	}

	if jsonKey != nil {
		if err := initProxy(ctx, jsonKey); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open(driver, connStr)

	if err == nil {
		err = db.PingContext(ctx)
	}

	return db, err
}

func initProxy(ctx context.Context, jsonKey []byte) error {
	cfg, err := goauth.JWTConfigFromJSON(jsonKey, sqlScope)
	if err != nil {
		return err
	}

	client := cfg.Client(ctx)

	if client != nil {
		// Initializing the proxy
		proxy.Init(client, nil, nil)
	}

	return nil
}
