package handlers

import (
	"fmt"
	"net/http"
)

func Uninstall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UNINSTALL!\n")
}
