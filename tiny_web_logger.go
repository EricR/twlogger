package main

import (
  "net/http"
  "math/rand"
  "time"
  "strconv"
  "log"
  "flag"
  "fmt"
)

var (
  server_name string
  port int
  listen_address string
)

func init() {
  flag.StringVar(&server_name, "n", random_server_name(), "Name of web server instance.")
  flag.IntVar(&port, "p", 8080, "Port to listen on.")
  flag.Parse()

  listen_address = fmt.Sprintf(":%s", strconv.Itoa(port))
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Fprintf(w, "Hello from %s", server_name)
    log.Printf("%s -> %s (form: %s)", r.RemoteAddr, r.URL.Path, r.Form)
  })

  log.Printf("Starting %s on port %d", server_name, port)

  if err := http.ListenAndServe(listen_address, nil); err != nil {
    log.Fatal("Could not listen and serve on port ", port, "!")
  }
}

func random_server_name() string {
  rand.Seed(time.Now().UTC().UnixNano())
  return fmt.Sprintf("TinyWebServer-%d", rand.Intn(100))
}