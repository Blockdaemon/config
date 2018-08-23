package main

import "fmt"
import "github.com/Blockdaemon/config"

func main() {
	config := new(config.Config)
	config.SetPrefix("EXAMPLE_")
	config.DescribeInt("PORT", "The port to listen to", false, 1234)
	config.DescribeString("HOST", "The host to listen to", false, "0.0.0.0")
	config.DescribeBool("DEBUG", "Start in debug mode", false, true)
	config.Parse()

	host := config.GetString("HOST")
	port := config.GetInt("PORT")
	debug := config.GetBool("DEBUG")

	fmt.Printf("Host: %s, Port: %d, Debug: %t\n", host, port, debug)
}
