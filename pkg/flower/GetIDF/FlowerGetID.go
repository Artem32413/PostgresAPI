package GetIDF

import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func GetFlowersByID(c *gin.Context) { //GetID
	slFlowers := []v.Flower{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	defer database.Close()
	query := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	if res.Next() {
		strFlowers := v.Flower{}

		err = res.Scan(&strFlowers.ID, &strFlowers.Name, &strFlowers.Quantity, &strFlowers.Price, &strFlowers.ArrivalDate)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
		slFlowers = append(slFlowers, strFlowers)
		c.IndentedJSON(http.StatusOK, slFlowers)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
	}

}
