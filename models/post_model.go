package models

type Post struct {
	Id     int          `gorm:"primaryKey" json:"id"`
	Title  string       `json:"title" gorm:"not null"`
	Body   string       `json:"body" gorm:"not null"`
	UserId int          `json:"user_id"`
	User   UserResponse `json:"user"`
}

type PostResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"-"`
}

func (PostResponse) TableName() string {
	return "posts"
}