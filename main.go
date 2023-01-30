package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gitlab.com/wiky.lyu/temgo/api"
	"gitlab.com/wiky.lyu/temgo/config"
	"gitlab.com/wiky.lyu/temgo/db"
	"gitlab.com/wiky.lyu/temgo/service/files"
	"gitlab.com/wiky.lyu/temgo/service/google"
	"gitlab.com/wiky.lyu/temgo/service/host"
)

const (
	APPName = "temgo"
)

func init() {
	if err := config.Init(APPName); err != nil {
		panic(err)
	}

	initLog()
	initDatabase()
	initFiles()
	initGoogle() // if you want to run this program without google play integration, just comment this line.
}

func initLog() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}

func initGoogle() {
	cfg := struct {
		Play struct {
			PackageName string `json:"packageName" yaml:"packageName"`
			JSONFile    string `json:"jsonFile" yaml:"jsonFile"`
		} `json:"play" yaml:"play"`
	}{}
	if err := config.Unmarshal("google", &cfg); err != nil {
		panic(err)
	}
	if err := google.Init(cfg.Play.PackageName, cfg.Play.JSONFile); err != nil {
		panic(err)
	}
}

func initDatabase() {

	cfg := struct {
		Debug bool `json:"debug" yaml:"debug"`
		PSQL  struct {
			DSN string `json:"dsn" yaml:"dsn"`
		}
	}{}
	if err := config.Unmarshal("db", &cfg); err != nil {
		panic(err)
	}

	if err := db.Init(cfg.PSQL.DSN, cfg.Debug); err != nil {
		panic(err)
	}
}

func initFiles() {
	cfg := struct {
		Path string `json:"path" yaml:"path"`
	}{}
	if err := config.Unmarshal("files", &cfg); err != nil {
		panic(err)
	}
	if err := files.Init(cfg.Path); err != nil {
		panic(err)
	}
}

func main() {

	cfg := struct {
		Host    string `json:"host" yaml:"host"`
		Listen  string `json:"listen" yaml:"listen"`
		Session string `json:"session" yaml:"session"`
		Debug   bool   `json:"debug" yaml:"debug"`
	}{}

	if err := config.Unmarshal("http", &cfg); err != nil {
		panic(err)
	}
	var listen string
	flag.StringVar(&listen, "http.listen", "", "bind address")
	flag.Parse()
	if listen != "" {
		cfg.Listen = listen
	}

	host.Init(cfg.Host)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:4200"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestedWith},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.Session))))
	e.HTTPErrorHandler = errorHandler

	api.Register(e)

	instanceName := fmt.Sprintf("%s %s", APPName, cfg.Listen)
	// initGraylog(instanceName)

	go func() {
		if err := e.Start(cfg.Listen); err != nil {
			log.Fatalf("shutting down the server: %v", err)
		}
	}()
	log.Infof("%s started", instanceName)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Warnf("%s quited", instanceName)
}

func errorHandler(err error, c echo.Context) {
	c.NoContent(500)
}
