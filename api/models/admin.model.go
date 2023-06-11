package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	Username  string    `gorm:"size:100;not null;unique" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func HashAdmin(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPasswordAdmin(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Admin) BeforeSaveAdmin() error {
	hashedPassword, err := HashAdmin(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Admin) PrepareAdmin() {
	u.Id = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Admin) ValidateAdmin(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Username == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}

func (u *Admin) SaveAdmin(db *gorm.DB) (*Admin, error) {

	err := u.BeforeSaveAdmin()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) FindAllAdmins(db *gorm.DB) (*[]Admin, error) {
	var err error
	admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Limit(100).Find(&admins).Error
	if err != nil {
		return &[]Admin{}, err
	}
	return &admins, err
}

func (u *Admin) FindAdminByID(db *gorm.DB, uid uint32) (*Admin, error) {
	var err error
	err = db.Debug().Model(Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("Admin Not Found")
	}
	return u, err
}

func (u *Admin) UpdateAdmin(db *gorm.DB, uid uint32) (*Admin, error) {

	// To hash the password
	err := u.BeforeSaveAdmin()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"name":       u.Name,
			"adminname":  u.Username,
			"password":   u.Password,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Admin{}, db.Error
	}
	// This is the display the updated admin
	err = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) DeleteAdmin(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).Delete(&Admin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
