package GetID

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

func GetCarsByID(c *gin.Context) { //GetID
	slCar := []v.Car{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}

	query := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		log.Println(con.ErrNotConnect, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
		return
	}
	if res.Next() {
		strCar := v.Car{}

		err = res.Scan(&strCar.ID, &strCar.Brand, &strCar.Model, &strCar.Mileage, &strCar.Owners)
		if err != nil {
			log.Println(con.ErrInternal, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrInternal})
			return
		}
		slCar = append(slCar, strCar)
		c.IndentedJSON(http.StatusOK, slCar)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}
	if err := database.Close(); err != nil {
		log.Println(con.ErrDBClose, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDBClose})
		return
	}
}
