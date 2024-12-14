package putf
import (
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

	var updateRequest v.Flower
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}
	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Flowers" SET "Name" = '%s' , "Quantity" = '%d', "Price" = '%d', "ArrivalDate" = '%s' WHERE "id" = %s`, updateRequest.Name, updateRequest.Quantity, updateRequest.Price, updateRequest.ArrivalDate, id)
		_, err := database.Exec(param)
		if err != nil {
			log.Println("Ошибка id данных:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
			return
		}
		c.IndentedJSON(http.StatusOK, updateRequest)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
	}
	defer database.Close()
}