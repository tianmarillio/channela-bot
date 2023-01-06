package models

type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`

	Timestamp
}
