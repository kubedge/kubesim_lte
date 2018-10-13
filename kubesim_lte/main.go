package main
import (
  "log"
  "os"
  "net/http"
  "strings"
  "github.com/kubedge/kubesim_base/grpc/go/kubedge_client"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

//arguments
// arg1=demotype
// arg2=demovalue
func main() {
  log.Printf("%s", "kubesim_lte client is running")
  demotype := os.Args[1]
  demovalue := os.Args[2]
  log.Printf("demotype=%s, demovalue=%s", demotype, demovalue)
  go client.Client(demotype, demovalue)

  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
  log.Printf("%s", "kubesim_lte client is exiting")
}
