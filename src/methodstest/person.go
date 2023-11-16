package methodstest

import (
	"fmt"
	"time"
)

type PersonTwo struct {
	First       string
	Last        string
	DateOfBirth time.Time
}

func NewPerson(first, last string, year, month, day int) *PersonTwo {
	return &PersonTwo{
		First:       first,
		Last:        last,
		DateOfBirth: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local),
	}
}

func (p PersonTwo) GetAge() int {
	timeSince := time.Since(p.DateOfBirth)
	return int(timeSince.Hours() / 24 / 365)
}

func (p PersonTwo) SayHello() string {
	return fmt.Sprintf("Hello %s\n", p.First)
}
