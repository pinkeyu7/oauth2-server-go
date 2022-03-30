package main

import (
	"flag"
	"log"
	"oauth2-server-go/api"
	"oauth2-server-go/config"
	oauth2lib "oauth2-server-go/internal/oauth2/library"
	"oauth2-server-go/pkg/logr"
	"oauth2-server-go/pkg/valider"
	"oauth2-server-go/route"

	_ "github.com/go-sql-driver/mysql"
)

var port string

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

	// init oauth2
	oauth2lib.InitOauth2()

	// init gin router
	r := route.Init()

	// start server
	err := r.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
