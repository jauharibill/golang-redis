package models

type ProfileModel struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DateBirth string `json:"date_birth"`
	UserID    int64  `json:"user_id"`
}
