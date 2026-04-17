package router

import (
	"flight/middleware"
	"flight/modules/user/controller"
	"flight/modules/user/repo"
	"flight/modules/user/service"
	"flight/pkg/transaction"

	scheduleController "flight/modules/schedule/controller"
	scheduleRepo "flight/modules/schedule/repo"
	scheduleService "flight/modules/schedule/service"

	adminController "flight/modules/admin/controller"
	adminRepo "flight/modules/admin/repo"
	adminService "flight/modules/admin/service"

	planeController "flight/modules/plane/controller"
	planeRepo "flight/modules/plane/repo"
	planeService "flight/modules/plane/service"

	bookingController "flight/modules/booking/controller"
	bookingRepo "flight/modules/booking/repo"
	bookingService "flight/modules/booking/service"

	airlineRepo "flight/modules/airline/repo"

	paymentController "flight/modules/payment/controller"
	paymentProvider "flight/modules/payment/provider"
	paymentRepo "flight/modules/payment/repo"
	paymentService "flight/modules/payment/service"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Controller struct {
	userController     controller.UserController
	scheduleController scheduleController.ScheduleController
	adminController    adminController.AdminController
	planeController    planeController.PlaneController
	paymentController  paymentController.PaymentController
	bookingController  bookingController.BookingController
}

type Service struct {
	UserService     service.UserService
	ScheduleService scheduleService.ScheduleService
	AdminService    adminService.AdminService
	PlaneService    planeService.PlaneService
	BookingService  bookingService.BookingService
	PaymentService  paymentService.PaymentService
}

type Repo struct {
	userRepo     repo.UserRepo
	scheduleRepo scheduleRepo.ScheduleRepo
	adminRepo    adminRepo.AdminRepo
	planeRepo    planeRepo.PlaneRepo
	bookingRepo  bookingRepo.BookingRepo
	airlineRepo  airlineRepo.AirlineRepo
	seatRepo     scheduleRepo.SeatRepo
	ticketRepo   bookingRepo.TicketRepo
	paymentRepo  paymentRepo.PaymentRepo
	tx           transaction.TransactorRepo
}

func SetupRepo(db *gorm.DB) *Repo {
	return &Repo{
		userRepo:     repo.NewUserRepo(db),
		scheduleRepo: scheduleRepo.NewScheduleRepo(db),
		adminRepo:    adminRepo.NewAdminRepo(db),
		planeRepo:    planeRepo.NewPlaneRepo(db),
		bookingRepo:  bookingRepo.NewBookingRepo(db),
		airlineRepo:  airlineRepo.NewAirlineRepo(db),
		seatRepo:     scheduleRepo.NewSeatRepo(db),
		ticketRepo:   bookingRepo.NewTicketRepo(db),
		paymentRepo:  paymentRepo.NewPaymentRepo(db),
		tx:           transaction.NewTransactorRepo(db),
	}
}

func SetupService(repos *Repo, mqch *amqp.Channel) *Service {
	return &Service{
		UserService:     service.NewUserService(repos.userRepo),
		ScheduleService: scheduleService.NewScheduleService(repos.scheduleRepo),
		AdminService:    adminService.NewAdminService(repos.adminRepo, repos.userRepo),
		PlaneService:    planeService.NewPlaneService(repos.planeRepo, repos.airlineRepo),
		BookingService:  bookingService.NewBookingService(repos.bookingRepo, repos.seatRepo, repos.ticketRepo, repos.tx, mqch),
		PaymentService:  paymentService.NewPaymentService(repos.paymentRepo, repos.bookingRepo, repos.ticketRepo, repos.seatRepo, repos.tx, paymentProvider.NewManualVAProvider()),
	}
}

func NewRouter(db *gorm.DB, mqch *amqp.Channel) *gin.Engine {
	repos := SetupRepo(db)
	services := SetupService(repos, mqch)

	userController := controller.NewUserController(services.UserService)
	scheduleController := scheduleController.NewScheduleController(services.ScheduleService)
	adminController := adminController.NewAdminController(services.AdminService)
	planeController := planeController.NewPlaneController(services.PlaneService)
	paymentController := paymentController.NewPaymentController(services.PaymentService)
	bookingController := bookingController.NewBookingController(services.BookingService)

	return setupRouter(Controller{
		userController:     userController,
		scheduleController: scheduleController,
		adminController:    adminController,
		planeController:    planeController,
		paymentController:  paymentController,
		bookingController:  bookingController,
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
	scheduleGeneral.GET("", c.scheduleController.GetFlights)

	userAuth := baseEndpoint.Group("/user/auth")
	userAuth.POST("login", c.userController.Login)
	userAuth.POST("register", c.userController.Register)

	protected := baseEndpoint.Group("/")
	protected.Use(middleware.CheckAuth)

	profileProtected := protected.Group("/")
	profileProtected.GET("users/:id", c.userController.GetUserById)

	userProtected := protected.Group("/")
	userProtected.Use(middleware.CheckUserAuth)
	userProtected.PATCH("users/:id", c.userController.UpdateUser)
	userProtected.PATCH("user/update-password", c.userController.UpdatePassword)
	userProtected.POST("payments", c.paymentController.CreatePayment)
	userProtected.GET("payments/:id", c.paymentController.GetPaymentByID)

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

	bookingProtected := protected.Group("/")
	bookingProtected.Use(middleware.CheckUserAuth)
	bookingProtected.GET("bookings", c.bookingController.GetBookings)
	bookingProtected.POST("bookings", c.bookingController.AddBookings)
	bookingProtected.GET("bookings/:id", c.bookingController.GetBookingsById)

	baseEndpoint.POST("payments/webhook", c.paymentController.HandleWebhook)

	return router
}
