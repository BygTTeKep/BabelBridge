package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"babelbridge/internal"
	"babelbridge/internal/company"
	"babelbridge/internal/config"
	"babelbridge/internal/database/repositories"
	kafkamanager "babelbridge/internal/kafkaManager"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	conf, err := config.LoadConfig("../../internal/config")
	if err != nil {
		panic(err)
	}
	pg := repositories.NewPG(conf)

	defer pg.DB.Close()

	kafka := &kafkamanager.KafkaConf{
		Host: "localhost",
		Port: 9092,
	}
	conn := kafka.ConnectKafka()

	router := gin.Default()
	kmRepo := kafkamanager.NewKafkaManagerRepositories(pg.DB)
	kafkaManagerServices := kafkamanager.NewKafkaManagerServices(kmRepo, conn)

	companyRepo := company.NewCompanyRepository(pg.DB)
	companyService := company.NewCompanyService(companyRepo, logger)

	services := internal.NewServices(kafkaManagerServices, companyService)
	newRoute := internal.NewRouters(router, *services)
	newRoute.Init()
	go func() {
		router.Run(fmt.Sprintf("%s:%s", conf.AppCfg.Host, conf.AppCfg.Port))
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
