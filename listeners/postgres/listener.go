package postgres

import (
	"time"

	"github.com/lib/pq"
	"github.com/mlesar/cloudsql-proxy/dialers/postgres"
)

func New(driver, name string, minReconnectInterval, maxReconnectInterval time.Duration, eventCallback pq.EventCallbackType) *pq.Listener {
	if driver == "cloudsqlpostgres" {
		return pq.NewDialListener(postgres.Dialer{}, name, minReconnectInterval, maxReconnectInterval, eventCallback)
	}
	return pq.NewListener(name, minReconnectInterval, maxReconnectInterval, eventCallback)
}
