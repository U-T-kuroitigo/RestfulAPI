package chapter

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

func Create(context echo.Context) error {
	c := &Chapter{}
	if err := context.Bind(c); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Incorrect structure",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateChapter(c); err != nil {
		return context.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Failed validation",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()


	if err := db.Create(&c).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error creatingr",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Created Successfully",
		Data:    c,
	}
	return context.JSON(http.StatusCreated, r)
}

func GetAll(context echo.Context) error {
	chapters := []Chapter{}
	db := configuration.GetConnection()


	if err := db.Find(&chapters).Error; err != nil {
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
		Data:    chapters,
	}
	return context.JSON(http.StatusOK, r)
}

func Delete(context echo.Context) error {
	var chapter Chapter
	id := context.QueryParam("id")

	db := configuration.GetConnection()


	if err := db.First(&chapter, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "not found",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&chapter).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Delete errorr",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Correctly Deleted",
		Data:    chapter,
	}
	return context.JSON(http.StatusAccepted, r)
}

func Update(context echo.Context) error {
	ci := context.QueryParam("chapter_id")
	ct := context.QueryParam("chapter_title")

	db := configuration.GetConnection()


	if err := db.Model(&Chapter{}).Where("chapter_id = ?", ci).Updates(Chapter{ChapterTitle: ct}).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error updating",
			Data:    err.Error(),
		})
	}

	chapters := []Chapter{}
	if err := db.Find(&chapters).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Query error",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Updated successfully",
		Data:    chapters,
	}
	return context.JSON(http.StatusAccepted, r)
}

func Get(context echo.Context) error {
	id := context.QueryParam("id")

	db := configuration.GetConnection()

	var chapter Chapter
	if err := db.First(&chapter, id).Error; err != nil {
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
		Data:    chapter,
	}
	return context.JSON(http.StatusOK, r)
}
