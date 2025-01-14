package Delete

import (
	db "apiGO/run/postgres"

	con "apiGO/run/constLog"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DeletedById(c *gin.Context) { //DeleteID
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

	selectId := fmt.Sprintf(`SELECT id FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		logger.Error(con.ErrNotConnect,
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		query := fmt.Sprintf(`DELETE FROM "Cars" WHERE "id" = %s`, id)
		res, err := database.Exec(query)
		if err != nil {
			logger.Error(con.ErrNotConnect,
				zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotConnect})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
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
