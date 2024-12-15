package patch

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
	var outstruct v.Car
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	defer database.Close()
	selectId := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка id"})
		return
	}
	if res.Next() {

		err = res.Scan(&outstruct.ID, &outstruct.Brand, &outstruct.Model, &outstruct.Mileage, &outstruct.Owners)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
		return
	}

	var instruct v.Car
	if err := c.ShouldBindJSON(&instruct); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
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
	fmt.Println(outstruct)
	fmt.Println(instruct)
	param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = '%s' , "Model" = '%s', "Mileage" = '%d', "Owners" = '%d' WHERE "id" = %s`, outstruct.Brand, outstruct.Model, outstruct.Mileage, outstruct.Owners, id)
	_, err = database.Exec(param)
	if err != nil {
		log.Println("Ошибка id данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	c.IndentedJSON(http.StatusOK, outstruct)
	
}