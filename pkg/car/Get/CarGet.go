package Get

import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) { //Get
	slCar := []v.Car{}
	database, err := db.Connect()
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	defer database.Close()
	res, err := database.Query(`SELECT * FROM "Cars"`)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	for res.Next() {
		strCar := v.Car{}
		err = res.Scan(&strCar.ID, &strCar.Brand, &strCar.Model, &strCar.Mileage, &strCar.Owners)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
		slCar = append(slCar, strCar)
	}
	c.IndentedJSON(http.StatusOK, slCar)
}
