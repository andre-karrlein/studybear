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
							app.Text("Wir freuen uns über Ihr Interesse an unserer Website und unserer App. Der Schutz Ihrer persönlichen Daten ist uns ein wichtiges Anliegen. Nachfolgend informieren wir Sie ausführlich über den Umgang mit Ihren Daten.
1. Verantwortliche Stelle
Verantwortlich für die Verarbeitung Ihrer personenbezogenen Daten im Sinne der Datenschutz-Grundverordnung (DSGVO) ist StudyBear, Tittmoninger Str. 2, 83410 Laufen, Deutschland, info@study-bear.de.
2. Erhebung, Verarbeitung und Nutzung personenbezogener Daten
2.1 Besuch der Website
Bei Ihrem Besuch auf unserer Website verarbeiten wir keine personenbezogenen Daten von Ihnen. Es werden lediglich allgemeine, anonyme Informationen erhoben, die keinen Rückschluss auf Ihre Person ermöglichen und ausschließlich zu statistischen Zwecken ausgewertet werden.
2.2 Nutzung der App
Für die Nutzung unserer App ist eine Registrierung erforderlich. Bei der Registrierung erheben wir Ihre E-Mail-Adresse, um Ihnen den Zugang zur App zu ermöglichen und Ihnen Informationen über Updates oder neue Funktionen zukommen zu lassen. Die Verarbeitung Ihrer E-Mail-Adresse erfolgt auf Grundlage von Art. 6 Abs. 1 lit. b DSGVO.
Darüber hinaus speichert die App die von Ihnen gekauften Inhalte. Diese Informationen dienen ausschließlich der Bereitstellung der von Ihnen erworbenen Inhalte und werden nicht an Dritte weitergegeben, sofern keine rechtliche Verpflichtung besteht oder Sie ausdrücklich eingewilligt haben.
3. Weitergabe von Daten an Dritte
Eine Übermittlung Ihrer personenbezogenen Daten an Dritte erfolgt nur, wenn dies zur Vertragserfüllung erforderlich ist, Sie eingewilligt haben oder eine gesetzliche Verpflichtung hierzu besteht. Dabei können Dienstleister, die uns bei der Bereitstellung der App unterstützen, Zugriff auf Ihre Daten erhalten. Diese Dienstleister handeln in unserem Auftrag und sind vertraglich dazu verpflichtet, Ihre Daten vertraulich zu behandeln und ausschließlich im Rahmen der App-Bereitstellung zu verarbeiten.
4. Einsatz von AWS Cloud-Diensten
Unsere App nutzt AWS Cloud-Dienste (Amazon Web Services) zur Speicherung und Verarbeitung Ihrer Daten. AWS stellt sicher, dass Ihre Daten in Rechenzentren gespeichert werden, die hohe Sicherheitsstandards erfüllen. Die Datenübermittlung erfolgt verschlüsselt, um einen Schutz vor unbefugtem Zugriff zu gewährleisten.
5. Ihre Rechte
Sie haben das Recht, Auskunft über die von uns verarbeiteten personenbezogenen Daten zu erhalten. Außerdem haben Sie das Recht auf Berichtigung, Löschung oder Einschränkung der Verarbeitung Ihrer Daten, soweit dem keine gesetzlichen Aufbewahrungspflichten entgegenstehen. Unter bestimmten Voraussetzungen haben Sie auch das Recht auf Datenübertragbarkeit.
Sofern Sie in die Verarbeitung Ihrer Daten eingewilligt haben, können Sie Ihre Einwilligung jederzeit mit Wirkung für die Zukunft widerrufen. Darüber hinaus haben Sie das Recht, sich bei einer Datenschutz-Aufsichtsbehörde über die Verarbeitung Ihrer personenbezogenen Daten durch uns zu beschweren.
6. Kontaktaufnahme
Für Fragen zum Datenschutz oder zur Ausübung Ihrer Rechte können Sie sich jederzeit an uns wenden. Die Kontaktdaten finden Sie am Anfang dieser Datenschutzerklärung."),
						),
					),
				),
			),
		),
	)
}
