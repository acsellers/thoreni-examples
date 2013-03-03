package main

import (
	"./views"
	"github.com/acsellers/thoreni"
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
	App := app.RunnableApp{Router, Renderer}
	app.ListenAndServe(&App, ":3456")
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
	render.NewRendererFromLists(views.LayoutList, views.ViewList)
	Renderer = render.MasterRenderer
}
func AboutHandler(ctx *thoreni.Contextable) {
	ctx.RenderStatic("pages/about")
}
func ContactHandler(ctx *thoreni.Contextable) {
	ctx.RenderStatic("pages/contact")
}
func PostHandler(ctx *thoreni.Contextable) {
	ctx.Render("posts/show")
}

func PostsHandler(ctx *thoreni.Contextable) {
	ctx.Render("posts/index")
}

func CreatePostsHandler(ctx *thoreni.Contextable) {
	ctx.Redirect("/posts")
}

func UpdatePostsHandler(ctx *thoreni.Contextable) {
	ctx.Redirect("/posts")
}

func DeletePostsHandler(ctx *thoreni.Contextable) {
	ctx.Redirect("/posts")
}

func IndexHandler(ctx *thoreni.Contextable) {
	ctx.Render("posts/index")
}
