package model

import "github.com/yguilai/timetable-micro/common"

type User struct {
	common.Model
	Email    string `json:"email" gorm:"uniqueIndex"`                  // 邮箱, 唯一
	Nickname string `json:"nickname" gorm:"index:idx_nickname,unique"` // 昵称, 唯一
	Avatar   string `json:"avatar" gorm:"default:default_avatar.png"`
	OwnWord  string `json:"ownWord" gorm:"default:这个人很懒, 什么都没留下"`
	Gender   int    `json:"gender" gorm:"default:0"` // 性别: 0: unknown, 1: male, 2: female
	State    int    `json:"state" gorm:"default:1"`                                                                   // 状态: 0: disable, 1: normal
}

type UserAuth struct {
	common.Model
	UserID   uint   `json:"userId"` // 关联用户
	Email    string `json:"email" gorm:"uniqueIndex"` // 登录邮箱
	Password string `json:"password"`                 // 密码, 使用bcrypt加密
}
