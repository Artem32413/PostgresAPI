package Put

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
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	selectId := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}

	var updateRequest v.Car
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println(con.ErrInvalidRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}
	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = '%s' , "Model" = '%s', "Mileage" = '%d', "Owners" = '%d' WHERE "id" = %s`, updateRequest.Brand, updateRequest.Model, updateRequest.Mileage, updateRequest.Owners, id)
		_, err := database.Exec(param)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		c.IndentedJSON(http.StatusOK, updateRequest)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}

}
