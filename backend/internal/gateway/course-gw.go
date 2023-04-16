package gateway

import "time"

type GetCourses struct {
	Limit int `json:"limit"`
}

type CreateCourse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Left        int       `json:"left"`
	Expiration  time.Time `json:"expiration"`
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Timetable   string    `json:"timetable"`
}

type GetCourseById struct {
	Id uint `json:"id"`
}
