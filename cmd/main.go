package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/m/v2/configs"
	"example.com/m/v2/internal/framework"
	"example.com/m/v2/internal/http/handler"
	"example.com/m/v2/internal/repository/exercises_repo"
	"example.com/m/v2/internal/repository/postgres"
	"example.com/m/v2/internal/repository/workout_instructions_repo"
	"example.com/m/v2/internal/services/exercises_service"
	"example.com/m/v2/internal/services/workout_instructions_service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg, err := configs.DNSFromEnv()
	if err != nil {
		panic("panica")
	}
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	workoutInstructionsRepo := workout_instructions_repo.NewRepo(db)
	workoutExercisesRepo := exercises_repo.NewRepo(db)

	txRunner := framework.NewTxRunner(db)

	workoutInstructionsService := workout_instructions_service.NewService(workoutInstructionsRepo, workoutExercisesRepo, *txRunner)
	workoutExercisesService := exercises_service.NewService(workoutExercisesRepo)

	handlers := handler.NewHandler(workoutInstructionsService, workoutExercisesService)

	router := handlers.InitRoutes()
	httpSrv := &http.Server{Addr: ":" + viper.GetString("port"), Handler: router}
	go func() {
		if err = httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err = httpSrv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
