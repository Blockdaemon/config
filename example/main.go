package main

import "fmt"
import "github.com/Blockdaemon/config"

func main() {
	config := new(config.Config)
	config.SetPrefix("EXAMPLE_")

	config.DescribeOptionalInt("PORT", "The port to listen to", 1234)
	config.DescribeOptionalString("HOST", "The host to listen to", "0.0.0.0")
	config.DescribeOptionalBool("DEBUG", "Start in debug mode", false)

	config.DescribeMandatoryInt("PEERS", "The amount of peers to connect to")
	config.DescribeMandatoryString("ID", "A unique id for this app")
	config.DescribeMandatoryBool("ACCEPT", "Accept the terms and conditions")

	config.Parse()

	port := config.GetInt("PORT")
	host := config.GetString("HOST")
	debug := config.GetBool("DEBUG")
	peers := config.GetInt("PEERS")
	id := config.GetString("ID")
	accept := config.GetBool("ACCEPT")

	fmt.Printf("Port: %d, Host: %s, Debug: %t, Peers: %d, Id: %s, Accept: %t\n", port, host, debug, peers, id, accept)
}
