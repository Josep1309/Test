package main

import (
	f "Server/functions"
	
    "flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type RequestBody struct {
	SearchType   string   `json:"search_type"`
	Query        Query    `json:"query"`
	SortFields   []string `json:"sort_fields"`
	From         int      `json:"from"`
	MaxResults   int      `json:"max_results"`
	SourceFields []string `json:"_source"`
}

type Query struct {
	Term       string `json:"term"`
	Field      string `json:"field"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

    r := chi.NewRouter()

    r.Use(middleware.Logger)

    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
      }))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hi! You are connected to the server."))
    })

    r.Get("/new_indexer", func(w http.ResponseWriter, r *http.Request) {
        var rootPath string = "/home/jose/Documents/Projects/Test/Database/"
        f.EmailsAdd(rootPath)
        f.PostZinc()
        w.Write([]byte("Loaded emails dataset!!!"))
    })

    r.Post("/search/", func(w http.ResponseWriter, r *http.Request) {

        var query RequestBody
        var databaseResponse map[string]map[string]interface{}

        var receivedJson map[string]string
        _ = json.NewDecoder(r.Body).Decode(&receivedJson)

        query.Query.Term = receivedJson["term"]
        query.MaxResults, _ = strconv.Atoi(receivedJson["maxresults"])
        query.SearchType = receivedJson["searchtype"]

        query.Query.StartTime = "2000-06-02T14:28:31.894Z"
        query.Query.EndTime = "2023-12-02T15:28:31.894Z"
        query.From = 0

        jsonquery, _ := json.Marshal(query)

        response := f.Search(string(jsonquery))

        _ = json.Unmarshal(response, &databaseResponse)
        
        response, err := json.Marshal(databaseResponse["hits"]["hits"])
        if err != nil {
            panic(err)
        }

        w.Write([]byte(response))
    })

    http.ListenAndServe(":3000", r)

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}