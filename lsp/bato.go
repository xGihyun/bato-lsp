package lsp

// TODO: 
// Make a proper type for this maybe
// Could put descriptions on separate markdown files
var CompletionMap = map[string]string{
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

  // Integer
  "ulit": "Ulitin.",
  "beses": "Ulitin.",
  // "kasunod"

  // Class
  "gumawa": "Gumawa ng bagong class.",
  "kumatawan": "Gumawa ng bagong class.",

  // Array
  // "baliktad"
  "umikot": "Paikutin",
  "itulak": "Maglagay ng bagong elemento sa array.",
  "pasok": "Maglagay ng bagong elemento sa array.",
  "ipasok": "Maglagay ng bagong elemento sa array.",
  "huli": "Kunin ang x huling elemento ng array.",
  "dulo": "Kunin ang x huling elemento ng array.",
  "una": "Kunin ang x unang elemento ng array.",
  "isaisahin": "Isa-isahin ang bawat elemento ng isang array.",
  "kada": "Isa-isahin ang bawat elemento ng isang array.",
  "bawat": "Isa-isahin ang bawat elemento ng isang array.",
  "kada_isa": "Isa-isahin ang bawat elemento ng isang array.",
  "sa_bawat_isa": "Isa-isahin ang bawat elemento ng isang array.",
  "bilang_bawat_isa": "Isa-isahin ang bawat elemento at index ng isang array.",
  // "haba"

	// Add the rest ...
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
