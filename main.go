package main

import (
	"net/http"

	"github.com/theofficialnar/webservice/controllers"
)

func main() {
	// Initialize all the controllers
	controllers.RegisterControllers()

	// Start the web server and listen to port 3000
	http.ListenAndServe(":3000", nil)
}
