package search

import (
	"strconv"
)

// ExpYears indicates the range of experience years.
// The format is like [start, end].
// start cannot be less than 0.
// But end can have negative value; Indicating that it has no limit.
// start should always be less or equal to end.
type ExpYears [2]int

func (e ExpYears) Start() int          { return e[0] }
func (e ExpYears) End() int            { return e[1] }
func (e ExpYears) HasNoLimit() bool    { return e[0] == 0 && e.HasNoEndLimit() }
func (e ExpYears) HasNoEndLimit() bool { return e[1] < 0 }
func (e ExpYears) Valid() bool         { return e[0] >= 0 && e[1] >= e[0] }

func (e *ExpYears) SetNoEndLimit() { e[1] = -1 }

func (e ExpYears) String() string {
	if e.HasNoLimit() {
		return "경력무관"
	}

	from := "신입"
	if e.Start() > 0 {
		from = strconv.Itoa(e.Start()) + "년"
	}

	if e.HasNoEndLimit() {
		return from + " 이상"
	}

	if e.Start() == e.End() {
		return from
	}

	to := strconv.Itoa(e.End()) + "년"

	return from + "~" + to
}
