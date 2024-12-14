package patchf

import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func PatchItem(c *gin.Context) { //Patch
	var outstruct v.Flower
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	selectId := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка id"})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Name, &outstruct.Quantity, &outstruct.Price, &outstruct.ArrivalDate)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
		return
	}

	var instruct v.Flower
	if err := c.ShouldBindJSON(&instruct); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
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
	fmt.Println(outstruct)
	fmt.Println(instruct)
	param := fmt.Sprintf(`UPDATE "Flowers" SET "Name" = '%s' , "Quantity" = '%s', "Price" = '%d', "ArrivalDate" = '%d' WHERE "id" = %s`, outstruct.Name, outstruct.Quantity, outstruct.Price, outstruct.ArrivalDate, id)
	_, err = database.Exec(param)
	if err != nil {
		log.Println("Ошибка id данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	c.IndentedJSON(http.StatusOK, outstruct)
	defer database.Close()
}