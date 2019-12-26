package einit

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	dbType = "mysql"
)

func initMysql() (*xorm.EngineGroup, error) {
	conf := config.Mysql
	engines := make([]*xorm.Engine, 0)
	for _, addr := range conf.Addr {
		dataSourceName := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", conf.User, conf.Password, addr, conf.DbName)
		if engine, err := xorm.NewEngine(dbType, dataSourceName); err == nil {
			engine.SetMaxIdleConns(conf.IDleConn)
			engine.SetMaxOpenConns(conf.IDleConn)
			engines = append(engines, engine)
		}
	}

	return xorm.NewEngineGroup(engines[0], engines[1:], xorm.RandomPolicy())
}

func (i initRes) Mysql() *xorm.EngineGroup {
	if m, ok := i[Mysql]; ok {
		return m.(*xorm.EngineGroup)
	}
	return nil
}
