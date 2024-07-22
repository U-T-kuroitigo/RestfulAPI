package chapter

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	chapter := &Chapter{}
	if err := c.Bind(chapter); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateChapter(chapter); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&chapter).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error creatingr",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Created Successfully",
		Data:    chapter,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	chapters := []Chapter{}
	db := configuration.GetConnection()

	if err := db.Find(&chapters).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Correctly consulted",
		Data:    chapters,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var chapter Chapter
	ci := c.QueryParam("chapter_id")

	db := configuration.GetConnection()

	if err := db.Where("chapter_id = ?", ci).First(&chapter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "404",
				Message: "not found",
				Data:    err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "500",
				Message: "Query error",
				Data:    err.Error(),
			})
		}
	}

	if err := db.Where("chapter_id = ?", ci).Delete(&chapter).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Delete error",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Correctly Deleted",
		Data:    chapter,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	ci := c.QueryParam("chapter_id")

	db := configuration.GetConnection()

	chapter := Chapter{}

	if err := db.Where("chapter_id = ?", ci).First(&chapter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "404",
				Message: "not found",
				Data:    err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "500",
				Message: "Query error",
				Data:    err.Error(),
			})
		}
	}

	// リクエストボディをマップに変換
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Invalid request body",
			Data:    err.Error(),
		})
	}

	// 更新対象のフィールドを明示的に指定
	updates := make(map[string]interface{})
	for key, value := range requestBody {
		updates[key] = value
	}

	if err := db.Model(&Chapter{}).Where("chapter_id = ?", ci).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("chapter_id = ?", ci).First(&chapter).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    chapter,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	ci := c.QueryParam("chapter_id")
	db := configuration.GetConnection()

	var chapter Chapter
	if err := db.Where("chapter_id = ?", ci).First(&chapter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "404",
				Message: "not found",
				Data:    err.Error(),
			})
		} else {
			return c.JSON(http.StatusInternalServerError, response.Model{
				Code:    "500",
				Message: "Query error",
				Data:    err.Error(),
			})
		}
	}

	r := response.Model{
		Code:    "200",
		Message: "Correctly consulted",
		Data:    chapter,
	}
	return c.JSON(http.StatusOK, r)
}
