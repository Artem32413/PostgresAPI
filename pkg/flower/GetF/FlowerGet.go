package GetF

import (
	con "apiGO/run/constLog"
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFlowers(c *gin.Context) { //Get
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	slFlower := []v.Flower{}
	database, err := db.Connect()

	if err != nil {
		logger.Error("Ошибка подключения к базе данных:",
			zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
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

	logger.Info("Успешно")
	c.IndentedJSON(http.StatusOK, slFlower)
}
