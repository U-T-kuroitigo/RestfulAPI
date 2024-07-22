package routes

import (
	"github.com/U-T-kuroitigo/RestfulAPI/tables/chapter"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/choice"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/extra_problem"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/extra_situation"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/problem"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/situation"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/theme"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/user"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/user_profile"
	"github.com/labstack/echo"
)

func userRoutes(e *echo.Echo) {
	e.GET("api/v2/users", user.GetAll)   //GetAll Users
	e.GET("api/v2/user", user.Get)       //GET one user
	e.POST("api/v2/user", user.Create)   //CREATE
	e.PUT("api/v2/user", user.Update)    //UPDATE
	e.DELETE("api/v2/user", user.Delete) //DELETE
}

func UserProfileRoutes(e *echo.Echo) {
	e.GET("api/v2/user_profiles", user_profile.GetAll)   //GetAll UserProfiles
	e.GET("api/v2/user_profile", user_profile.Get)       //GET one UserProfile
	e.POST("api/v2/user_profile", user_profile.Create)   //CREATE
	e.PUT("api/v2/user_profile", user_profile.Update)    //UPDATE
	e.DELETE("api/v2/user_profile", user_profile.Delete) //DELETE
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
	e.GET("api/v2/situations", situation.GetAll)   //GetAll situations
	e.GET("api/v2/situation", situation.Get)       //GET one situation
	e.POST("api/v2/situation", situation.Create)   //CREATE
	e.PUT("api/v2/situation", situation.Update)    //UPDATE
	e.DELETE("api/v2/situation", situation.Delete) //DELETE
}

func problemRoutes(e *echo.Echo) {
	e.GET("api/v2/problems", problem.GetAll)   //GetAll problems
	e.GET("api/v2/problem", problem.Get)       //GET one problem
	e.POST("api/v2/problem", problem.Create)   //CREATE
	e.PUT("api/v2/problem", problem.Update)    //UPDATE
	e.DELETE("api/v2/problem", problem.Delete) //DELETE
}

func choiceRoutes(e *echo.Echo) {
	e.GET("api/v2/choices", choice.GetAll)   //GetAll choices
	e.GET("api/v2/choice", choice.Get)       //GET one choice
	e.POST("api/v2/choice", choice.Create)   //CREATE
	e.PUT("api/v2/choice", choice.Update)    //UPDATE
	e.DELETE("api/v2/choice", choice.Delete) //DELETE
}

func extra_situationRoutes(e *echo.Echo) {
	e.GET("api/v2/extra_situations", extra_situation.GetAll)   //GetAll chapters
	e.GET("api/v2/extra_situation", extra_situation.Get)       //GET one situation
	e.POST("api/v2/extra_situation", extra_situation.Create)   //CREATE
	e.PUT("api/v2/extra_situation", extra_situation.Update)    //UPDATE
	e.DELETE("api/v2/extra_situation", extra_situation.Delete) //DELETE
}

func extra_problemRoutes(e *echo.Echo) {
	e.GET("api/v2/extra_problems", extra_problem.GetAll)   //GetAll problems
	e.GET("api/v2/extra_problem", extra_problem.Get)       //GET one problem
	e.POST("api/v2/extra_problem", extra_problem.Create)   //CREATE
	e.PUT("api/v2/extra_problem", extra_problem.Update)    //UPDATE
	e.DELETE("api/v2/extra_problem", extra_problem.Delete) //DELETE
}

func StartRoutes(e *echo.Echo) {
	userRoutes(e)
	UserProfileRoutes(e)
	themeRoutes(e)
	chapterRoutes(e)
	situationRoutes(e)
	problemRoutes(e)
	choiceRoutes(e)
	extra_problemRoutes(e)
	extra_situationRoutes(e)
}
