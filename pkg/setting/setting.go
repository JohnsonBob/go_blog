package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize  int
	JwtSecret string

	RuntimeRootPath string
	PrefixUrl       string
	ImageSavePath   string

	ImageMaxSize   int64
	ImageAllowExts []string

	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	TimeFormat     string
	ExportSavePath string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Setting struct {
	App      *App
	Server   *Server
	Database *Database
	Redis    *Redis
}

var Config Setting

func Setup() {
	var err error
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	err = cfg.MapTo(&Config)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}

	Config.Server.WriteTimeout *= time.Second
	Config.Server.ReadTimeout *= time.Second
	Config.Redis.IdleTimeout *= time.Second
}
