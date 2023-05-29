package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type privacy struct {
	app.Compo
}

func (privacy *privacy) Render() app.UI {
	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Datenschutz"),
						app.P().Body(
							app.Strong().Text(""),
							app.Br(),
							app.Text(""),
						),
					),
				),
			),
		),
	)
}
