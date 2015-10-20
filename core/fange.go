package core

//DTO brukt til transport av fange som skal settes i isolat.
type IsolatFange struct {
	FangeTilIsolat Fange             `json:"fangeTilIsolat"`
	IsoleringsTid  int               `json:"isoleringsTid"`
	CallbackUrl    string            `json:"callbackUrl"`
	Method         string            `json:"method"`
	Headers        map[string]string `json:"headers"`
}

//Domeneklasse for en fangerepresentasjon.
type Fange struct {
	Id   string `json:"id"`
	Navn string `json:"navn"`
}

func (f Fange) String() string {
	return "[Fangenr: " + f.Id + ", navn: " + f.Navn + "]"
}
