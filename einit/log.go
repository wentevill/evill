package einit

import log "github.com/sirupsen/logrus"

func initLog() {
	conf := config.Log
	log.SetLevel(log.Level(conf.Level))
}
