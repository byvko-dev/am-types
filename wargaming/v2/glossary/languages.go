package glossary

const (
	LangEN   = "en"    // English
	LangRU   = "ru"    // Russian
	LangPL   = "pl"    // Polish
	LangDE   = "de"    // German
	LangFR   = "fr"    // French
	LangES   = "es"    // Spanish
	LangTR   = "tr"    // Turkish
	LangCS   = "cs"    // Czech
	LangTH   = "th"    // Thai
	LangKO   = "ko"    // Korean
	LangVI   = "vi"    // Vietnamese
	LangZhCH = "zh-cn" // Simplified Chinese
	LangZhTW = "zh-tw" // Traditional Chinese
)

var AllLanguages = []string{LangEN, LangRU, LangPL, LangDE, LangFR, LangES, LangTR, LangCS, LangTH, LangKO, LangVI, LangZhCH, LangZhTW}

func GetLocale(locale string) string {
	for _, l := range AllLanguages {
		if l == locale {
			return l
		}
	}
	return LangEN
}
