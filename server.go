package main

import (
	"fmt"
	"github.com/Traceableai/goagent"
	"github.com/Traceableai/goagent/config"
	"github.com/Traceableai/goagent/instrumentation/net/traceablehttp"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func startHttpServer() {
	fmt.Println("Starting app on port 8080")
	cfg := config.LoadFromFile("./config.yaml")

	closer := goagent.Init(cfg)
	defer closer()

	r := mux.NewRouter()
	r.Handle("/foo",
		traceablehttp.NewHandler(http.HandlerFunc(fooHandler), "/foo"))
	//	http.HandlerFunc(extCapHandler), "/extcap"))
	//r.Handle("/v1/traces", traceablehttp.NewHandler(
	//	http.HandlerFunc(otlpHttpHandler), "/v1/traces"))
	// Using log.Fatal(http.ListenAndServe(":8081", r)) causes a gosec timeout error.
	// G114 (CWE-676): Use of net/http serve function that has no support for setting timeouts (Confidence: HIGH, Severity: MEDIUM)
	srv := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

type Person struct {
	Name      string   `json:"name"`
	Telephone int      `json:"telephone"`
	City      string   `json:"curr_city"`
	Weather   string   `json:"weather"`
	Nested    []Person `json:"nested"`
}

type Request struct {
	Data []Person `json:"data"`
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	sBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("here")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(r.Body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(sBody)
}
