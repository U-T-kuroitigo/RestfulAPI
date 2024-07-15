package situation

import (
	"net/http"
	"strconv"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	s := &Situation{}
	if err := c.Bind(s); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateSituation(s); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&s).Error; err != nil {
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
		Data:    s,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	situations := []Situation{}
	db := configuration.GetConnection()

	if err := db.Find(&situations).Error; err != nil {
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
		Data:    situations,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var situation Situation
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	if err := db.First(&situation, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&situation).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Delete errorr",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Correctly Deleted",
		Data:    situation,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	si := c.QueryParam("situation_id")
	st := c.QueryParam("situation_title")
	sl, err := strconv.Atoi(c.QueryParam("situation_level"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Invalid parameter",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Model(&Situation{}).Where("situation_id = ?", si).Updates(Situation{SituationTitle: st, SituationLevel: uint(sl)}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	situations := []Situation{}
	if err := db.Find(&situations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    situations,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	var situation Situation
	if err := db.First(&situation, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Correctly consulted",
		Data:    situation,
	}
	return c.JSON(http.StatusOK, r)
}
