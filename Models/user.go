package Models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
}
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (UserResponse) TableName() string {
	return "users"
}
