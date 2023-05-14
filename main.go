package main

import (
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/server"
	pkg_postgres "github.com/engineerXIII/Diploma-server/pkg/db/postgres"
	"github.com/engineerXIII/Diploma-server/pkg/db/redis"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	cfgFile, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	var appLogger = logger.NewServerLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	appLogger.Info("Initating DB connection")
	db, err := pkg_postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	appLogger.Info("Starting migration")
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/postgres",
		"diploma", driver)
	if err != nil {
		appLogger.Fatalf("Migration error: %s", err)
	}
	err = m.Up()
	if err != nil {
		appLogger.Infof("Migration state: %s", err)
		appLogger.Info("Migration completed")
	}

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	s := server.NewServer(cfg, db, redisClient, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
