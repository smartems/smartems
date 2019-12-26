package sqlstore

import (
	"github.com/smartems/smartems/pkg/bus"
	m "github.com/smartems/smartems/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
