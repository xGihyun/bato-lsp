// NOTE: For now, this is where I'll put all the hover definitions
package lsp

// TODO: Make a proper type for this maybe
var HoverDocs = map[string]string{
	// Kernel
	"mag_print":   "Magprint ng something.",
	"ulit_ulitin": "Ulit-ulitin ang code na nasa loob.",

	// String
	"baliktad":               "Baliktarin ang string.",
	"haba":                   "Kunin ang haba ng string.",
	"sa_malaking_titik":      "Kinakapitalisa ang unang letra ng string.",
	"sa_malalaking_titik":    "Ginagawang malalaking titik ang lahat ng letra sa string.",
	"sa_malaking_mga_titik":  "Ginagawang malalaking titik ang lahat ng letra sa string.",
	"sa_maliit_na_titik":     "Ginagawang maliliit na titik ang lahat ng letra sa string.",
	"sa_maliit_na_mga_titik": "Ginagawang maliliit na titik ang lahat ng letra sa string.",
	"igitna":                 "Ibinabalik ang string na nakagitna.",
	"kagatan":                "Tinatanggal ang huling newline sa string.",
	"palitan":                "Pinapalitan ang mga karakter sa string.",
	"siyasatin":              "Ibinabalik ang isang string na representasyon ng object para sa inspeksyon.",
	"kasunod":                "Ibinabalik ang kasunod na karakter ng string.",

	// TODO: Add the rest
}

// NOTE: I probably don't need this
func MapToCompletionItems(m map[string]string) []CompletionItem {
	items := make([]CompletionItem, 0, len(m))
	for label, detail := range m {
		items = append(items, CompletionItem{
			Label:  label,
			Detail: detail,
		})
	}
	return items
}
