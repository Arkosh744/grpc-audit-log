package main

import (
	"context"
	"fmt"
	"github.com/Arkosh744/grpc-audit-log/internal/config"
	"github.com/Arkosh744/grpc-audit-log/internal/repository"
	"github.com/Arkosh744/grpc-audit-log/internal/server"
	"github.com/Arkosh744/grpc-audit-log/internal/service"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New("internal/config")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DBUsername,
		Password: cfg.DBPassword,
	})

	opts.ApplyURI(cfg.DBURI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DBDatabase)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
