package main

import (
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) { 
        _, err := w.Write([]byte(
`
<html>
    <body>
        Sorry, we are currently down for maintenance. We will be back shortly.
    </body>
</html>
`,
        ))
        if err != nil {
            http.Error(w, "unexpected error, please try again", 500)
            log.Printf("error, when attempting to write response page. Error: %v", err)
            return
        }
    })
    http.HandleFunc("/health", func (w http.ResponseWriter, r *http.Request) {})
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatalf("error, when starting http server. Error: %v", err)
    }
}
