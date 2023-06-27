package controllers

import (
	"log"
	"net/http"
	"sesi6/configs"
	"sesi6/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateNewsController(c echo.Context) error {
	var news models.News
	c.Bind(&news)

	log.Println("req1: ", c)
	log.Println("req2: ", news.Content)

	result := configs.DB.Create(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    news,
	})

}

func UpdateController(c echo.Context) error {
	var news models.News
	var updatedNews models.News
	c.Bind(&updatedNews)

	var id, _ = strconv.Atoi(c.Param("id"))

	configs.DB.First(&news, id)

	news.Title = updatedNews.Title
	news.Content = updatedNews.Content

	result := configs.DB.Save(&news)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    news,
	})

}

func DetailNewsController(c echo.Context) error {
	log.Println("detail")
	var id, _ = strconv.Atoi(c.Param("id"))
	log.Println("id: ", id)

	var news models.News

	result := configs.DB.First(&news, id)
	log.Println("result:", result.Error)

	if result.Error != nil {
		log.Println("error:", result.Error)
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    news,
	})
}

func NewsController(c echo.Context) error {
	var data []models.News

	result := configs.DB.Find(&data)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    data,
	})
}
