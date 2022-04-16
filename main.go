package main

import (
	"log"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"

	"github.com/jufabeck2202/piScraper/internal/adaptors"
	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
	"github.com/jufabeck2202/piScraper/internal/core/services/alertsrv"
	"github.com/jufabeck2202/piScraper/internal/core/services/captchasrv"
	"github.com/jufabeck2202/piScraper/internal/core/services/notificationsrv"
	"github.com/jufabeck2202/piScraper/internal/core/services/validatesrv"
	"github.com/jufabeck2202/piScraper/internal/core/services/websitesrv"
	"github.com/jufabeck2202/piScraper/internal/handlers"
	"github.com/jufabeck2202/piScraper/internal/repositories/platforms/mail"
	"github.com/jufabeck2202/piScraper/internal/repositories/platforms/pushover"
	"github.com/jufabeck2202/piScraper/internal/repositories/platforms/webhook.go"
	"github.com/jufabeck2202/piScraper/internal/repositories/redis"
)

/*
TODO:
- Implement Verify Email Server Route
- Add React Router to Frontend
- Add React route to verify email
- excape unwanted characters from verify token
*/
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found", err)
	}

	// Initialize the app
	redisRepository, err := redis.NewRedisRepository()
	if err != nil {
		panic("Could not connect to redis")
	}
	websiteService := websitesrv.New(redisRepository)
	alertService := alertsrv.New(redisRepository)
	captchaService, err := captchasrv.New()
	if err != nil {
		panic("Could not connect to captcha service")
	}
	validateService := validatesrv.New()

	// logEnv()
	go startScraper(websiteService, alertService)

	app := fiber.New()
	app.Use(helmet.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "Too many request",
			})
		},
	}))
	// Used for local testing
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	//Monitoring
	prometheus := fiberprometheus.New("pi-stock-de")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	//controllers
	getController := handlers.NewGetHandler(websiteService)
	alertController := handlers.NewAlertHandler(websiteService, validateService, captchaService, alertService)
	deleteController := handlers.NewDeleteHandler(websiteService, validateService, captchaService, alertService)

	//routes
	app.Static("/", "./frontend/build")
	app.Get("/api/v1/status", getController.Get)
	app.Post("/api/v1/alert", alertController.Post)
	app.Delete("/api/v1/alert/", deleteController.Delete)

	app.Listen(":3001")
}

func startScraper(websiteService ports.WebsiteService, alertService ports.AlertService) {
	mailService := mail.NewMail()
	pushoverServie := pushover.NewPushover()
	webhookService := webhook.NewWebhook()
	notificationService := notificationsrv.NewNotificationService(,)
	c := cron.New()
	searchPi(true, websiteService, alertService)
	c.AddFunc("*/5 * * * *", func() {
		searchPi(false, websiteService, alertService)
	})
	c.Start()
}

func searchPi(firstRun bool, websiteService ports.WebsiteService, alertService ports.AlertService) {
	adaptorsList := make([]ports.Adaptor, 0)
	websiteService.Load()
	c := colly.NewCollector(
		colly.Async(true),
	)
	adaptorsList = append(adaptorsList, adaptors.NewBechtle(c, websiteService), adaptors.NewRappishop(c, websiteService), adaptors.NewOkdo(c, websiteService), adaptors.NewBerryBase(c, websiteService), adaptors.NewSemaf(c, websiteService), adaptors.NewBuyZero(c, websiteService), adaptors.NewELV(c, websiteService), adaptors.NewWelectron(c, websiteService), adaptors.NewPishop(c, websiteService), adaptors.NewFunk24(c, websiteService), adaptors.NewReichelt(c, websiteService))
	for _, site := range adaptorsList {
		site.Run()
	}

	for _, site := range adaptorsList {
		site.Wait()
	}

	if !firstRun {
		changes := websiteService.CheckForChanges()
		if len(changes) > 0 {
			log.Println("Found Updates: ", len(changes))
			scheduleUpdates(changes, alertService)
		}
	}
	websiteService.Save()
}

func scheduleUpdates(websites domain.Websites, alertService ports.AlertService) {
	tasksToSchedule := make([]domain.AlertTask, len(websites))

	for _, w := range websites {
		log.Printf("%s changed \n", w.URL)
		alert := alertService.LoadAlerts(w.URL)
		for _, t := range alert {
			tasksToSchedule = append(tasksToSchedule, domain.AlertTask{Website: w, Recipient: t.Recipient, Destination: t.Destination})
			log.Printf("scheduling update for %s and %s \n", w.URL, t.Recipient)
		}
	}
	log.Println("Found websites to update: ", len(tasksToSchedule))
}
