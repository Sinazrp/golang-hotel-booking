package forms

import (
	// Importing the net/http package for handling HTTP requests and responses.
	"net/http"
	// Importing the net/url package for URL manipulation.
	"net/url"
)

// Definition of the Form struct which embeds url.Values for form data and includes a field for storing validation errors.
type Form struct {
	url.Values
	Errors FormErrors
}

// Constructor function for creating a new instance of the Form struct. It takes url.Values as input and initializes the Errors field with an empty map.
func New(data url.Values) *Form {
	return &Form{
		data,
		FormErrors(map[string][]string{}),
	}
}

// Method on the Form struct to check if a specific field exists in the form data submitted via an HTTP request.
func (f *Form) Has(field string, r *http.Request) bool {
	// Retrieves the value of the specified field from the form data of the HTTP request.
	x := r.Form.Get(field)
	// If the retrieved value is an empty string, it means the field does not exist or its value is empty.
	if x == "" {
		return false
	}
	// If the retrieved value is not an empty string, it means the field exists with a non-empty value.
	return true
}
