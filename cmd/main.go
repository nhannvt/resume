package main

import (
	"flag"
	"github.com/nhannvt/resume/configs"
	"github.com/nhannvt/resume/internal/interface"
	"net/http"
	"strings"
	"time"
)

const timezone = "Asia/Ho_Chi_Minh"

func init() {
	// Set local time zone
	location, err := time.LoadLocation(timezone)
	if err != nil {
		location = time.FixedZone(timezone, 9*60*60)
	}
	time.Local = location

	// Set number of max idle connections for using keepAlive more efficiently
	http.DefaultTransport.(*http.Transport).MaxIdleConns = 3000
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
}

func main() {

	config := configs.GetConfig()

	port := flag.String("port", config.Get("defaultServerPort"), "API Server Port")
	without := flag.String("without", "", "Disabled Middlewares. For example \"authenticate,validate\"")
	stage := flag.String("stage", config.Get("stage"), "SForum API Stage Name")
	debug := flag.Bool("debug", false, "Enabled Debug mode")
	terminationPeriod := flag.Int64("termination-period", config.GetInt64("defaultTerminationPeriod"), "Termination graceful period seconds")

	flag.Parse()

	if *debug {
		*terminationPeriod = int64(1)
	}

	options := &api.Options{
		Port:              *port,
		Without:           strings.Split(*without, ","),
		Debug:             *debug,
		Stage:             *stage,
		TerminationPeriod: *terminationPeriod,
	}

	server := api.NewServer(options)
	server.Start()
}
