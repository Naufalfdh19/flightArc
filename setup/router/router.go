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

	planeController "flight/modules/plane/controller"
	planeRepo "flight/modules/plane/repo"
	planeService "flight/modules/plane/service"

	tokenController "flight/modules/token/controller"
	tokenRepo "flight/modules/token/repo"
	tokenService "flight/modules/token/service"

	airlineRepo "flight/modules/airline/repo"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userController     controller.UserController
	scheduleController scheduleController.ScheduleController
	adminController    adminController.AdminController
	planeController    planeController.PlaneController
	tokenControlelr    tokenController.TokenController
}

func NewRouter(db *sql.DB) *gin.Engine {
	userController := setupUserController(db)
	scheduleController := setupScheduleController(db)
	adminController := setupAdminController(db)
	planeController := setupPlaneController(db)
	tokenController := setupTokenController(db)

	return setupRouter(Controller{
		userController:     userController,
		scheduleController: scheduleController,
		adminController:    adminController,
		planeController:    planeController,
		tokenControlelr:    tokenController,
	})
}

func setupRouter(c Controller) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.LogMiddleware)
	router.Use(middleware.ErrorMiddleware)

	baseEndpoint := router.Group("api/v1")

	userGeneral := baseEndpoint.Group("users")
	userGeneral.DELETE("/:id", c.userController.DeleteUserById)

	scheduleGeneral := baseEndpoint.Group("schedules")
	scheduleGeneral.GET("", c.scheduleController.GetSchedules)

	userAuth := baseEndpoint.Group("/user/auth")
	userAuth.POST("login", c.userController.Login)
	userAuth.POST("register", c.userController.Register)

	tokenGeneral := baseEndpoint.Group("/tokens")
	tokenGeneral.GET("", c.tokenControlelr.GenerateNewAccessToken)

	protected := baseEndpoint.Group("/")
	protected.Use(middleware.CheckAuth)

	profileProtected := protected.Group("/")
	profileProtected.GET("users/:id", c.userController.GetUserById)

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

	airlineAdminProtected := protected.Group("/")
	airlineAdminProtected.Use(middleware.CheckAirlineAdminAuth)
	airlineAdminProtected.POST("planes", c.planeController.AddPlane)
	airlineAdminProtected.PATCH("planes/:id/update-seats", c.planeController.UpdateSeats)

	return router
}

func setupUserController(db *sql.DB) controller.UserController {
	tokenRepo := tokenRepo.NewTokenRepo(db)
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo, tokenRepo)
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
	adminService := adminService.NewAdminService(adminRepo, userRepo)
	return adminController.NewAdminController(adminService)
}

func setupPlaneController(db *sql.DB) planeController.PlaneController {
	planeRepo := planeRepo.NewPlaneRepo(db)
	airlineRepo := airlineRepo.NewAirlineRepo(db)
	planeService := planeService.NewPlaneService(planeRepo, airlineRepo)
	return planeController.NewPlaneController(planeService)
}

func setupTokenController(db *sql.DB) tokenController.TokenController {
	tokenRepo := tokenRepo.NewTokenRepo(db)
	tokenService := tokenService.NewTokenService(tokenRepo)
	return tokenController.NewTokenController(tokenService)
}
