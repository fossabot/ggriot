package ggriot

// JSONError is the struct that Riot returns the error codes.
// This gives us more information that just getting the HTTP Status
type JSONError struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	} `json:"status"`
}
