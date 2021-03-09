package controllers

import (
	"net/http"
	"regexp"
)

/* Works just like a class.
 *
 * *Note: userIDPattern will be initialized with a regexp to help the
 * controller know that the request is to get data for a specific user.
 *
 * **Note: userController automatically implements the Handler interface
 * because of the ServeHTTP method.
 */
type userController struct {
	userIDPattern *regexp.Regexp
}

// ServeHTTP Method that will send out the response
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

/* Constructor for the userController class.
 *
 * Initializes the controller with the regex to match urls like /user/123123
 */
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
