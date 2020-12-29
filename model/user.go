package model

type User struct {
	Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Nickname string `json:"nickname" gorm:"index:idx_nickname,unique"`
	Gender   int    `json:"gender"`
	State    int    `json:"state"`
}

type UserAuth struct {
	Model
	UserID   uint   `json:"userId"`
	User     User   `json:"user"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
