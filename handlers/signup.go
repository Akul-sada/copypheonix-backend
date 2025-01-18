package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// User represents the user model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// To varify the email is correct
func VarifyEmail() {

}

//

// SignUp handles user registration
func CreateUser(r *http.Request, _ httprouter.Params) *http.Response {
	// Declare the user variable
	var user User
	var err = json.NewDecoder(r.Body).Decode(&user)

	// Create response
	response := &http.Response{
		Header: make(http.Header),
	}

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = io.NopCloser(strings.NewReader(err.Error()))
		return response
	}

	response.Header.Set("Content-Type", "application/json")
	response.StatusCode = http.StatusCreated

	// Create response body
	message, _ := json.Marshal(map[string]string{"message": "User created successfully"})
	response.Body = io.NopCloser(bytes.NewBuffer(message))

	return response
}
