package access

import "gorm.io/gorm"

type GroupPerms struct {
	gorm.Model
	Write   bool   `json:"write"`
	Read    bool   `json:"read"`
	Path    string `json:"path"`
	GroupID uint
	Group   Group
}
