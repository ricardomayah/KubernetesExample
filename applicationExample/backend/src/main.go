package main

import (
    "encoding/json"
    "net/http"
    "os"
    "time"
)

type HandsOn struct{
    Time time.Time  `json:"time"`
    Hostname string `json:"hostname"`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {

    resp := HandsOn{
        Time: time.Now(),
        Hostname: os.Getenv("HOSTNAME"),
    }

    jsonresp, error := json.Marshal(&resp)

    if error!= nil {
        w.Write([]byte("Error"));
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Write(jsonresp)
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}