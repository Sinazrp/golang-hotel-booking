package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"

	// Importing the net/http package for handling HTTP requests and responses.
	"net/http"
	"strings"
	"unicode/utf8"

	// Importing the net/url package for URL manipulation.
	"net/url"
)

func (f *Form) Valid() bool {
	return len(f.Errors) == 0

}

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
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	// If the retrieved value is not an empty string, it means the field exists with a non-empty value.
	return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}

}
func (f *Form) MaxLength(field string, length int) bool {

	if utf8.RuneCountInString(f.Get(field)) > length {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", length))
		return false
	}
	return true

}
func (f *Form) MinLength(field string, length int) bool {

	if utf8.RuneCountInString(f.Get(field)) < length {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", length))
		return false
	}
	return true

}
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}

}
