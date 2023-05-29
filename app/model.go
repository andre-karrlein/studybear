package main

type climate struct {
	Heading      string `json:"heading"`
	Content   string `json:"content"`
}

func getClimateData() []climate {
	climateData := []climate{
		{
			Heading:   "Klimaschutz",
			Content:   "Für eine ambitionierte Klimapolitik einsetzen, um katastrophale Auswirkungen des Klimawandels zu verhindern."
		},
		{
			Heading:   "Energiewende",
			Content:   "Forderung einer rasche Umstellung auf erneuerbare Energien, um fossile Brennstoffe zu reduzieren und eine nachhaltige Zukunft zu gewährleisten"
		},
		{
			Heading:   "Grünes Wachstum",
			Content:   "Ich sehe den Klimawandel als Chance für grünes Wachstum und fördern Investitionen in nachhaltige Technologien und Arbeitsplätze"
		},
		{
			Heading:   "Mobilität",
			Content:   "Anstreben einer schrittweise Abschaffung von Verbrennungsmotoren zugunsten von emissionsfreien Alternativen wie Elektroantriben."
		},
		{
			Heading:   "Heizen",
			Content:   "Für die Umstellung von fossilen Brennstoffen auf erneuerbare Energien in der Gebäudeheizung einsetzen , um den CO2-Ausstoß zu reduzieren und die Energiewende voranzutreiben."
		},
	}

	return climateData
}