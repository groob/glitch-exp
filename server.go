package main

import (
  "fmt"
  "log"
  "html"
  "net/http"
  "net/http/httputil"
  "github.com/groob/plist"
)

// https://go-example.glitch.me/
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Go!")
}

// https://go-example.glitch.me/love/Go
func love(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", html.EscapeString(r.URL.Path[6:]))
}

// https://go-example.glitch.me/hacking
func hacking(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hacking away on Go on Glitch.")
}

func echo(w http.ResponseWriter, r *http.Request) {
  out, err := httputil.DumpRequest(r, true)
  if err != nil {
    panic(err)
  }
  _ =out
  //w.Write(out)
  var t = struct {Key string}{Key: "foo"}
  plist.NewEncoder(w).Encode(&t)
}

func main() {
    http.HandleFunc("/hacking", hacking)
    http.HandleFunc("/love/", love)
    http.HandleFunc("/", handler)
  http.HandleFunc("/echo", echo)
  log.Fatal(http.ListenAndServe(":3000", nil))
}
