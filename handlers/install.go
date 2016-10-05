package handlers

import (
	"fmt"
	"net/http"
)

func Install(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "INSTALL!\n")
}
