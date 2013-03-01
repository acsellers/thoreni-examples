package main

import (
	"github.com/acsellers/thor/views"
	"github.com/acsellers/thoreni/app"
	"github.com/acsellers/thoreni/render"
	"github.com/acsellers/thoreni/router"
)

var (
	Router   *router.Router
	Renderer *render.TemplateRenderer
)

func init() {
	SetupRouter()
	SetupTemplates()
}

func main() {
	App := app.Runnable{Router, Renderer}
	app.Run(&App)
}

func SetupRouter() {
	Router = router.NewRouter()
	Router.Get("/about$", AboutHandler)
	Router.Get("/contact$", ContactHandler)
	Router.Get("/posts/:id", PostHandler)
	Router.Get("/posts/", PostsHandler)
	Router.Post("/posts/", CreatePostsHandler)
	Router.Post("/posts/:id", UpdatePostsHandler)
	Router.Delete("/posts/:id", DeletePostsHandler)
	Router.Root(IndexHandler)
}

func SetupTemplates() {

}
func AboutHandler(ctx Contextable) {
	ctx.RenderStatic("pages/about", nil)
}
func ContactHandler(ctx Contextable) {
	ctx.RenderStatic("pages/contact", nil)
}
func PostHandler(ctx Contextable) {
	ctx.Render("posts/show", nil)
}

func PostsHandler(ctx Contextable) {
	ctx.Render("posts/index", nil)
}

func CreatePostsHandler(ctx Contextable) {
	ctx.Redirect("/posts")
}

func UpdatePostsHandler(ctx Contextable) {
	ctx.Redirect("/posts")
}

func DeletePostsHandler(ctx Contextable) {
	ctx.Redirect("/posts")
}

func IndexHandler(ctx Contextable) {
	ctx.Render("posts/index", nil)
}
