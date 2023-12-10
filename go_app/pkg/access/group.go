package access

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID        uint   `gorm:"type:integer;not null;autoIncrement;primary_key"`
	Name      string `gorm:"unique" json:"name"`
	UserID    uuid.UUID
	User      User
	CreatedAt *time.Time     `gorm:"not null;default:now"`
	UpdatedAt *time.Time     `gorm:"not null;default:now"`
	DeletedAt gorm.DeletedAt `gorm:"not null;index"`
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) FindOne(db *gorm.DB, id int) *Group {
	var user Group
	db.First(&user, id)
	return &user
}

func (g *Group) FindAll(db *gorm.DB) []Group {
	var groups []Group
	db.Order("created_at ASC").Find(&groups)
	return groups
}

func (g *Group) Create(db *gorm.DB, gName string) *Group {
	group := &Group{
		Name: strings.ToUpper(gName),
	}
	db.Create(&group)
	return group
}

func (g *Group) Update(db *gorm.DB, group Group) *Group {
	db.Save(&group)
	return &group
}

func (g *Group) Delete(db *gorm.DB, id int) *Group {
	var group Group
	db.First(&group, id)
	db.Delete(&group)
	return &group
}

func (g *Group) FindUsersGroups(db *gorm.DB, userID uuid.UUID) []Group {
	groups := []Group{}
	db.Table("Groups").Where("user_id=?", userID).Find(&groups)
	return groups
}

func (g *Group) ContainsGroup(groups []Group, groupID uint) bool {

	for _, a := range groups {
		if a.ID == groupID {
			return true
		}
	}
	return false
}

func (g *Group) AddGroup(groups []Group, groupID uint) []Group {
	group := Group{
		ID: groupID,
	}
	groups = append(groups, group)

	return groups

}
