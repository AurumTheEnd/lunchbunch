package error

import (
	"fmt"
	"log"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, method string) {
	http.Error(w, fmt.Sprintf("Method %s not allowed", method), http.StatusMethodNotAllowed)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func UnauthorizedError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func BadRequestError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Bad request", http.StatusBadRequest)
}
