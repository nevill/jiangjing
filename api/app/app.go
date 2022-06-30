package app

import (
	"github.com/nevill/jiangjing/api"
)

type API struct {
	Engines   *Engines
	Synonyms  *Synonyms
	Documents *Documents
	Search    Search
}

func New(t api.Transport) *API {
	return &API{
		Engines: &Engines{
			List:   newEnginesListFunc(t),
			Get:    newEnginesGetFunc(t),
			Create: newEngineCreateFunc(t),
			Delete: newEngineDeleteFunc(t),
		},
		Synonyms: &Synonyms{
			List:   newSynonymsListFunc(t),
			Get:    newSynonymsGetFunc(t),
			Create: newSynonymsCreateFunc(t),
			Update: newSynonymsUpdateFunc(t),
			Delete: newSynonymsDeleteFunc(t),
		},
		Documents: &Documents{
			Create: newDocumentsCreateFunc(t),
			Delete: newDocumentsDeleteFunc(t),
			List:   newDocumentsListFunc(t),
		},
		Search: newSearchFunc(t),
	}
}
