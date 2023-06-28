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

	result := configs.DB.Create(&news)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, models.BaseResponse{
			Message: "Failed",
			Data:    news,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    news,
	})

}

func UpdateController(c echo.Context) error {
	var id, err1 = strconv.Atoi(c.Param("id"))
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Message: "Request Invalid",
			Data:    nil,
		})
	}

	var news, updatedNews models.News
	c.Bind(&updatedNews)

	result := configs.DB.First(&news, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	news.Title = updatedNews.Title
	news.Content = updatedNews.Content

	if err := configs.DB.Save(&news); err != nil {
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
	var id, _ = strconv.Atoi(c.Param("id"))

	var news models.News

	result := configs.DB.First(&news, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Not Found",
			Data:    nil,
		})
	}
	log.Println("rowAffected:", result.RowsAffected)

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

func DeleteNewsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var news models.News

	result := configs.DB.Where("id = ?", id).Delete(&news)

	log.Println("result:", result.RowsAffected)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, models.BaseResponse{
			Message: "Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    nil,
	})
}
