package userDto

type UserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Token    string `json:"token"`
}
type CheckAuthResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	Username string `json:"username"`
}
