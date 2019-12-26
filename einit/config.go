package einit

import (
	"errors"
	"strings"

	"evill/internal/file"
)

var (
	ErrAssembly = errors.New("unknown assembly")
)

var config = struct {
	Log struct {
		Level int
	}
	Mysql struct {
		DbName   string
		Addr     []string
		User     string
		Password string
		Port     int
		IDleConn int
		MaxConn  int
	}
}{}

var (
	ErrFileType = errors.New("unknown file suffix")
)

func configInit(path string) (err error) {
	var reader file.File
	switch {
	case strings.HasSuffix(path, "yml"):
		reader = new(file.YML)
	default:
		return ErrFileType
	}
	if err = reader.Load(&config, path); err != nil {
		return
	}
	return
}
