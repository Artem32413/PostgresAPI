package router

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"

	"go.uber.org/zap"
)

func TestingDb() {
	_, err := db.Connect()
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		return
	}
	logger.Info("Успешно")
}
