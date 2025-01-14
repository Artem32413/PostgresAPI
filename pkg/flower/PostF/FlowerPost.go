package PostF

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

func PostFlowers(c *gin.Context) { //Post
	logger, _ := zap.NewDevelopment()
	if err := logger.Sync(); err != nil {
		zap.Error(err)
	}
	database, err := db.Connect()

	if err != nil {
		logger.Error(con.ErrDB,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	var updateRequest v.Flower
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		logger.Error(con.ErrInvalidRequest,
			zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}

	param := fmt.Sprintf(`INSERT INTO "Flowers" ("Name", "Quantity", "Price", "ArrivalDate") VALUES ('%s', '%d', '%f', '%s') RETURNING id`, updateRequest.Name, updateRequest.Quantity, updateRequest.Price, updateRequest.ArrivalDate)
	res, err := database.Query(param)
	if err != nil {
		logger.Error(con.ErrInternal,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if res.Next() {
		err = res.Scan(&updateRequest.ID)
		if err != nil {
			logger.Error(con.ErrNotConnect,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
	}
	if err := database.Close(); err != nil {
		logger.Error(con.ErrDBClose,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}

	c.IndentedJSON(http.StatusOK, updateRequest)
}
