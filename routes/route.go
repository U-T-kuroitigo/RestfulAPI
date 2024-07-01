package routes

import (
	"github.com/U-T-kuroitigo/RestfulAPI/tables/chapter"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/situation"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/theme"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/user"
	"github.com/labstack/echo"
)

func userRoutes(e *echo.Echo) {
	e.GET("api/v2/users", user.GetAll)   //GetAll Users
	e.GET("api/v2/user", user.Get)       //GET one user
	e.POST("api/v2/user", user.Create)   //CREATE
	e.PUT("api/v2/user", user.Update)    //UPDATE
	e.DELETE("api/v2/user", user.Delete) //DELETE
}

func themeRoutes(e *echo.Echo) {
	e.GET("api/v2/themes", theme.GetAll)   //GetAll themes
	e.GET("api/v2/theme", theme.Get)       //GET one theme
	e.POST("api/v2/theme", theme.Create)   //CREATE
	e.PUT("api/v2/theme", theme.Update)    //UPDATE
	e.DELETE("api/v2/theme", theme.Delete) //DELETE
}

func chapterRoutes(e *echo.Echo) {
	e.GET("api/v2/chapters", chapter.GetAll)   //GetAll chapters
	e.GET("api/v2/chapter", chapter.Get)       //GET one chapter
	e.POST("api/v2/chapter", chapter.Create)   //CREATE
	e.PUT("api/v2/chapter", chapter.Update)    //UPDATE
	e.DELETE("api/v2/chapter", chapter.Delete) //DELETE
}

func situationRoutes(e *echo.Echo) {
	e.GET("api/v2/situations", situation.GetAll)   //GetAll chapters
	e.GET("api/v2/situation", situation.Get)       //GET one situation
	e.POST("api/v2/situation", situation.Create)   //CREATE
	e.PUT("api/v2/situation", situation.Update)    //UPDATE
	e.DELETE("api/v2/situation", situation.Delete) //DELETE
}

func StartRoutes(e *echo.Echo) {
	userRoutes(e)
	themeRoutes(e)
	chapterRoutes(e)
	situationRoutes(e)
}
