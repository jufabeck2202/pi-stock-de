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

	"github.com/jufabeck2202/piScraper/adaptors"
	"github.com/jufabeck2202/piScraper/messaging"
	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/routes"
	"github.com/jufabeck2202/piScraper/utils"
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
	// logEnv()
	go startScraper()

	go messaging.Init()

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
	// // Used for local testing
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	//Monitoring
	prometheus := fiberprometheus.New("pi-stock-de")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	//routes
	app.Static("/", "./frontend/build")
	app.Get("/api/v1/status", routes.GetTasksController)
	app.Post("/api/v1/alert", routes.AddTaskController)
	app.Delete("/api/v1/alert/", routes.DeleteTaskController)

	app.Listen(":3001")
}

func startScraper() {
	routes.Websites.Init()
	c := cron.New()
	searchPi(true)
	c.AddFunc("*/5 * * * *", func() {
		searchPi(false)
	})
	c.Start()
}

func searchPi(firstRun bool) {
	adaptorsList := make([]adaptors.Adaptor, 0)
	routes.Websites.Load()
	c := colly.NewCollector(
		colly.Async(true),
	)
	adaptorsList = append(adaptorsList, adaptors.NewBechtle(c), adaptors.NewRappishop(c), adaptors.NewOkdo(c), adaptors.NewBerryBase(c), adaptors.NewSemaf(c), adaptors.NewBuyZero(c), adaptors.NewELV(c), adaptors.NewWelectron(c), adaptors.NewPishop(c), adaptors.NewFunk24(c), adaptors.NewReichelt(c))
	for _, site := range adaptorsList {
		site.Run(routes.Websites)
	}

	for _, site := range adaptorsList {
		site.Wait()
	}

	if !firstRun {
		changes := routes.Websites.CheckForChanges()
		if len(changes) > 0 {
			log.Println("Found Updates: ", len(changes))
			scheduleUpdates(changes)
		}
	}
	routes.Websites.Save()
}

func scheduleUpdates(websites []utils.Website) {
	tasksToSchedule := []types.AlertTask{}

	for _, w := range websites {
		log.Printf("%s changed \n", w.URL)
		alert := routes.AlertManager.LoadAlerts(w.URL)
		for _, t := range alert {
			tasksToSchedule = append(tasksToSchedule, types.AlertTask{Website: w, Recipient: t.Recipient, Destination: t.Destination})
			log.Printf("scheduling update for %s and %s \n", w.URL, t.Recipient)
		}
	}
	log.Println("Found websites to update: ", len(tasksToSchedule))
	messaging.AddToQueue(tasksToSchedule)
}
