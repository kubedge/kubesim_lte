/*
Copyright 2018 Kubedge

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/kubedge/kubesim_base/config"
	"github.com/kubedge/kubesim_base/connected"
	"github.com/kubedge/kubesim_base/grpc/go/kubedge_client"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const SIM_NAME = "KUBESIM LTE"
const SIM_CONFIG_FILE = "/etc/kubedge/kubesim_conf.yaml"
const SIM_CONNECTED_UE_FILE = "/etc/kubedge/connected_ue.yaml"
const SIMPLE_HTTP_SERVER = false

func sim_message(msg string) {
	log.Printf("%s: %s", SIM_NAME, msg)
}

func configAPI(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = SIM_NAME + " : " + message
	w.Write([]byte(message))
}

func main() {
	sim_message("Starting")

	demotype := os.Args[1]
	log.Printf("%s: demotype=%s", SIM_NAME, demotype)

	var conf config.Configdata
	conf.Config(SIM_NAME, SIM_CONFIG_FILE)
	log.Printf("%s: product_name=%s, product_type=%s, product_family=%s, product_release=%s, feature_set1=%s, feature_set2=%s",
		SIM_NAME, conf.Product_name, conf.Product_type, conf.Product_family, conf.Product_release, conf.Feature_set1, conf.Feature_set2)

	var conn connected.Connecteddata

	if !SIMPLE_HTTP_SERVER {
		sim_message("Monitoring UE Connection")
		for {
			conn.Readconnectvalues(SIM_NAME, SIM_CONNECTED_UE_FILE)
			log.Printf("%s: Connected UEs: %s", SIM_NAME, conn.Connected)

			imsi := strings.Join(conn.Connected, ",")
                        for index,message := range conf.Feature_set1 {
			    log.Printf("%s: Message %s to send: %s %s", SIM_NAME, index, message, imsi)
			    client.Client(SIM_NAME, message, imsi)
                        }

			time.Sleep(15 * time.Second) //every 15 seconds
		}
	} else {
		// Simple HTTP Server instead of endless loop
		http.HandleFunc("/", configAPI)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}
	sim_message("Exiting")
}
