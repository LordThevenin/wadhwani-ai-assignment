package constants

import "golang.org/x/text/language"

var Matcher = language.NewMatcher([]language.Tag{
	language.English, // The first language is used as fallback.
	language.Hindi,
	language.Marathi,
	language.Telugu,
	language.Punjabi,
})
