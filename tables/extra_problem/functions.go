package extra_problem

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	ep := &ExtraProblem{}
	if err := c.Bind(ep); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateExtraProblem(ep); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&ep).Error; err != nil {
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
		Data:    ep,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	extra_problems := []ExtraProblem{}
	db := configuration.GetConnection()

	if err := db.Find(&extra_problems).Error; err != nil {
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
		Data:    extra_problems,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var extra_problem ExtraProblem
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	if err := db.First(&extra_problem, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&extra_problem).Error; err != nil {
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
		Data:    extra_problem,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	epi := c.QueryParam("extra_problem_id")
	epti := c.QueryParam("extra_problem_title")
	epte := c.QueryParam("extra_problem_text")
	epe := c.QueryParam("extra_problem_explanation")

	db := configuration.GetConnection()

	if err := db.Model(&ExtraProblem{}).Where("extra_problem_id = ?", epi).Updates(ExtraProblem{ExtraProblemTitle: epti, ExtraProblemText: epte, ExtraProblemExplanation: epe}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	extra_problems := []ExtraProblem{}
	if err := db.Find(&extra_problems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    extra_problems,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	var extra_problem ExtraProblem
	if err := db.First(&extra_problem, id).Error; err != nil {
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
		Data:    extra_problem,
	}
	return c.JSON(http.StatusOK, r)
}
