package Post

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

func PostCars(c *gin.Context) { //Post
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
	var updateRequest v.Car
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println(con.ErrInvalidRequest, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": con.ErrInvalidData})
		return
	}

	fmt.Println(updateRequest)
	param := fmt.Sprintf(`INSERT INTO "Cars" ("Brand", "Model", "Mileage", "Owners") VALUES ('%s', '%s', '%d', '%d') RETURNING id`, updateRequest.Brand, updateRequest.Model, updateRequest.Mileage, updateRequest.Owners)
	res, err := database.Query(param)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		err = res.Scan(&updateRequest.ID)
		if err != nil {
			log.Println(con.ErrNotConnect, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, updateRequest)
}
