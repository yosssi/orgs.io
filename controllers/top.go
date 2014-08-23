package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TopIndex is the Top controller's index action.
func TopIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello!")
}
