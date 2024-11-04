package main

import (
	"os"
	"todoapp/internal/handler"
	"todoapp/internal/repository"
	todo "todoapp/internal/server"
	"todoapp/internal/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Не удалось загрузить файл env", err)
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка инициализации конфига: %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Ошибка подлкючения к базе данных: %s", err)
	}

	rep := repository.NewRepository(db)      // слой работы с бд
	services := service.NewService(rep)      // слой бизнес логики, зависит от реп
	handlers := handler.NewHandler(services) // слой роутов, зависит от сервиса

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Ошибка при запуске сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
