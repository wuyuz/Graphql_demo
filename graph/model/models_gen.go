// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `json:"user" pg:"-"`
	UserID      string `json:"userId"`
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups" pg:"-"`
}
