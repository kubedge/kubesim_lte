package main
import (
  "log"
  "os"
  "net/http"
  "strings"
  "time"
  "github.com/kubedge/kubesim_base/grpc/go/kubedge_client"
  "github.com/kubedge/kubesim_base/config"
  "github.com/kubedge/kubesim_base/connected"
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

  var conf config.Configdata
  conf.Config()
  log.Printf("lte client:  product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
             conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)


  var conn connected.Connecteddata
  var message string

  log.Printf("starting for loop")
  for {
     conn.Readconnectvalues()
     log.Printf("epc server:  connected=%s", conn.Connected)

     message = strings.Join(conn.Connected, ",") 
     log.Printf("message to send=" + message)
     client.Client(demotype, message)

     time.Sleep(15 * time.Second)  //every 15 seconds
  }

  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
  log.Printf("%s", "kubesim_lte client is exiting")
}
