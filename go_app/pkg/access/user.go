package access

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email     string     `gorm:"type:varchar(100);not null; unique"`
	Groups    []Group
	Verified  bool           `gorm:"not null;default:false"`
	CreatedAt *time.Time     `gorm:"not null;default:now"`
	UpdatedAt *time.Time     `gorm:"not null;default:now"`
	DeletedAt gorm.DeletedAt `gorm:"not null;index"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) FindOne(db *gorm.DB, id uuid.UUID) *User {
	var user User
	db.Preload("Group").First(&user, id)
	return &user
}

func (u *User) FindAll(db *gorm.DB) []User {
	var users []User

	db.Preload("Group").Order("created_at ASC").Find(&users)
	return users
}

func (u *User) Create(db *gorm.DB, email, password string, group uint) *User {

	s := NewShadow()
	groups := u.AddGroup(db, group)

	user := User{
		Email:  strings.ToLower(email),
		Groups: *groups,
	}
	db.Create(&user)
	s.Create(db, *user.ID, password)

	return &user
}

func (u *User) Update(db *gorm.DB, user User) *User {

	userUpdated := u.FindOne(db, *user.ID)
	userUpdated.Email = user.Email
	db.Save(&userUpdated)
	return userUpdated
}

func (u *User) Delete(db *gorm.DB, id int) *User {
	var user User
	db.First(&user, id)
	db.Delete(&user)
	return &user
}

func (u *User) AddGroup(db *gorm.DB, groupID uint) *[]Group {
	g := NewGroup()
	groups := g.FindUsersGroups(db, *u.ID)
	var groups2Return []Group
	if !g.ContainsGroup(groups, groupID) {
		groups2Return = g.AddGroup(groups, groupID)
	}
	return &groups2Return
}

func (u *User) VerifyUser(db *gorm.DB, userID uuid.UUID) bool {
	user := u.FindOne(db, userID)
	user.Verified = true
	db.Save(&user)
	return user.Verified
}

func (u *User) Login(db *gorm.DB, email, password string) *User {
	var user *User
	db.Select("u.*").Table("users u").InnerJoins("shadow s").Where("u.email=? and s.shadow=?", email, password).Find(&user)
	return user
}
