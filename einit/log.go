package einit

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os"
)

func initLog() {
	conf := config.Log
	log.SetLevel(log.Level(conf.Level))
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	tf := &log.TextFormatter{}
	tf.FieldMap = log.FieldMap{
		log.FieldKeyTime:  "@times",
		log.FieldKeyLevel: "@level",
		log.FieldKeyMsg:   "@message",
		log.FieldKeyFunc:  "@caller",
		log.FieldKeyFile:  "@file",
	}
	log.SetFormatter(&format{tf: tf})
}

type format struct {
	tf *log.TextFormatter
}

func (f *format) Format(entry *log.Entry) ([]byte, error) {
	serialized, err := f.tf.Format(entry)
	if entry.Context != nil {
		if requestId, ok := entry.Context.Value("request_id").(string); ok {
			var b *bytes.Buffer
			b = &bytes.Buffer{}

			b.WriteString("[")
			b.WriteString(requestId)
			b.WriteString("]")

			if entry.Buffer != nil {
				b.Write(entry.Buffer.Bytes())
			}
			return b.Bytes(), nil
		}
	}

	return serialized, err
}
