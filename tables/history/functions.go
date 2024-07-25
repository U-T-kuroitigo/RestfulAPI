package history

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	history := &History{}
	if err := c.Bind(history); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateHistory(history); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&history).Error; err != nil {
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
		Data:    history,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	historys := []History{}
	db := configuration.GetConnection()

	if err := db.Find(&historys).Error; err != nil {
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
		Data:    historys,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var history History
	hi := c.QueryParam("history_id")

	db := configuration.GetConnection()

	if err := db.Where("history_id = ?", hi).First(&history).Error; err != nil {
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

	if err := db.Where("history_id = ?", hi).Delete(&history).Error; err != nil {
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
		Data:    history,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	hi := c.QueryParam("history_id")

	db := configuration.GetConnection()

	history := History{}

	if err := db.Where("history_id = ?", hi).First(&history).Error; err != nil {
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

	if err := db.Model(&History{}).Where("history_id = ?", hi).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("history_id = ?", hi).First(&history).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    history,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	hi := c.QueryParam("history_id")
	db := configuration.GetConnection()

	var history History
	if err := db.Where("history_id = ?", hi).First(&history).Error; err != nil {
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
		Data:    history,
	}
	return c.JSON(http.StatusOK, r)
}
