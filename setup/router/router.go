package router

import (
	"flight/middleware"
	"flight/modules/user/controller"
	"flight/modules/user/repo"
	"flight/modules/user/service"

	scheduleController "flight/modules/schedule/controller"
	scheduleRepo "flight/modules/schedule/repo"
	scheduleService "flight/modules/schedule/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Controller struct {
	userController controller.UserController
	scheduleController scheduleController.ScheduleController
}

func NewRouter(db *pgx.Conn) *gin.Engine {
	userController := setupUserController(db)
	scheduleController := setupScheduleController(db)

	return setupRouter(Controller{
		userController: userController,
		scheduleController: scheduleController,
	})
}

func setupRouter(c Controller) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.LogMiddleware)
	router.Use(middleware.ErrorMiddleware)

	baseEndpoint := router.Group("api/v1")

	userGeneral := baseEndpoint.Group("users")
	userGeneral.GET("/:id", c.userController.GetUserById)
	userGeneral.DELETE("/:id", c.userController.DeleteUserById)
	
	scheduleGeneral := baseEndpoint.Group("schedules")
	scheduleGeneral.GET("", c.scheduleController.GetSchedules)

	userAuth := baseEndpoint.Group("/user/auth")
	userAuth.POST("login", c.userController.Login)

	protected := baseEndpoint.Group("/")
	protected.Use(middleware.CheckAuth)

	userProtected := protected.Use(middleware.CheckUserAuth)
	userProtected.PATCH("/:id", c.userController.UpdateUserById)

	adminProtected := protected.Use(middleware.CheckAdminAuth)
	adminProtected.GET("users", c.userController.GetUsers)	

	return router
}
// test

func setupUserController(db *pgx.Conn) controller.UserController {
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	return controller.NewUserController(userService)
}

func setupScheduleController(db *pgx.Conn) scheduleController.ScheduleController {
	scheduleRepo := scheduleRepo.NewScheduleRepo(db)
	scheduleService := scheduleService.NewScheduleService(scheduleRepo)
	return scheduleController.NewScheduleController(scheduleService)
}