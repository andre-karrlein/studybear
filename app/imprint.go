package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type imprint struct {
	app.Compo
}

func (imprint *imprint) Render() app.UI {
	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Impressum"),
							app.P().Body(
								app.Text("Angaben gemäß § 5 TMG"),
								app.Br(),
								app.Br(),
								app.H1().Text("StudyBear"),
								app.Br(),
								app.H3().Text("Betreiber der Website:"),
								app.P().Text("André Karrlein"),
								app.Br(),
								app.H3().Text("Adresse:"),
								app.P().Text("Tittmoninger Str. 2"),
								app.P().Text("83410 Laufen"),
								app.Br(),
								app.H3().Text("Kontakt:"),
								app.P().Text("Email: info@study-bear.de"),
								app.P().Text("Internet: study-bear.de"),
							),
					),
				),
			),
		),
	)
}

