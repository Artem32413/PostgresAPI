package PutFu

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

func PutItem(c *gin.Context) { //Put
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Furnitures" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}

	var updateRequest v.Furniture
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		logger.Error(con.ErrInvalidRequest,
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}
	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Furnitures" SET "Name" = '%s' , "Manufacturer" = '%s', "Height" = '%d', "Width" = '%d', "Length" = '%d' WHERE "id" = %s`, updateRequest.Name, updateRequest.Manufacturer, updateRequest.Height, updateRequest.Width, updateRequest.Length, id)
		_, err := database.Exec(param)
		if err != nil {
			logger.Error(con.ErrNotConnect,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
		c.IndentedJSON(http.StatusOK, updateRequest)
	} else {
		logger.Error(con.ErrInternal)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
