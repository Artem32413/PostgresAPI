package Delete

import (
	db "apiGO/run/postgres"

	con "apiGO/run/constLog"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func DeletedById(c *gin.Context) { //DeleteID
	id := c.Param("id")
	database, err := db.Connect()
	if err != nil {
		log.Println(con.ErrDB, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrDB})
		return
	}
	defer database.Close()
	selectId := fmt.Sprintf(`SELECT id FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка id"})
		return
	}
	if res.Next() {
		query := fmt.Sprintf(`DELETE FROM "Cars" WHERE "id" = %s`, id)
		res, err := database.Exec(query)
		if err != nil {
			log.Println("Ошибка id данных:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": con.ErrNotFound})
	}

}
