package models

// swagger:model target
type Target struct {
	// Enumeration of `field`, `parameter`, `header`
	// example: field
	Type string `json:"type"`

	// The name of the problematic field
	// example: urcapID
	Name string `json:"name"`
}
