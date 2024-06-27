package theme

import (
	"net/http"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

// Create crea un nuevo usuario
func Create(context echo.Context) error {
	t := &Theme{}
	if err := context.Bind(t); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Estructura incorrecta",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusBadRequest, r)
	}

	if err := ValidateTheme(t); err != nil {
		return context.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Validación fallida",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()


	if err := db.Create(&t).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al crear",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Creado Correctamente",
		Data:    t,
	}
	return context.JSON(http.StatusCreated, r)
}

// GetAll Obtiene todos los datos
func GetAll(context echo.Context) error {
	themes := []Theme{}
	db := configuration.GetConnection()


	if err := db.Find(&themes).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al consultar",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Consultado Correctamente",
		Data:    themes,
	}
	return context.JSON(http.StatusOK, r)
}

// Delete elimina un usuario por su id
func Delete(context echo.Context) error {
	var usuario Theme
	id := context.QueryParam("id")

	db := configuration.GetConnection()


	if err := db.First(&usuario, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "Usuario no encontrado",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&usuario).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al eliminar",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Eliminado Correctamente",
		Data:    usuario,
	}
	return context.JSON(http.StatusAccepted, r)
}

// Update actualiza los campos
func Update(context echo.Context) error {
	ti := context.QueryParam("theme_id")
	tt := context.QueryParam("theme_title")

	db := configuration.GetConnection()


	if err := db.Model(&Theme{}).Where("theme_id = ?", ti).Updates(Theme{ThemeTitle: tt}).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error al actualizar",
			Data:    err.Error(),
		})
	}

	themes := []Theme{}
	if err := db.Find(&themes).Error; err != nil {
		return context.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error al consultar",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Actualizado Correctamente",
		Data:    themes,
	}
	return context.JSON(http.StatusAccepted, r)
}

// Get trae un solo usuario por su ID
func Get(context echo.Context) error {
	id := context.QueryParam("id")

	db := configuration.GetConnection()

	var theme Theme
	if err := db.First(&theme, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "Usuario no encontrado",
			Data:    err.Error(),
		}
		return context.JSON(http.StatusNotFound, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Consultado correctamente",
		Data:    theme,
	}
	return context.JSON(http.StatusOK, r)
}