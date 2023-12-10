package access

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Shadow struct {
	gorm.Model
	Salt   string    // the salt
	Shadow string    `gorm:"type:text;index"` // encrypted password
	UserID uuid.UUID `gorm:"type:uuid;index"`
	User   User
}

func NewShadow() *Shadow {
	return &Shadow{}
}

func (s *Shadow) FindOne(db *gorm.DB, userID uuid.UUID) *Shadow {
	var Shadow Shadow
	db.Preload("Group").First(&Shadow, userID)
	return &Shadow
}

func (s *Shadow) FindAll(db *gorm.DB) []Shadow {
	var Shadows []Shadow

	db.Preload("Group").Order("created_at ASC").Find(&Shadows)
	return Shadows
}

func (s *Shadow) Create(db *gorm.DB, userID uuid.UUID, password string) *Shadow {

	Shadow := Shadow{
		Shadow: strings.ToLower(password),
		UserID: userID,
	}
	db.Create(&Shadow)
	return &Shadow
}

func (s *Shadow) Update(db *gorm.DB, userID uuid.UUID, password string) *Shadow {
	ShadowUpdated := s.FindOne(db, userID)
	passwordEnc := s.EncryptPassword(password)
	if passwordEnc != "" {
		ShadowUpdated.Shadow = s.EncryptPassword(password)
		db.Save(&ShadowUpdated)
	}
	return ShadowUpdated
}

func (s *Shadow) Delete(db *gorm.DB, id int) *Shadow {
	var Shadow Shadow
	db.First(&Shadow, id)
	db.Delete(&Shadow)
	return &Shadow
}

func (s *Shadow) EncryptPassword(password string) string {
	passwordEnc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		return string(passwordEnc)
	}
}
