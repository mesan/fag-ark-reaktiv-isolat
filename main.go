//fag-ark-reaktiv-isolatfaolatfaolatfaolatfaolatfaolatfaolatfaolatfaolat
package main

import (
	"fag-ark-reaktiv-isolat/core"
	"flag"
	"github.com/goarne/web"
	"net/http"
	"strconv"
)

func main() {
	appConfig := lastAppKonfig()

	core.InitLoggers(appConfig.Logging)

	skrivOppstartsMelding(appConfig)

	http.ListenAndServe(":"+strconv.FormatInt(appConfig.Server.Port, 10), opprettIsolatRessurs(appConfig))
}

//Logger applikasjonskonfigurasjon
func skrivOppstartsMelding(appConfig core.AppConfig) {
	core.Info.Println("Laster appconfig:", appConfig.Logging.Filename)
	core.Info.Println("Starter isolat på port:", strconv.FormatInt(appConfig.Server.Port, 10))
	core.Info.Println("Isolat er tilgjengelig på ", appConfig.Server.Root)
	core.Info.Println("Skriver til loggfil: ", appConfig.Logging.Filename)
}

//Laster inn applikasjonskonfigurasjon fra en JSON konfigurasjonsfil.
func lastAppKonfig() core.AppConfig {

	configFile := flag.String("config", "./config/appconfig.json", "Fullt navn til applikasjonens konfigurasjonsfil (json)")
	flag.Parse()

	appConfig := core.AppConfig{}
	appConfig.ReadConfig(*configFile)

	return appConfig
}

//Oppretter REST ressurs for applikasjon
func opprettIsolatRessurs(appConfig core.AppConfig) *web.WebRouter {
	r := web.NewRoute()

	r.Path(appConfig.Server.Root)
	r.Method(web.HttpGet).Method(web.HttpPost)
	r.Handler(core.NyRestHandler())

	router := web.NewWebRouter()
	router.AddRoute(r)
	return router
}
