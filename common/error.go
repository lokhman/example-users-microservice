package common

// Error which nicely translates in the HTTP response.
type HTTPError struct {
	Err string `json:"error" example:"Error message"`
}

func (e HTTPError) Error() string {
	return e.Err
}
