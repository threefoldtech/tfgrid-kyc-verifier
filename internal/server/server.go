package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "example.com/tfgrid-kyc-service/api/docs"
	"example.com/tfgrid-kyc-service/internal/clients/idenfy"
	"example.com/tfgrid-kyc-service/internal/clients/substrate"
	"example.com/tfgrid-kyc-service/internal/configs"
	"example.com/tfgrid-kyc-service/internal/handlers"
	"example.com/tfgrid-kyc-service/internal/middleware"
	"example.com/tfgrid-kyc-service/internal/repository"
	"example.com/tfgrid-kyc-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/storage/mongodb"
	"github.com/gofiber/swagger"
)

// implement server struct that have fiber app and config
type Server struct {
	app    *fiber.App
	config *configs.Config
}

func New(config *configs.Config) *Server {
	// debug log
	app := fiber.New()

	// Setup Limter Config and store
	/* ipLimiterstore := mongodb.New(mongodb.Config{
		ConnectionURI: config.MongoURI,
		Database:      config.DatabaseName,
		Collection:    "ip_limit",
		Reset:         false,
	}) */
	/* ipLimiterConfig := limiter.Config{ // TODO: use configurable parameters, also check if it works well after passing the request through an SSL gateway
		Max:                    5,
		Expiration:             24 * time.Hour,
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		Store:                  ipLimiterstore,
		// skip the limiter for localhost
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
	} */
	clientLimiterStore := mongodb.New(mongodb.Config{
		ConnectionURI: config.MongoURI,
		Database:      config.DatabaseName,
		Collection:    "client_limit",
		Reset:         false,
	})

	clientLimiterConfig := limiter.Config{ // TODO: use configurable parameters
		Max:                    10,
		Expiration:             24 * time.Hour,
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		Store:                  clientLimiterStore,
		// Use client id as key to limit the number of requests per client
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("X-Client-ID")
		},
	}

	// Global middlewares
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())
	app.Use(recover.New())
	app.Use(helmet.New())

	// Database connection
	db, err := repository.ConnectToMongoDB(config.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	database := db.Database(config.DatabaseName)

	// Initialize repositories
	tokenRepo := repository.NewMongoTokenRepository(database)
	verificationRepo := repository.NewMongoVerificationRepository(database)

	// Initialize services
	idenfyClient := idenfy.New(config.Idenfy)

	if err != nil {
		log.Fatalf("Failed to initialize idenfy client: %v", err)
	}

	substrateClient, err := substrate.New(config.TFChain)
	if err != nil {
		log.Fatalf("Failed to initialize substrate client: %v", err)
	}
	tokenService := services.NewTokenService(tokenRepo, idenfyClient, substrateClient, config.MinBalanceToVerifyAccount)
	verificationService := services.NewVerificationService(verificationRepo, idenfyClient, &config.Verification)
	coordinatorService := services.NewCoordinatorService(tokenService, verificationService)

	// Initialize handler
	handler := handlers.NewHandler(tokenService, verificationService, coordinatorService)

	// Routes
	app.Get("/docs/*", swagger.HandlerDefault)

	v1 := app.Group("/api/v1", middleware.AuthMiddleware(config.ChallengeWindow)) // TODO: add limiter.New(ipLimiterConfig)
	v1.Post("/token", limiter.New(clientLimiterConfig), handler.GetorCreateVerificationToken())
	v1.Get("/data", handler.GetVerificationData())
	v1.Get("/status", handler.GetVerificationStatus())

	// Webhook routes
	webhooks := app.Group("/webhooks/idenfy")
	webhooks.Post("/verification-update", handler.ProcessVerificationResult())
	webhooks.Post("/id-expiration", handler.ProcessDocExpirationNotification())

	return &Server{app: app, config: config}
}

func (s *Server) Start() {
	// Start server
	go func() {
		if err := s.app.Listen(":" + s.config.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
