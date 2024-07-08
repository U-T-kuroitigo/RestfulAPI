package extra_situation

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	es := &ExtraSituation{}
	if err := c.Bind(es); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateExtraSituation(es); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&es).Error; err != nil {
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
		Data:    es,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	extra_situations := []ExtraSituation{}
	db := configuration.GetConnection()

	if err := db.Find(&extra_situations).Error; err != nil {
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
		Data:    extra_situations,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var extra_situation ExtraSituation
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	if err := db.First(&extra_situation, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&extra_situation).Error; err != nil {
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
		Data:    extra_situation,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	esi := c.QueryParam("extra_situation_id")
	est := c.QueryParam("extra_situation_title")

	db := configuration.GetConnection()

	if err := db.Model(&ExtraSituation{}).Where("extra_situation_id = ?", esi).Updates(ExtraSituation{ExtraSituationTitle: est}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	extra_situations := []ExtraSituation{}
	if err := db.Find(&extra_situations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    extra_situations,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	var extra_situation ExtraSituation
	if err := db.First(&extra_situation, id).Error; err != nil {
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
		Data:    extra_situation,
	}
	return c.JSON(http.StatusOK, r)
}
