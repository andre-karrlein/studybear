package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo
}

func (navbar *navbar) Render() app.UI {
	return app.Nav().Class("navbar").Body(
		app.Div().Class("logo").Body(
			app.A().Href("/").Body(
				app.Img().Width(80).Src("/web/images/logo.png"),
			),
		),
		app.Ul().Class("menu").Body(
			app.Li().Body(
				app.A().Href("/privacy.html").Text("Datenschutz"),
			),
			app.Li().Body(
				app.A().Href("/imprint.html").Text("Impressum"),
			),
		),
	)
}
