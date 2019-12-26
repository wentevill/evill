package einit

import (
	"os"
)

const (
	Log = iota << 1
	Mysql
	Redis
	Kafka
	Etcd
)

type initRes map[int]interface{}

func Init(n int, env string) (res initRes, err error) {
	configPath := os.Getenv(env)
	if configPath == "" {
		configPath = env
	}
	if err = configInit(configPath); err != nil {
		return
	}
	initLog()
	res = make(initRes, 0)
	for n != 0 {
		var mid interface{}
		var err error
		switch {
		case n&Mysql != 0:
			n &= ^Mysql
			mid, err = initMysql()
			res[Mysql] = mid
		default:
			err = ErrAssembly
		}
		if err != nil {
			return nil, err
		}
	}
	return
}
