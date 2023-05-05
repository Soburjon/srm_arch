package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"log"
	_ "srm_arch/internal/controller/docs"
	task_controller "srm_arch/internal/controller/http/v1"
	"srm_arch/internal/controller/routes"
	"srm_arch/internal/middleware"
	configs "srm_arch/internal/pkg/config"
	"srm_arch/internal/pkg/repository"
	"srm_arch/internal/pkg/utils"
	"srm_arch/internal/repository/postgres/admin"
	"srm_arch/internal/repository/postgres/programmer"
	"srm_arch/internal/repository/postgres/register"
	admin2 "srm_arch/internal/service/admin"
	programmer2 "srm_arch/internal/service/programmer"
	register2 "srm_arch/internal/service/register"
	admin3 "srm_arch/internal/usecase/admin"
	programmer3 "srm_arch/internal/usecase/programmer"
	register3 "srm_arch/internal/usecase/register"
)

var (
	fiberConfig = configs.FiberConfig()
	appConfig   = configs.Config()
)

// @title API
// @version 1
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name Soburjon
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// database
	postgresDB := repository.NewPostgres()
	//_, err := postgresDB.NewCreateTable().Model((*entity.User)(nil)).Exec(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}

	// repository
	adminRepo := admin.NewRepository(postgresDB)
	programmerRepo := programmer.NewRepository(postgresDB)
	registerRepo := register.NewRepository(postgresDB)

	//service
	adminService := admin2.NewService(adminRepo)
	programmerService := programmer2.NewService(programmerRepo)
	registerService := register2.NewService(registerRepo)

	// use case
	adminUseCase := admin3.NewUseCase(adminService)
	programmerUseCase := programmer3.NewUseCase(programmerService)
	registerUseCase := register3.NewUseCase(registerService)

	// controller
	taskController := task_controller.NewController(adminUseCase, programmerUseCase, registerUseCase)

	utils.MigrationsUp()

	app := fiber.New(fiberConfig)

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	middleware.FiberMiddleware(app)

	jwtRoleAuthorizer, err := middleware.NewJWTRoleAuthorizer(appConfig)
	if err != nil {
		log.Fatal("Could not initialize JWT Role Authorizer")
	}
	app.Use(middleware.NewAuthorizer(jwtRoleAuthorizer))
	routes.SwaggerRoute(app)
	routes.AdminRoutes(app, taskController)
	routes.RegisterRoutes(app, taskController)
	routes.ProgrammerRoutes(app, taskController)

	app.Listen(":8000")

}
