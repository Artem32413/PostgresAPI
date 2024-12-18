package PutF

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func PutItem(c *gin.Context) { //Put
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	selectId := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}

	var updateRequest v.Flower
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println(con.ErrInvalidRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}
	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Flowers" SET "Name" = '%s' , "Quantity" = '%d', "Price" = '%f', "ArrivalDate" = '%s' WHERE "id" = %s`, updateRequest.Name, updateRequest.Quantity, updateRequest.Price, updateRequest.ArrivalDate, id)
		_, err := database.Exec(param)
		if err != nil {
			log.Println(con.ErrNotConnect, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
		c.IndentedJSON(http.StatusOK, updateRequest)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
