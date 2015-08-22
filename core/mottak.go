package core

type Mottak struct {
	innKo chan IsolatFange
}

func NyttMottak() *Mottak {
	m := &Mottak{innKo: make(chan IsolatFange)}
	go m.HandterMottaksKo()
	return m
}

//Motaksmetode som tar i mot en fange og legger den inn i mottakskø.
func (m *Mottak) Motta(f IsolatFange) {
	Trace.Println(f.FangeTilIsolat.String(), "Mottar fange.")
	m.innKo <- f
}

//Metoden håndterer mottakskøen. Den tar mottar fanger, oppretter isolat og starter soningen i en GO routine.
func (m *Mottak) HandterMottaksKo() {
	for f := range m.innKo {
		isolat := &Isolat{fange: f.FangeTilIsolat, isoleringsTid: f.IsoleringsTid, callbackUrl: f.CallbackUrl, method: f.Method}
		go isolat.StartSoning()
	}

	close(m.innKo)
}
