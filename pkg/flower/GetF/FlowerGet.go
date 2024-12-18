package GetF

import (
	con "apiGO/run/constLog"
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
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}
	res, err := database.Query(`SELECT * FROM "Flowers"`)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	for res.Next() {
		strFlower := v.Flower{}
		err = res.Scan(&strFlower.ID, &strFlower.Name, &strFlower.Quantity, &strFlower.Price, &strFlower.ArrivalDate)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFlower = append(slFlower, strFlower)
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	c.IndentedJSON(http.StatusOK, slFlower)
}
