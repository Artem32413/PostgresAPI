package PostF

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

func PostFlowers(c *gin.Context) { //Post
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
	var updateRequest v.Flower
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	param := fmt.Sprintf(`INSERT INTO "Flowers" ("Name", "Quantity", "Price", "Arrivaldate") VALUES ('%s', '%d', '%d', '%s') RETURNING id`, updateRequest.Name, updateRequest.Quantity, updateRequest.Price, updateRequest.ArrivalDate)
	res, err := database.Query(param)
	if err != nil {
		log.Println(con.ErrInternal, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
		return
	}
	if res.Next() {
		err = res.Scan(&updateRequest.ID)
		if err != nil {
			log.Println(con.ErrNotConnect, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, updateRequest)

}
