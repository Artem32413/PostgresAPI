package Get

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetCars(c *gin.Context) { //Get
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	slCar := []v.Car{}
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	res, err := database.Query(`SELECT * FROM "Cars"`)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	for res.Next() {
		strCar := v.Car{}
		err = res.Scan(&strCar.ID, &strCar.Brand, &strCar.Model, &strCar.Mileage, &strCar.Owners)
		if err != nil {
			logger.Error(con.ErrInternal,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slCar = append(slCar, strCar)
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	c.IndentedJSON(http.StatusOK, slCar)
}
