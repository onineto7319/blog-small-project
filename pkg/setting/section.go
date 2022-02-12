package setting

import (
	"fmt"
	"time"
)

type ServerSettings struct {
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseMysqlSettings struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (s *Setting) ReadSection(k string, v interface{}) error {

	fmt.Println(k, v)
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
