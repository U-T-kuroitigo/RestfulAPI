package user

import (
	"net/http"
	"strconv"

	"github.com/U-T-kuroitigo/RestfulAPI/configuration"
	"github.com/U-T-kuroitigo/RestfulAPI/response"
	"github.com/labstack/echo"
)

// Create crea un nuevo usuario
func Create(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		r := response.Model{
			Code:    "400",
			Message: "Estructura incorrecta",
			Data:    err.Error(), // errをそのまま返すのではなく、メッセージに変換する
		}
		return c.JSON(http.StatusBadRequest, r)
	}

	db := configuration.GetConnection()


	if err := db.Create(&u).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al crear",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "201",
		Message: "Creado Correctamente",
		Data:    u,
	}
	return c.JSON(http.StatusCreated, r)
}

// GetAll Obtiene todos los datos
func GetAll(c echo.Context) error {
	users := []User{}
	db := configuration.GetConnection()


	if err := db.Find(&users).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al consultar",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Consultado Correctamente",
		Data:    users,
	}
	return c.JSON(http.StatusOK, r)
}

// Delete elimina un usuario por su id
func Delete(c echo.Context) error {
	var usuario User
	id := c.QueryParam("id")

	db := configuration.GetConnection()


	if err := db.First(&usuario, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "Usuario no encontrado",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	if err := db.Delete(&usuario).Error; err != nil {
		r := response.Model{
			Code:    "500",
			Message: "Error al eliminar",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, r)
	}

	r := response.Model{
		Code:    "202",
		Message: "Eliminado Correctamente",
		Data:    usuario,
	}
	return c.JSON(http.StatusAccepted, r)
}

// Update actualiza los campos
func Update(c echo.Context) error {
	i := c.QueryParam("id")
	fn := c.QueryParam("firstname")
	ltn := c.QueryParam("lastname")
	ag, err := strconv.Atoi(c.QueryParam("age"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Model{
			Code:    "400",
			Message: "Parámetro inválido",
			Data:    err.Error(),
		})
	}

	db := configuration.GetConnection()


	if err := db.Model(&User{}).Where("ID = ?", i).Updates(User{FirstName: fn, LastName: ltn, Age: ag}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error al actualizar",
			Data:    err.Error(),
		})
	}

	users := []User{}
	if err := db.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.Model{
			Code:    "500",
			Message: "Error al consultar",
			Data:    err.Error(),
		})
	}

	r := response.Model{
		Code:    "202",
		Message: "Actualizado Correctamente",
		Data:    users,
	}
	return c.JSON(http.StatusAccepted, r)
}

// Get trae un solo usuario por su ID
func Get(c echo.Context) error {
	id := c.QueryParam("id")

	db := configuration.GetConnection()

	var user User
	if err := db.First(&user, id).Error; err != nil {
		r := response.Model{
			Code:    "404",
			Message: "Usuario no encontrado",
			Data:    err.Error(),
		}
		return c.JSON(http.StatusNotFound, r)
	}

	r := response.Model{
		Code:    "200",
		Message: "Consultado correctamente",
		Data:    user,
	}
	return c.JSON(http.StatusOK, r)
}
