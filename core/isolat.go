package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Isolat struct {
	fange         Fange
	isoleringsTid int
	callbackUrl   string
	method        string
}

func (i *Isolat) StartSoning() {
	Trace.Println(i.fange.String(), "Fange starter isolat.")
	time.Sleep(time.Duration(i.isoleringsTid) * time.Second)

	if err := i.AvsluttSoning(); err != nil {
		Error.Println(i.fange.String(), "Fange ble ikke løslat.", err)
		return
	}

	Trace.Println(i.fange.String(), "Fange slippes fri fra isolat.")
}

func (i *Isolat) AvsluttSoning() error {
	fangeBytes, _ := json.Marshal(i.fange)

	client := &http.Client{}
	req, _ := http.NewRequest(i.method, i.callbackUrl, bytes.NewReader(fangeBytes))
	resp, e := client.Do(req)

	if e != nil {
		return errors.New("Det oppstod en feil." + e.Error())
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("Feil svar fra server." + resp.Status)
	}

	return nil
}
