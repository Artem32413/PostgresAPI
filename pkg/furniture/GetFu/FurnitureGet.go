package GetFu

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFurnitures(c *gin.Context) { //Get
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	slFurniture := []v.Furniture{}
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	res, err := database.Query(`SELECT * FROM "Furnitures"`)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	for res.Next() {
		strFurniture := v.Furniture{}
		err = res.Scan(&strFurniture.ID, &strFurniture.Name, &strFurniture.Manufacturer, &strFurniture.Height, &strFurniture.Width, &strFurniture.Length)
		if err != nil {
			logger.Error(con.ErrInternal,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFurniture = append(slFurniture, strFurniture)
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	c.IndentedJSON(http.StatusOK, slFurniture)
}
