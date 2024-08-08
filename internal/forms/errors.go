package forms

// Type definition for FormErrors, which is a map where keys are strings representing field names and values are slices of strings representing error messages.
type FormErrors map[string][]string

// Method on the FormErrors type to add an error message to a specific field. It appends the error message to the slice of errors for that field.
func (fe FormErrors) Add(field string, err string) {
	fe[field] = append(fe[field], err)
}

// Method on the FormErrors type to retrieve the first error message for a specific field. If there are no errors for the field, it returns an empty string.
func (fe FormErrors) Get(field string) string {
	es := fe[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
