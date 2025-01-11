package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	rou "apiGO/run/router/routers"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GracefulShotdown() {
	router := gin.Default()
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	rou.RunRouters(router)
	
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Ошибка при завершении работы: ", zap.Error(err))
		}
	}()

	logger.Info("Сервер запущен")
	logger.Info("Нажмите Ctrl+C для завершения")

	<-signalChan
	logger.Info("Ожидаем завершения работы...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Ошибка при завершении работы сервера: %v\n", err)
	}

	logger.Info("Сервер завершил работу")
}
