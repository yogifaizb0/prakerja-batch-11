package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	InitDatabase()
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.Start(":8000")
}

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja11?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error init database")
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    users,
	})
}
