package GetIDFu

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFurnituresByID(c *gin.Context) { //GetID
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	slFurnitures := []v.Furniture{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	query := fmt.Sprintf(`SELECT * FROM "Furnitures" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		strFurniture := v.Furniture{}

		err = res.Scan(&strFurniture.ID, &strFurniture.Name, &strFurniture.Manufacturer, &strFurniture.Height, &strFurniture.Width, &strFurniture.Length)
		if err != nil {
			logger.Error(con.ErrInternal,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFurnitures = append(slFurnitures, strFurniture)
		c.IndentedJSON(http.StatusOK, slFurnitures)
	} else {
		logger.Error(con.ErrNotFound)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
