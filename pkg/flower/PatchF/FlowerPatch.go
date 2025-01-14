package PatchF

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
	var outstruct v.Flower
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Name, &outstruct.Quantity, &outstruct.Price, &outstruct.ArrivalDate)
		if err != nil {
			logger.Error(con.ErrNotConnect,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
	} else {
		logger.Error(con.ErrInternal)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}

	var instruct v.Flower
	if err := c.ShouldBindJSON(&instruct); err != nil {
		logger.Error(con.ErrInternal,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if instruct.Name != "" {
		outstruct.Name = instruct.Name
	}
	if instruct.Quantity != 0 {
		outstruct.Quantity = instruct.Quantity
	}
	if instruct.Price != 0 {
		outstruct.Price = instruct.Price
	}
	if instruct.ArrivalDate != "" {
		outstruct.ArrivalDate = instruct.ArrivalDate
	}
	param := fmt.Sprintf(`UPDATE "Flowers" SET "Name" = '%s' , "Quantity" = '%d', "Price" = '%f', "ArrivalDate" = '%s' WHERE "id" = %s`, outstruct.Name, outstruct.Quantity, outstruct.Price, outstruct.ArrivalDate, id)
	_, err = database.Exec(param)
	if err != nil {
		logger.Error(con.ErrInternal,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}

	c.IndentedJSON(http.StatusOK, outstruct)
}
