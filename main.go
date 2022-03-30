package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	oauthLib "oauth2-server-go/internal/oauth/library"
	"oauth2-server-go/pkg/logr"
	"oauth2-server-go/pkg/valider"
	"oauth2-server-go/route"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

var port string

// Interrupt handler.
var errChan = make(chan error, 1)

//go:embed web/view/*
var fs embed.FS

func main() {
	// init http port
	flag.StringVar(&port, "port", "8080", "Initial port number")
	flag.Parse()

	// init config
	config.InitEnv()

	// init logger
	logr.InitLogger()

	// init validation
	valider.Init()

	// init driver
	_ = api.InitXorm()
	_ = api.InitRedis()
	_ = api.InitRedisCluster()

	// init oauth
	oauthLib.InitOauth2()

	// init gin router
	r := route.Init()

	// Start gin server
	go func() {
		tmpl := template.Must(template.New("").Funcs(template.FuncMap{"basepath": func() string {
			return config.GetHtmlBasePath()
		}}).ParseFS(fs, "web/view/*.tmpl"))
		r.SetHTMLTemplate(tmpl)
		r.StaticFS("/public", http.FS(fs))
		logr.L.Info("Auth-Backend server start", zap.String("port", port))
		errChan <- r.Run(":" + port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	<-errChan
}
