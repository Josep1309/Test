package functions

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)


func Search(query string) []byte{
    req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_search", strings.NewReader(query))
    if err != nil {
        log.Fatal(err)
    }
    req.SetBasicAuth("admin", "Complexpass#123")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    log.Println(resp.StatusCode)
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    return body
}

func PostZinc(){
    data, err := os.Open("/home/jose/Documents/Projects/Test/Backend/emails.ndjson")
    fmt.Println(data)
    if err != nil {
        log.Fatal(err)
    }
    req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", data)
    if err != nil {
        log.Fatal(err)
    }
    req.SetBasicAuth("admin", "Complexpass#123")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    log.Println(resp.StatusCode)
    body , err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(body)
}