package core

import (
	"encoding/json"
	. "github.com/goarne/logging"
	"io/ioutil"
	"net/http"
)

type RestHandler struct {
	mottak *Mottak
}

func NyRestHandler() *RestHandler {
	return &RestHandler{mottak: OpprettMottak()}
}

func (h RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.ReceiveGet(w, r)
	case "POST":
		h.ReceivePost(w, r)
	case "OPTIONS":
		w.Write([]byte("The handler supports GET and POST!"))
		body, _ := ioutil.ReadAll(r.Body)

		w.Write([]byte(body))
	default:

		http.Error(w, "Method not supported.", 405)
		Error.Println("Metode er ikke støttet: ", r.Method)
	}
}

func (h RestHandler) ReceiveGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eksempel på payload: "))
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	fange := IsolatFange{FangeTilIsolat: Fange{Id: "1ES532KD1", Navn: "Albert Åbert"}, IsoleringsTid: 5, CallbackUrl: "http://dummy.url/", Method: "GET", Headers: headers}
	jsonData, _ := json.Marshal(&fange)
	w.Write(jsonData)
}

func (h RestHandler) ReceivePost(w http.ResponseWriter, r *http.Request) {
	//	return
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
