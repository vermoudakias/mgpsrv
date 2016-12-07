//
// https://blog.bracelab.com/achieving-perfect-ssl-labs-score-with-go
//

package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request) {
    w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
    w.Write([]byte("This is an example server.\n"))
  })
  /* Generate key/cert:
   * go run /usr/local/go/src/crypto/tls/generate_cert.go --host="localhost"
   */
  log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}

