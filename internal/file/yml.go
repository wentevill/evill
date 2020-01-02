package file

import (
	"gopkg.in/yaml.v2"
	"os"
)

type YML struct {
}

func (Y YML) Load(c interface{}, path string) (err error) {
	var f *os.File
	if f, err = os.OpenFile(path, os.O_RDONLY, 0666); err != nil {
		return
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(c)
}
