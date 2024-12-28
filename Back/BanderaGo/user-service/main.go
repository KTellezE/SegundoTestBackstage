// main.go
package main

import (
	"application/config"
	"application/controllers"
	facadeImpl "application/facade/impl"
	"application/persistence/contexts"
	"application/persistence/repositories"
	repoImpl "application/persistence/repositories/impl"
	serviceImpl "application/services/impl"
	"fmt"
	"log"

	docs "application/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := config.LoadEnvVariables(); err != nil {
		log.Fatalf("No se pudieron cargar las variables de entorno: %v", err)
	}
	// Obtener el valor del puerto desde las variables de entorno
	port, err := config.GetEnvVariable("APP_PORT")
	if err != nil {
		log.Fatalf("No se pudo obtener el valor del puerto: %v", err)
	}

	// Configurar el enrutador Gin
	router := gin.Default()

	docs.SwaggerInfo.Title = "GoLang"

	// Inicializar la conexión a la base de datos
	userConfig, err := config.NewUserConfig()
	if err != nil {
		panic(err)
	}

	mySQLDB, err := contexts.NewMySQLDB(userConfig)
	if err != nil {
		log.Fatal(err)
	}

	var myGormDB repositories.GormDB = mySQLDB.DB

	// Crear instancia de UserRepositoryImpl
	userRepo := repoImpl.NewUserRepository(myGormDB)

	// Crear instancia de UserServiceImpl usando UserRepository
	userService := serviceImpl.NewUserService(userRepo)

	// Crear instancia de UserFacadeImpl usando UserService
	userFacade := facadeImpl.NewUserFacade(userService)

	// Crear instancia de UserController usando UserFacade
	userController := controllers.NewUserController(userFacade)

	// Ruta base para el grupo de endpoints de usuarios
	userGroup := router.Group("/api/users")
	{
		// Definir endpoints CRUD para usuarios dentro del grupo
		userGroup.POST("", userController.CreateUser)
		userGroup.GET("", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetSingleUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	// Configurar middleware de Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Servidor escuchando en el puerto %s", port)
	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
