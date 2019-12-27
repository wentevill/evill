package einit

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func initLog() {
	conf := config.Log
	log.SetLevel(log.Level(conf.Level))
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}
