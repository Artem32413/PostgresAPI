package Patch

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

func PatchItem(c *gin.Context) { //Patch
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	var outstruct v.Car
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		err = res.Scan(&outstruct.ID, &outstruct.Brand, &outstruct.Model, &outstruct.Mileage, &outstruct.Owners)
		if err != nil {
			logger.Error(con.ErrInternal,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
	} else {
		logger.Error(con.ErrNotFound)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}

	var instruct v.Car
	if err := c.ShouldBindJSON(&instruct); err != nil {
		logger.Error(con.ErrInternal,
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}
	if instruct.Brand != "" {
		outstruct.Brand = instruct.Brand
	}
	if instruct.Model != "" {
		outstruct.Model = instruct.Model
	}
	if instruct.Mileage != 0 {
		outstruct.Mileage = instruct.Mileage
	}
	if instruct.Owners != 0 {
		outstruct.Owners = instruct.Owners
	}
	param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = '%s' , "Model" = '%s', "Mileage" = '%d', "Owners" = '%d' WHERE "id" = %s`, outstruct.Brand, outstruct.Model, outstruct.Mileage, outstruct.Owners, id)
	_, err = database.Exec(param)
	if err != nil {
		logger.Error(con.ErrNotFound,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	c.IndentedJSON(http.StatusOK, outstruct)
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
