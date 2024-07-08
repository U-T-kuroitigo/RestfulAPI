package theme

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(context echo.Context) error {
	t := &Theme{}
	if err := context.Bind(t); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateTheme(t); err != nil {
		return context.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()


	if err := db.Create(&t).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error creating",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Created Successfully",
		Data:    t,
	}
	return context.JSON(http.StatusCreated, r)
}

func GetAll(context echo.Context) error {
	themes := []Theme{}
	db := configuration.GetConnection()


	if err := db.Find(&themes).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Correctly consulted",
		Data:    themes,
	}
	return context.JSON(http.StatusOK, r)
}

func Delete(context echo.Context) error {
	var theme Theme
	ti := context.QueryParam("theme_id")

	db := configuration.GetConnection()


	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusNotFound, r)
	}

	if err := db.Where("theme_id = ?", ti).Delete(&theme).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Delete error",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Correctly Deleted",
		Data:    theme,
	}
	return context.JSON(http.StatusAccepted, r)
}

func Update(context echo.Context) error {
	ti := context.QueryParam("theme_id")
	tt := context.QueryParam("theme_title")

	db := configuration.GetConnection()

	theme := Theme{}

	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			print("\n\nerr is ",err , "\n")
			// print("\n\nerr code is ",err. , "\n")
			return context.JSON(http.StatusInternalServerError, response.Model{
				Code:    "404",
				Message: "not found",
				Data:    err.Error(),
			})
		} else {
			return context.JSON(http.StatusInternalServerError, response.Model{
				Code:    "500",
				Message: "Query error",
				Data:    err.Error(),
			})
		}
	}

	if err := db.Model(&Theme{}).Where("theme_id = ?", ti).Updates(Theme{ThemeTitle: tt}).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("theme_id = ?", ti).First(&theme).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
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
	return context.JSON(http.StatusAccepted, r)
}

func Get(context echo.Context) error {
	ti := context.QueryParam("theme_id")
	db := configuration.GetConnection()

	var theme Theme
	if err := db.Where("theme_id = ?", ti).First(&theme).Error;err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusNotFound, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Correctly consulted",
		Data:    theme,
	}
	return context.JSON(http.StatusOK, r)
}
