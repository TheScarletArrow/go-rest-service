package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go-rest-service/config"
	"go-rest-service/models"
	"strconv"
)

func CreateExcel(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	for i, user := range users {
		index := strconv.Itoa(i)
		f.SetCellValue("Sheet1", "A"+index, user.Id)
		f.SetCellValue("Sheet1", "B"+index, user.Name)
		f.SetCellValue("Sheet1", "C"+index, user.Password)
		f.SetCellValue("Sheet1", "D"+index, user.Email)
		f.SetCellValue("Sheet1", "E"+index, user.CreatedAt)
		f.SetActiveSheet(1)
		if err := f.SaveAs("Book.xlsx"); err != nil {
			fmt.Println(err)
		}
	}

}
