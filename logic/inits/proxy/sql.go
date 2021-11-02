package proxy

import (
	"github.com/lfxnxf/frame/BackendPlatform/golang/sql"
	"github.com/lfxnxf/frame/logic/inits"
	"golang.org/x/net/context"
)

type SQL struct {
	name []string
}

func InitSQL(name ...string) *SQL {
	if len(name) == 0 {
		return nil
	}
	return &SQL{name}
}

func (s *SQL) Master(ctx context.Context, name ...string) *sql.Client {
	var gName string
	if len(name) == 0 {
		gName = s.name[0]
	} else {
		gName = name[0]
	}
	return inits.SQLClient(ctx, gName).Master()
}

func (s *SQL) Slave(ctx context.Context, name ...string) *sql.Client {
	var gName string
	if len(name) == 0 {
		gName = s.name[0]
	} else {
		gName = name[0]
	}
	return inits.SQLClient(ctx, gName).Slave()
}
