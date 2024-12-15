package postfu

import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func PostFurnitures(c *gin.Context) { //Post
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	defer database.Close()
	var updateRequest v.Furniture
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}
	param := fmt.Sprintf(`INSERT INTO "Furnitures" ("Name" , "Manufacturer" , "Height" , "Width" , "Length") VALUES ('%s', '%s', '%d', '%d', '%d') RETURNING id`, updateRequest.Name, updateRequest.Manufacturer, updateRequest.Height, updateRequest.Width, updateRequest.Length)
	res, err := database.Query(param)
	if err != nil {
		log.Println("Ошибка id данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	if res.Next() {
		err = res.Scan(&updateRequest.ID)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, updateRequest)
	
}