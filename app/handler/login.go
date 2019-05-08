package handler

import(
	"net/http"
	"strings"
)

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		
	}
}