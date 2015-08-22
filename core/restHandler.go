package core

import (
	"encoding/json"
	"net/http"
)

type RestHandler struct {
	mottak *Mottak
}

func NyRestHandler() *RestHandler {
	return &RestHandler{mottak: NyttMottak()}
}

func (h RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.ReceiveGet(w, r)
	case "POST":
		h.ReceivePost(w, r)
	case "OPTIONS":
		w.Write([]byte("The handler supports GET and POST!"))
	default:
		http.Error(w, "Method not supported.", 405)
		Error.Println("Metode er ikke støttet: ", r.Method)
	}
}

func (h RestHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eksempel på payload: "))

	fange := IsolatFange{FangeTilIsolat: Fange{Id: "1ES532KD1", Navn: "Albert Åbert"}, IsoleringsTid: 5, CallbackUrl: "http://dummy.url/", Method: "GET"}

	encoder := json.NewEncoder(w)
	encoder.Encode(&fange)
}

func (h RestHandler) ReceivePost(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	fange := IsolatFange{}
	err := decoder.Decode(&fange)

	if err != nil {
		http.Error(w, "Invalid content.", 400)
		Error.Println("Ugyldig innhold.", err.Error())

		return
	}

	h.mottak.Motta(fange)
}
