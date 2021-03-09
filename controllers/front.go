package controllers

import (
	"net/http"
)

// RegisterControllers is used to register new routes & determine what controller handles which route.
func RegisterControllers() {
	// Initialize a new instance of the userController
	uc := newUserController()

	/* Let the program know that when there are requests for the urls
	 * below, it should route it to the appropriate controller.
	 *
	 * *Note: In golang, /users and /users/ are different so we need to handle both.
	 *
	 * **Note: In the case of /users/ anything after the declared still is still
	 * considered valid e.g. /users/5 or /users/any/kind/of/path/
	 */
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
