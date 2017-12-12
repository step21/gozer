package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var onlineOnly bool

func init() {
	flag.BoolVar(&onlineOnly, "online", false, "Show only online network members")
}

func main() {
	flag.Parse()
	if len(apiToken) == 0 {
		var token_buf []byte
		token_buf, err := ioutil.ReadFile(os.ExpandEnv(apiTokenFile))
		if err != nil {
			if !os.IsNotExist(err) {
				log.Fatal("Cannot read", err)
			}
			log.Println("No token file " + apiTokenFile)
		} else {
			apiToken = strings.TrimSpace(string(token_buf))
		}
	}
	if len(apiToken) == 0 {
		log.Fatal("No API token supplied")
	}

	client := ZeroTierClient{}

	network_names := make([]string, 0, 100)

	if flag.NArg() == 0 {
		networks, err := client.ListNetworks(false)
		if err != nil {
			logger.Fatal(err)
		}
		for name := range networks.id_index {
			network_names = append(network_names, name)
		}
	} else {
		for _, name := range flag.Args() {
			network_names = append(network_names, name)
		}
	}
	logger.Verboseln(fmt.Sprintf("Showing detail for: %v", network_names))

	for _, arg := range network_names {

		// Get the information for a named network
		network, err := client.GetNetworkDetails(arg)
		if err != nil {
			fmt.Println(arg, " NOT FOUND ", err)
			continue
		}
		fmt.Println(network.SummaryString())

		// Get a list of members for the network, and iterate
		members := client.GetNetworkMemberDetails(network, onlineOnly)
		for _, member := range members {
			fmt.Println("    ", member.SummaryString())
		}
		fmt.Println()
	}

}
