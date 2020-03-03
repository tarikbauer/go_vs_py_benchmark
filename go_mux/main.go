package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	LogWarning(r.URL.String() + " " + r.Method)
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("Not Found!"))
}

func notAllowed(w http.ResponseWriter, r *http.Request) {
	LogWarning(r.URL.String() + " " + r.Method)
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = w.Write([]byte("Method not Allowed!"))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		LogInfo(r.URL.String() + " " + r.Method)
		next.ServeHTTP(w, r)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	t := r.FormValue("t")
	timeList := strings.Split(t, ",")
	c := make(chan int, len(timeList))
	for _, currentTime := range timeList {
		wg.Add(1)
		value, err := strconv.Atoi(currentTime)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Invalid Parameter!"))
			return
		}
		go Sleep(value, c, &wg)
	}
	wg.Wait()
	close(c)
	response := 0
	for value := range c{
		response += value
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strconv.Itoa(response)))
}

func main() {
	port := os.Getenv("GO_MUX_PORT")
	router := mux.NewRouter().StrictSlash(true)
	router.Use(loggingMiddleware)
	router.Path("/api").Queries("t", "{t}").HandlerFunc(Handler).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(notFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(notAllowed)
	LogInfo("Server Listening on 127.0.0.1:" + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}