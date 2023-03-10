package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	// By default Amber engine minifies the template before render.
	app.RegisterView(iris.Amber("./views", ".amber"))

	app.Get("/", index)

	app.Listen(":8080")
}

func index(ctx iris.Context) {
	data := iris.Map{
		"Title":      "Page Title",
		"FooterText": "Footer contents",
		"Message":    "Main contents",
	}

	// On Amber this is ignored:  ctx.ViewLayout("layouts/main")
	// Layouts are only rendered from inside the index page itself
	// using the "extends" keyword.
	if err := ctx.View("index", data); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
