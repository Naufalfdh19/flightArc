package router

import (
	"database/sql"
	"flight/middleware"
	"flight/modules/user/controller"
	"flight/modules/user/repo"
	"flight/modules/user/service"

	scheduleController "flight/modules/schedule/controller"
	scheduleRepo "flight/modules/schedule/repo"
	scheduleService "flight/modules/schedule/service"

	adminController "flight/modules/admin/controller"
	adminRepo "flight/modules/admin/repo"
	adminService "flight/modules/admin/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userController     controller.UserController
	scheduleController scheduleController.ScheduleController
	adminController    adminController.AdminController
}

func NewRouter(db *sql.DB) *gin.Engine {
	userController := setupUserController(db)
	scheduleController := setupScheduleController(db)
	adminController := setupAdminController(db)

	return setupRouter(Controller{
		userController:     userController,
		scheduleController: scheduleController,
		adminController:    adminController,
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
	userAuth.POST("register", c.userController.Register)

	protected := baseEndpoint.Group("/")
	protected.Use(middleware.CheckAuth)

	userProtected := protected.Group("/")
	userProtected.Use(middleware.CheckUserAuth)
	userProtected.PATCH("", c.userController.UpdateUser)
	userProtected.PATCH("user/update-password", c.userController.UpdatePassword)

	adminAuth := baseEndpoint.Group("/admin/auth")
	adminAuth.POST("login", c.adminController.Login)
	adminAuth.POST("register", c.adminController.Register)

	adminProtected := protected.Group("/")
	adminProtected.Use(middleware.CheckAdminAuth)
	adminProtected.GET("users", c.userController.GetUsers)

	return router
}

func setupUserController(db *sql.DB) controller.UserController {
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	return controller.NewUserController(userService)
}

func setupScheduleController(db *sql.DB) scheduleController.ScheduleController {
	scheduleRepo := scheduleRepo.NewScheduleRepo(db)
	scheduleService := scheduleService.NewScheduleService(scheduleRepo)
	return scheduleController.NewScheduleController(scheduleService)
}

func setupAdminController(db *sql.DB) adminController.AdminController {
	userRepo := repo.NewUserRepo(db)
	adminRepo := adminRepo.NewAdminRepo(db)
	adminService := adminService.NewUserService(adminRepo, userRepo)
	return adminController.NewAdminController(adminService)
}
