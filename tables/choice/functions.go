package choice

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	ch := &Choice{}
	if err := c.Bind(ch); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateChoice(ch); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&ch).Error; err != nil {
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
		Data:    ch,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	choices := []Choice{}
	db := configuration.GetConnection()

	if err := db.Find(&choices).Error; err != nil {
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
		Data:    choices,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var choice Choice
	ci := c.QueryParam("choice_id")

	db := configuration.GetConnection()

	if err := db.Where("choice_id = ?", ci).First(&choice).Error; err != nil {
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

	if err := db.Where("choice_id = ?", ci).Delete(&choice).Error; err != nil {
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
		Data:    choice,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	ci := c.QueryParam("choice_id")

	db := configuration.GetConnection()

	choice := Choice{}

	if err := db.Where("choice_id = ?", ci).First(&choice).Error; err != nil {
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

	if err := db.Model(&Choice{}).Where("choice_id = ?", ci).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("choice_id = ?", ci).First(&choice).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    choice,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	ci := c.QueryParam("choice_id")
	db := configuration.GetConnection()

	var choice Choice
	if err := db.Where("choice_id = ?", ci).First(&choice).Error; err != nil {
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
		Data:    choice,
	}
	return c.JSON(http.StatusOK, r)
}
