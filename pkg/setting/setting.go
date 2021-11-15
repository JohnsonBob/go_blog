package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

const (
	RUN_MODE      = "RUN_MODE"
	APP           = "app"
	PAGE_SIZE     = "PAGE_SIZE"
	JWT_SECRET    = "JWT_SECRET"
	SERVER        = "server"
	HTTP_PORT     = "HTTP_PORT"
	READ_TIMEOUT  = "READ_TIMEOUT"
	WRITE_TIMEOUT = "WRITE_TIMEOUT"
	DATABASE      = "database"
	TYPE          = "TYPE"
	USER          = "USER"
	PASSWORD      = "PASSWORD"
	HOST          = "HOST"
	NAME          = "NAME"
	TABLE_PREFIX  = "TABLE_PREFIX"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key(RUN_MODE).MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection(SERVER)
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key(HTTP_PORT).MustInt(8000)
	ReadTimeout = time.Duration(sec.Key(READ_TIMEOUT).MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key(WRITE_TIMEOUT).MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection(APP)
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key(JWT_SECRET).MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key(PAGE_SIZE).MustInt(10)
}
