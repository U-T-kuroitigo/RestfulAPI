package user_profile

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {
	user_profile := &UserProfile{}
	if err := c.Bind(user_profile); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateUserProfile(user_profile); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&user_profile).Error; err != nil {
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
		Data:    user_profile,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	user_profiles := []UserProfile{}
	db := configuration.GetConnection()

	if err := db.Find(&user_profiles).Error; err != nil {
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
		Data:    user_profiles,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var user_profile UserProfile
	ui := c.QueryParam("user_id")

	db := configuration.GetConnection()

	if err := db.Where("user_id = ?", ui).First(&user_profile).Error; err != nil {
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

	if err := db.Where("user_id = ?", ui).Delete(&user_profile).Error; err != nil {
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
		Data:    user_profile,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	ui := c.QueryParam("user_id")

	db := configuration.GetConnection()

	user_profile := UserProfile{}

	if err := db.Where("user_id = ?", ui).First(&user_profile).Error; err != nil {
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

	if err := db.Model(&UserProfile{}).Where("user_id = ?", ui).Updates(updates).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	if err := db.Where("user_id = ?", ui).First(&user_profile).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    user_profile,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	ui := c.QueryParam("user_id")
	db := configuration.GetConnection()

	var user_profile UserProfile
	if err := db.Where("user_id = ?", ui).First(&user_profile).Error; err != nil {
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
		Data:    user_profile,
	}
	return c.JSON(http.StatusOK, r)
}
