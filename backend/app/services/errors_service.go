package services

import (
	"github.com/go-raptor/raptor"
)

type ErrorsService struct {
	raptor.Service

	messages map[string]map[string]string
	lang     string
}

func NewErrorsService(lang string) *ErrorsService {
	return &ErrorsService{
		lang: lang,
		messages: map[string]map[string]string{
			"INVALID_URL": {
				"hr": "Neispravan link",
			},
			"URL_NOT_ABSOLUTE": {
				"hr": "Link nije apsolutan",
			},
			"URL_NOT_HTTP": {
				"hr": "Link ne sadrži HTTP ili HTTPS prefiks",
			},
			"HOST_NOT_FOUND": {
				"hr": "Host nije pronađen",
			},
			"INVALID_DOMAIN": {
				"hr": "Link sadrži neispravnu domenu",
			},
			"REDIRECTOR_DOMAIN": {
				"hr": "Link sadrži domenu preusmjerivača",
			},
			"BLACKLISTED_DOMAIN": {
				"hr": "Link sadrži domenu na crnoj listi",
			},
			"LINK_NOT_VALID": {
				"hr": "Link nije valjan",
			},
			"ERROR_HASHING_PASSWORD": {
				"hr": "Pogreška kod heširanja lozinke",
			},
			"ERROR_CREATING_LINK": {
				"hr": "Pogreška prilikom stvaranja linka",
			},
			"INVALID_PASSWORD": {
				"hr": "Neispravna lozinka",
			},
			"TOO_MANY_REQUESTS": {
				"hr": "Previše zahtjeva",
			},
		},
	}
}

func (es *ErrorsService) Message(identifier string) string {
	langMessages, ok := es.messages[identifier]
	if !ok {
		return "Unknown error"
	}

	if message, ok := langMessages[es.lang]; ok {
		return message
	}

	return "Language not supported"
}
