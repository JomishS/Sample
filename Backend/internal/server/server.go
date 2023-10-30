package server

import (
	"context"
	"example/Project3/database"
	"example/Project3/internal/config"
	"example/Project3/internal/inject"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/labstack/gommon/log"
)

func Init() error {
	var port string
	conf := config.GetConfig()
	err := database.InitPostgres(conf.Database.Postgres.ToPostgresConnectionConfig())
	if err != nil {
		return err
	}
	routeMapper := inject.GetRouteMapper()
	routeMapper.SetupRoutes()
	
	port=config.ViperConfig.GetString("SERVER_PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routeMapper.GetHandler(),
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		fmt.Println("inside go")
		log.Infof("Server starting on port:%s", port)

		err := srv.ListenAndServe()
		fmt.Println(" route mapping")
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error:%s\n", err)
		}

	}()

	<-ctx.Done()
	return nil
}

