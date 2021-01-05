// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewSchedule struct {
	UserID  *int    `json:"userId"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Memo    *string `json:"memo"`
	StartAt *string `json:"start_at"`
	EndAt   *string `json:"end_at"`
}

type NewUser struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

type Schedule struct {
	ID      *int    `json:"id"`
	UserID  *int    `json:"userId"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Memo    *string `json:"memo"`
	StartAt *string `json:"start_at"`
	EndAt   *string `json:"end_at"`
}

type User struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}
