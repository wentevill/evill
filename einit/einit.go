package einit

import (
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	Log = iota << 1
	Mysql
	Redis
	Kafka
	Etcd
)

var assemblyEnum = map[int]string{
	Log:   "log",
	Mysql: "mysql",
	Redis: "redis",
	Kafka: "kafka",
	Etcd:  "etcd",
}

type assembly map[int]interface{}

func Init(n int, env string) (res assembly, err error) {
	configPath := os.Getenv(env)
	if configPath == "" {
		configPath = env
	}
	if err = configInit(configPath); err != nil {
		return
	}
	initLog()
	log.Infof("init %s success...", assemblyEnum[Log])
	res = make(assembly, 0)
	for n != 0 {
		var mid interface{}
		var err error
		var done int
		switch {
		case n&Mysql != 0:
			n &= ^Mysql
			done = Mysql
			mid, err = initMysql()
			res[Mysql] = mid
		default:
			err = ErrAssembly
		}
		if err != nil {
			return nil, err
		}
		log.Infof("init %s success...", assemblyEnum[done])
	}
	return
}
