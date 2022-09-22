package server

import (
	"context"
	"errors"
	"fmt"
	userController "go-rest-api/src/controller"
	env "go-rest-api/src/functions"
	userService "go-rest-api/src/services/users"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	router      *gin.Engine
	ctx         context.Context
	cancel      context.CancelFunc
	err         error
	srv         *http.Server
	quit        chan os.Signal
	mongoClient *mongo.Client
	userDb      *mongo.Collection
)

func Server() {

	uri := env.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	db := env.GetEnv("DB", "go_test")

	mongoconn := options.Client().ApplyURI(uri)
	mongoClient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("Error while connecting with mongo", err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while trying to ping mongo", err)
	}

	fmt.Println("MongoDB connection established successfully.")

	userDb = mongoClient.Database(db).Collection("Users")
	userService := userService.InitUserService(userDb, ctx)
	userController := userController.InitUserController(userService)

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	router = gin.Default()

	srv = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	basepath := router.Group("/v1")
	userController.RegisterUserRoutes(basepath)

	quit = make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
