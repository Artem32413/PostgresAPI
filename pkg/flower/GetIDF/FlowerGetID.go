package GetIDF

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

func GetFlowersByID(c *gin.Context) { //GetID
	slFlowers := []v.Flower{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}
	query := fmt.Sprintf(`SELECT * FROM "Flowers" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		strFlowers := v.Flower{}

		err = res.Scan(&strFlowers.ID, &strFlowers.Name, &strFlowers.Quantity, &strFlowers.Price, &strFlowers.ArrivalDate)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slFlowers = append(slFlowers, strFlowers)
		c.IndentedJSON(http.StatusOK, slFlowers)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
