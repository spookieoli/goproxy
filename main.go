package main

// Importing the package
import (
	"goproxy/goproxy"
)

// main Function to start the proxy
func main() {
	// Proxy instance
	proxy := goproxy.NewProxy()
	// Start the proxy
	proxy.Serve()
}
