package main

import (
	"githbu.com/niklastreml/htmxual/pkg/handlers"
	"githbu.com/niklastreml/htmxual/pkg/renderer"
	"githbu.com/niklastreml/htmxual/pkg/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = &renderer.TemplRender{}

	counterGroup := r.Group("/count")
	cs := services.NewCounterService()
	handlers.NewCountHandler(*cs).Register(counterGroup)

	todoGroup := r.Group("/todo")
	ts := services.NewTodoService()
	handlers.NewTodoHandler(ts).Register(todoGroup)

	basicAuthGroup := r.Group("/basic")
	bas := services.BasicAuthService()
	handlers.NewBasicAuthHandler(bas).Register(basicAuthGroup)

	socketGroup := r.Group("/sse")
	ss := services.NewSseService()
	handlers.NewSseHandler(ss).Register(socketGroup)

	r.Run()
}
