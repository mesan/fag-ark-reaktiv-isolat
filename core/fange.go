package core

//DTO brukt til transport av fange som skal settes i isolat.
type IsolatFange struct {
	FangeTilIsolat Fange
	IsoleringsTid  int
	CallbackUrl    string
	Method         string
}

//Domeneklasse for en fangerepresentasjon.
type Fange struct {
	Id   string
	Navn string
}

func (f Fange) String() string {
	return "[Fangenr: " + f.Id + ", navn: " + f.Navn + "]"
}
