package main

import (
	"log"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &home{})
	app.Route("/privacy.html", &privacy{})
	app.Route("/imprint.html", &imprint{})
	app.RunWhenOnBrowser()

	if os.Getenv("GITHUB") == "TRUE" {
		err := app.GenerateStaticWebsite(".", &app.Handler{
			Name:        "Karrlein.com",
			Title:       "Karrlein.com",
			Description: "Home of André Karrlein",
			Icon: app.Icon{
				Default:    "/web/images/logo.png", // Specify default favicon.
				Large:      "/web/images/logo.png", // Specify large favicon
				AppleTouch: "/web/images/logo.png", // Specify icon on IOS devices.
			},
			Resources: app.GitHubPages("studybear"),
			Styles: []string{
				"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
				"/web/css/style.css",
			},
			ThemeColor: "#AF52DE",
		})

		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := app.GenerateStaticWebsite(".", &app.Handler{
			Name:        "StudyBear",
			Title:       "StudyBear",
			Description: "StudyBear learning platform",
			Icon: app.Icon{
				Default:    "/web/images/logo.png", // Specify default favicon.
				Large:      "/web/images/logo.png", // Specify large favicon
				AppleTouch: "/web/images/logo.png", // Specify icon on IOS devices.
			},
			Styles: []string{
				"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.2/css/all.min.css",
				"/web/css/style.css",
			},
			ThemeColor: "#AF52DE",
		})

		if err != nil {
			log.Fatal(err)
		}
	}
}
