package main

import (
	"./server"
	"flag"
	"fmt"
	"strconv"
)

// available topics
var availableTopics = map[string]string{
	"PFM": "PLATEFORMFILE_MAKER",
	"DFM": "DOCKERFILE_MAKER",
	"DPY": "DEPLOY",
	"GFI": "GETCAASFILE_INTEGRITY",
}

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {

	flag.Parse()
	pbserv := server.NewPubSubServer()
	pbserv.Start(strconv.Itoa(*port))
	fmt.Println("Serveur lancé !")
}
