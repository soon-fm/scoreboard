package db

// Implements the Point interface to store scores consistently
type Score struct {
	value int
	user  string
}

// Returns the measurement name
func (s Score) Name() string {
	return "scores"
}

// Returns the Tags for this score
func (s Score) Tags() map[string]string {
	return map[string]string{
		"user": s.user,
	}
}

// Returns the field values for the score
func (s Score) Fields() map[string]interface{} {
	return map[string]interface{}{
		"value": s.value,
	}
}

// Returns the value
func (s Score) Value() int {
	return s.value
}

// Returns the user id
func (s Score) User() string {
	return s.user
}

// Constructs a new score
func NewScore(user string, val int) Score {
	return Score{user: user, value: val}
}
