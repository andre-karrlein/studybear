package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type home struct {
	app.Compo
}

func (home *home) Render() app.UI {
	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Class("title").Text("StudyBear"),
						app.P().Text("Holen dir bessere Noten mit unserer Lern-App StudyBear."),
						app.P().Text("Unsere Lern-App ist die perfekte Lösung für Schüler, die ihr Wissen in verschiedenen Fächern erweitern möchten. Mit einer Vielzahl von Übungsaufgaben kannst du deine Leistungen verbessern."),
						app.Br(),
						app.Br(),
						app.H2().Text("Bald verfügbar!"),
						//app.Raw("<a href='https://apps.apple.com/de/app/studybear/id6449251194?itsct=apps_box_badge&amp;itscg=30200' style='display: inline-block; overflow: hidden; border-radius: 13px; width: 250px; height: 83px;'><img src='https://tools.applemediaservices.com/api/badges/download-on-the-app-store/black/de-de?size=250x83&amp;releaseDate=1643673600' alt='Download on the App Store' style='border-radius: 13px; width: 250px; height: 83px;'></a>"),
					),
					app.Div().Class("img-box").Body(
						app.Img().Src("/web/images/logo.jpeg"),
					),
				),
			),
		),
	)
}
