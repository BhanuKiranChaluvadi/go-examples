package models

// Error Model
// swagger:model error
type Error struct {
	Result int `json:"result"`

	// Enumeration string succinctly identifying the problem
	// example: MetadataInvalid
	Code string `json:"code"`

	// Message is developer-oriented explanation of the solution to the problem
	// example: The `urcapID` field is required
	Message error `json:"message"`

	// A publicly-accessible URL where information about the error can be read in a web browser.
	// example: https://docs.api.example.com/v1/urcap/create_urcap#urcapID
	Details string `json:"details,omitempty"`

	// Error Target :: Pointer to eliminate struct entirely with "omitempty"
	Target *Target `json:"target,omitempty"`
}
