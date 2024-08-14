package search

// ExpYears indicates the range of experience years.
// The format is like [start, end].
// start cannot be less than 0.
// But end can have negative value; Indicating that it has no limit.
// start should always be less or equal to end.
type ExpYears [2]int

func (e ExpYears) Start() int       { return e[0] }
func (e ExpYears) End() int         { return e[1] }
func (e ExpYears) HasNoLimit() bool { return e[0] == 0 && e[1] < 0 }
func (e ExpYears) Valid() bool      { return e[0] >= 0 && e[1] >= e[0] }
