package router

import (
	fuDelete "apiGO/pkg/car/Delete"
	fuGet "apiGO/pkg/car/Get"
	fuGetID "apiGO/pkg/car/GetID"
	fuPatch "apiGO/pkg/car/Patch"
	fuPost "apiGO/pkg/car/Post"
	fuPut "apiGO/pkg/car/Put"
	
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	fuDelete2 "apiGO/pkg/flower/DeleteF"
	fuGet2 "apiGO/pkg/flower/GetF"
	fuGetID2 "apiGO/pkg/flower/GetIDF"
	fuPatch2 "apiGO/pkg/flower/PatchF"
	fuPost2 "apiGO/pkg/flower/PostF"
	fuPut2 "apiGO/pkg/flower/PutF"

	fuDelete3 "apiGO/pkg/furniture/DeleteFu"
	fuGet3 "apiGO/pkg/furniture/GetFu"
	fuGetID3 "apiGO/pkg/furniture/GetIDFu"
	fuPatch3 "apiGO/pkg/furniture/PatchFu"
	fuPost3 "apiGO/pkg/furniture/PostFu"
	fuPut3 "apiGO/pkg/furniture/PutFu"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() {
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	router := gin.Default()
	//flowers
	router.GET("/flowers", fuGet2.GetFlowers)
	router.GET("/flowers/:id", fuGetID2.GetFlowersByID)
	router.DELETE("/flowers/:id", fuDelete2.DeletedById)
	router.POST("/flowers", fuPost2.PostFlowers)
	router.PUT("/flowers/:id", fuPut2.PutItem)
	router.PATCH("/flowers/:id", fuPatch2.PatchItem)
	//cars
	router.GET("/cars", fuGet.GetCars)
	router.GET("/cars/:id", fuGetID.GetCarsByID)
	router.DELETE("/cars/:id", fuDelete.DeletedById)
	router.POST("/cars", fuPost.PostCars)
	router.PUT("/cars/:id", fuPut.PutItem)
	router.PATCH("/cars/:id", fuPatch.PatchItem)
	//furniture
	router.GET("/furniture", fuGet3.GetFurnitures)
	router.GET("/furniture/:id", fuGetID3.GetFurnituresByID)
	router.DELETE("/furniture/:id", fuDelete3.DeletedById)
	router.POST("/furniture", fuPost3.PostFurnitures)
	router.PUT("/furniture/:id", fuPut3.PutItem)
	router.PATCH("/furniture/:id", fuPatch3.PatchItem)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("Сервер запущен на порту 8080",
			zap.Duration("Продолжительность выполнения", time.Second),
		)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Ошибка сервера: %s", zap.Error(err))
		}
	}()

	<-stop
	logger.Info("Ожидаем завершения работы...",
		zap.Duration("Продолжительность выполнения", time.Second),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Ошибка при завершении работы: %s", zap.Error(err))
	}
	logger.Info("Сервер завершил работу",
		zap.Duration("Продолжительность завершения", time.Second),
	)
}
