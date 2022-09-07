package server

import (
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	logger "github.com/sirupsen/logrus"
)

type webApp struct {
	app *fiber.App
}

var defaultApp *webApp

func init() {
	defaultApp = InitWebApp(html.New("./views", ".html"))
}

func InitWebApp(views fiber.Views) *webApp {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 views,
		BodyLimit:             500 * 1024 * 1024,
	})
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(ctx *fiber.Ctx, e interface{}) {
			logger.Error(e)
		},
	}))
	app.Use(cors.New())

	return &webApp{app}
}

func (a *webApp) Listener(ln net.Listener) error {
	return a.app.Listener(ln)
}
func (a *webApp) Static(dir string) {
	a.app.Static("/", dir, fiber.Static{
		Compress: true,
	})
}
func (a *webApp) Group(prefix string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	return a.app.Group(prefix, handlers...)
}
func (a *webApp) Use(args ...interface{}) fiber.Router {
	return a.app.Use(args...)
}
func (a *webApp) All(path string, handlers ...fiber.Handler) fiber.Router {
	return a.app.All(path, handlers...)
}
func (a *webApp) Get(path string, handlers ...fiber.Handler) fiber.Router {
	return a.app.Get(path, handlers...)
}
func (a *webApp) Put(path string, handlers ...fiber.Handler) fiber.Router {
	return a.app.Put(path, handlers...)
}
func (a *webApp) Post(path string, handlers ...fiber.Handler) fiber.Router {
	return a.app.Post(path, handlers...)
}
func (a *webApp) Delete(path string, handlers ...fiber.Handler) fiber.Router {
	return a.app.Delete(path, handlers...)
}
func (a *webApp) Start(addr string) error {
	stack := a.app.Stack()
	for _, s := range stack {
		for _, r := range s {
			logger.Debug("Router registered: ", r.Method, " - ", r.Path)
		}
	}

	return a.app.Listen(addr)
}
func (a *webApp) Shutdown() error {
	return a.app.Shutdown()
}

func Listener(ln net.Listener) error {
	return defaultApp.Listener(ln)
}
func Static(dir string) {
	defaultApp.Static(dir)
}

func Group(prefix string, handlers ...func(*fiber.Ctx) error) fiber.Router {
	return defaultApp.Group(prefix, handlers...)
}
func Use(args ...interface{}) fiber.Router {
	return defaultApp.Use(args...)
}
func All(path string, handlers ...fiber.Handler) fiber.Router {
	return defaultApp.All(path, handlers...)
}
func Get(path string, handlers ...fiber.Handler) fiber.Router {
	return defaultApp.Get(path, handlers...)
}
func Put(path string, handlers ...fiber.Handler) fiber.Router {
	return defaultApp.Put(path, handlers...)
}
func Post(path string, handlers ...fiber.Handler) fiber.Router {
	return defaultApp.Post(path, handlers...)
}
func Delete(path string, handlers ...fiber.Handler) fiber.Router {
	return defaultApp.Delete(path, handlers...)
}
func Start(addr string) error {
	return defaultApp.Start(addr)
}
func Shutdown() error {
	return defaultApp.Shutdown()
}
