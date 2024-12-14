package getf
import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func GetFlowers(c *gin.Context) { //Get
	slFlower := []v.Flower{}
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	res, err := database.Query(`SELECT * FROM "Flowers"`)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	for res.Next() {
		strFlower := v.Flower{}
		err = res.Scan(&strFlower.ID, &strFlower.Name, &strFlower.Quantity, &strFlower.Price, &strFlower.ArrivalDate)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
		slFlower = append(slFlower, strFlower)
	}
	defer database.Close()
	c.IndentedJSON(http.StatusOK, slFlower)
}