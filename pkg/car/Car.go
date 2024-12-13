package car

import (
	db "apiGO/run/postgres"
	v "apiGO/structFile"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
		fmt.Println(slCar)
	}
	defer database.Close()
	c.IndentedJSON(http.StatusOK, slCar)
}
func GetCarsByID(c *gin.Context) { //GetID
	slCar := []v.Car{}
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	query := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(query)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	if res.Next() {
		strCar := v.Car{}

		err = res.Scan(&strCar.ID, &strCar.Brand, &strCar.Model, &strCar.Mileage, &strCar.Owners)
		if err != nil {
			log.Println("Ошибка чтения из БД:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка чтения из БД"})
			return
		}
		slCar = append(slCar, strCar)

		fmt.Println(strCar)
		c.IndentedJSON(http.StatusOK, slCar)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
	}
	defer database.Close()
}

func DeletedById(c *gin.Context) { //DeleteID
	id := c.Param("id")
	database, err := db.Connect()
	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
	}
	defer database.Close()
}
func PostCars(c *gin.Context) { //Post
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	selectId := fmt.Sprintf(`SELECT id FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка id"})
		return
	}
	selectBrand := fmt.Sprintf(`SELECT Brand FROM "Cars" WHERE "id" = %s`, id)
	selectModel := fmt.Sprintf(`SELECT Model FROM "Cars" WHERE "id" = %s`, id)
	selectMileage := fmt.Sprintf(`SELECT Mileage FROM "Cars" WHERE "id" = %s`, id)
	selectOwners := fmt.Sprintf(`SELECT Owners FROM "Cars" WHERE "id" = %s`, id)
	fmt.Println(selectOwners)

	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = %s , "Model" = %s, "Mileage" = %s, "Owners" = %s `, selectBrand, selectModel, selectMileage, selectOwners)
		res, err := database.Exec(param)
		if err != nil {
			log.Println("Ошибка id данных:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
			return
		}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "По такому id данные не найдены"})
	}
	defer database.Close()
}
func PutItem(c *gin.Context) { //Put
	id := c.Param("id")
	database, err := db.Connect()

	if err != nil {
		log.Println("Ошибка подключения к базе данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка подключения к базе данных"})
		return
	}
	selectId := fmt.Sprintf(`SELECT * FROM "Cars" WHERE "id" = %s`, id)
	res, err := database.Query(selectId)
	if err != nil {
		log.Println("Ошибка подключения данных:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка id"})
		return
	}

	var updateRequest v.Car
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}
	if res.Next() {
		param := fmt.Sprintf(`UPDATE "Cars" SET "Brand" = '%s' , "Model" = '%s', "Mileage" = '%d', "Owners" = '%d' WHERE "id" = %s`, updateRequest.Brand, updateRequest.Model, updateRequest.Mileage, updateRequest.Owners, id)
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
func PatchItem(c *gin.Context) { //Patch
	file, err := os.Open("file.json")
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Файл не найден"})
		return
	}
	defer file.Close()

	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка чтения файла:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении файла"})
		return
	}

	var items []v.Inventory
	if err := json.Unmarshal(readFile, &items); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при декодировании JSON"})
		return
	}

	carsID := c.Param("id")
	var carsToUpdate *v.Car
	for i := range items[0].Cars {
		if items[0].Cars[i].ID == carsID {
			carsToUpdate = &items[0].Cars[i]
			break
		}
	}

	var updateRequest v.Car
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		log.Println("Ошибка связывания данных:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные запроса"})
		return
	}

	if carsToUpdate != nil {
		if updateRequest.Brand != "" {
			carsToUpdate.Brand = updateRequest.Brand
		}
		if updateRequest.Model != "" {
			carsToUpdate.Model = updateRequest.Model
		}
		if updateRequest.Mileage != 0 {
			carsToUpdate.Mileage = updateRequest.Mileage
		}
		if updateRequest.Owners != 0 {
			carsToUpdate.Owners = updateRequest.Owners
		}

		if err := writeFile("file.json", items); err != nil {
			log.Println("Ошибка при записи в файл:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при записи в файл"})
			return
		}
		c.JSON(http.StatusOK, carsToUpdate)
	} else {
		c.JSON(http.StatusNoContent, nil)
	}

	if err := writeFile("file.json", items); err != nil {
		log.Println("Ошибка при записи в файл:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при записи в файл"})
		return
	}
}
func writeFile(filename string, data interface{}) error {
	fileWrite, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fileWrite.Close()

	updatedDataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if _, err := fileWrite.Write(updatedDataJSON); err != nil {
		return err
	}

	return nil
}
