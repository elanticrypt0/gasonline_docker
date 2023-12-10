package access

import "gorm.io/gorm"

type UserLog struct {
	gorm.Model
	User_uuid string `json:"user"`
	Action    string `json:"action"`
	Url       string `json:"url"`
	Details   string `json:"details"`
}
