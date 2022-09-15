package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"gitlab.com/wiky.lyu/temgo/config"
	"gitlab.com/wiky.lyu/temgo/db"
	"gitlab.com/wiky.lyu/temgo/db/temtem"
)

const (
	appName = "temgo"
)

func init() {
	if err := config.Init(appName); err != nil {
		panic(err)
	}
	initLog()
	InitDB()
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

func InitDB() {
	cfg := struct {
		PSQL struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"psql" yaml:"psql"`
		Debug bool `json:"debug" yaml:"debug"`
	}{}
	if err := config.Unmarshal("db", &cfg); err != nil {
		log.Fatalf("读取DB配置出错: %v", err)
	}
	if err := db.Init(cfg.PSQL.DSN, cfg.Debug); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
}

func main() {

	cfg := struct {
		Listen string `json:"listen" yaml:"listen"`
		Debug  bool   `json:"debug" yaml:"debug"`
	}{}
	if err := config.Unmarshal("http", &cfg); err != nil {
		log.Fatalf("读取HTTP配置出错:%v", err)
	}
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(
			reqlog.WithEnabled(cfg.Debug),
		)),
		bunrouter.WithNotFoundHandler(notFoundHandler),
		bunrouter.WithMethodNotAllowedHandler(methodNotAllowedHandler),
	)

	router.GET("/", indexHandler)

	router.WithGroup("/api", func(g *bunrouter.Group) {
		g.GET("/users/:id", debugHandler)
		g.GET("/users/current", debugHandler)
		g.GET("/users/*path", debugHandler)
	})

	httpServer := http.Server{
		Addr:    cfg.Listen,
		Handler: router,
	}
	go func() {
		log.Infof("监听 http://%s", cfg.Listen)
		if err := httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Errorf("ListenAndServe Error: %v", err)
			}
		}
	}()

	types := make([]*temtem.TemtemType, 0)
	if err := db.PG().NewSelect().Model(&types).Scan(context.Background()); err != nil {
		log.Errorf("DB Error: %v", err)
	}
	for _, t := range types {
		log.Infof("%v:  %s", t.Name, t.Trivia)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Errorf("%s 退出出错: %v", err)
	}
	log.Infof("%s 退出成功", appName)
}

func indexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"hello": "world",
	})
}

func debugHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
	})
}

func notFoundHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusNotFound)
	return nil
}

func methodNotAllowedHandler(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return nil
}
