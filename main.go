package main

import (
    "net/http"
    "errors"
    "fmt"
    "log"
)

func main() {
    serviceName := "maint"
    config := Nats{
        Host: "127.0.0.1",
        Port: 3000,
    }
    healthy, err := New(config, serviceName, healthStatusRefresh)
    if err != nil {
        log.Fatalf("error, when attempting to connect to nats. Error: %v", err)
    }
    defer healthy.Close()
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) { 
        err = errors.New("error, maintenance page is being hit. Check the health of all your services")
        healthy.ReportUnexpectedError(w, err)
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
            err = fmt.Errorf("error, when attempting to write response page. Error: %v", err)
            http.Error(w, "unexpected error, please try again", http.StatusInternalServerError)
            healthy.ReportUnexpectedError(w, err)
            return
        }
    })
    http.HandleFunc("/health", func (w http.ResponseWriter, r *http.Request) {})
    err = http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatalf("error, when starting http server. Error: %v", err)
    }
}

func healthStatusRefresh(_ HealthStatus) error {
    // nothing to check yet
    return nil
}
