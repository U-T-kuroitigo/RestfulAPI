package problem

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	p := &Problem{}
	if err := c.Bind(p); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateProblem(p); err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()

	if err := db.Create(&p).Error; err != nil {
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
		Data:    p,
	}
	return c.JSON(http.StatusCreated, r)
}

func GetAll(c echo.Context) error {
	problems := []Problem{}
	db := configuration.GetConnection()

	if err := db.Find(&problems).Error; err != nil {
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
		Data:    problems,
	}
	return c.JSON(http.StatusOK, r)
}

func Delete(c echo.Context) error {
	var problem Problem
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	if err := db.First(&problem, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&problem).Error; err != nil {
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
		Data:    problem,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Update(c echo.Context) error {
	pi := c.QueryParam("problem_id")
	pti := c.QueryParam("problem_title")
	pte := c.QueryParam("problem_text")
	pe := c.QueryParam("problem_explanation")

	db := configuration.GetConnection()

	if err := db.Model(&Problem{}).Where("problem_id = ?", pi).Updates(Problem{ProblemTitle: pti, ProblemText: pte, ProblemExplanation: pe}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	problems := []Problem{}
	if err := db.Find(&problems).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    problems,
	}
	return c.JSON(http.StatusAccepted, r)
}

func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	var problem Problem
	if err := db.First(&problem, id).Error; err != nil {
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
		Data:    problem,
	}
	return c.JSON(http.StatusOK, r)
}
