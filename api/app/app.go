package app

import (
	"github.com/nevill/jiangjing/api"
)

type API struct {
	Engines  *Engines
	Synonyms *Synonyms
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
	}
}
