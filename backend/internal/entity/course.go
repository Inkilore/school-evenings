package entity

import "time"

type Course struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Left        int       `json:"left"`
	Archive     bool      `json:"archive"`
	ExpiresOn   time.Time `json:"expires"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Timetable   string    `json:"timetable"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewCourse(title, description string, left int, expiration, from, to time.Time, timetable string) *Course {
	return &Course{
		Title:       title,
		Description: description,
		Left:        left,
		ExpiresOn:   expiration,
		From:        from,
		To:          to,
		Timetable:   timetable,
		CreatedAt:   time.Now(),
	}
}
