package theme

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	theme := &Theme{}
	if err := c.Bind(theme); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateTheme(theme); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&theme).Error; err != nil {
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
		Data:    theme,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	themes := []Theme{}
	db := configuration.GetConnection()

	if err := db.Find(&themes).Error; err != nil {
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
		Data:    themes,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var theme Theme
	ti := c.QueryParam("theme_id")

	db := configuration.GetConnection()

	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
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

	if err := db.Where("theme_id = ?", ti).Delete(&theme).Error; err != nil {
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
		Data:    theme,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	ti := c.QueryParam("theme_id")

	db := configuration.GetConnection()

	theme := Theme{}

	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
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

	if err := db.Model(&Theme{}).Where("theme_id = ?", ti).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    theme,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	ti := c.QueryParam("theme_id")
	db := configuration.GetConnection()

	var theme Theme
	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
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
		Data:    theme,
	}
	return c.JSON(http.StatusOK, r)
}
