package access

import (
	"gorm.io/gorm"
)

// Install the access pkg
func SetupAccess(db *gorm.DB) {
	AutoMigrate(db)
	SeedGroups(db)
	SeedGroupsPerms(db)
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Shadow{}, &Group{}, &GroupPerms{}, &AuthFactor{}, &User{}, &UserLog{})
}
