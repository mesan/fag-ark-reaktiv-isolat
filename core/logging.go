package core

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLoggers(logConfig LogConfig) {
	logFil := opprettLoggFil(logConfig)

	initLoggers(logFil, logFil, logFil, logFil)
}

func opprettLoggFil(config LogConfig) io.Writer {
	file, err := os.OpenFile(config.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Greide ikke åpne loggfil.", config.Filename, ":", err)
	}

	return io.MultiWriter(file, os.Stdout)
}

func initLoggers(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
