package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Participant struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	RegNumber string `gorm:"size:100;not null;unique" json:"regnumber"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func HashParticipant(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPasswordParticipant(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *Participant) BeforeSaveParticipant() error {
	hashedPassword, err := HashParticipant(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *Participant) PrepareParticipant() {
	u.Id = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.RegNumber = html.EscapeString(strings.TrimSpace(u.RegNumber))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
}

func (u *Participant) ValidateParticipant(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.RegNumber == "" {
			return errors.New("Required RegNumber")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.RegNumber == "" {
			return errors.New("Required RegNumber")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.RegNumber == "" {
			return errors.New("Required RegNumber")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}

func (u *Participant) SaveParticipant(db *gorm.DB) (*Participant, error) {

	err := u.BeforeSaveParticipant()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Participant{}, err
	}
	return u, nil
}

func (u *Participant) FindAllParticipants(db *gorm.DB) (*[]Participant, error) {
	var err error
	participants := []Participant{}
	err = db.Debug().Model(&Participant{}).Limit(100).Find(&participants).Error
	if err != nil {
		return &[]Participant{}, err
	}
	return &participants, err
}

func (u *Participant) FindParticipantByID(db *gorm.DB, uid string) (*Participant, error) {
	var err error
	err = db.Debug().Model(Participant{}).Where("reg_number = ?", uid).Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Participant{}, errors.New("Participant Not Found")
	}
	if err != nil {
		return &Participant{}, err
	}
	return u, err
}

func (u *Participant) UpdateParticipant(db *gorm.DB, uid string) (*Participant, error) {

	// To hash the password
	err := u.BeforeSaveParticipant()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Participant{}).Where("reg_number = ?", uid).Take(&Participant{}).UpdateColumns(
		map[string]interface{}{
			"name":       u.Name,
			"reg_number": u.RegNumber,
			"password":   u.Password,
		},
	)
	if db.Error != nil {
		return &Participant{}, db.Error
	}
	// This is the display the updated participant
	err = db.Debug().Model(&Participant{}).Where("reg_number = ?", uid).Take(&u).Error
	if err != nil {
		return &Participant{}, err
	}
	return u, nil
}

func (u *Participant) DeleteParticipant(db *gorm.DB, uid string) (int64, error) {

	db = db.Debug().Model(&Participant{}).Where("reg_number = ?", uid).Take(&Participant{}).Delete(&Participant{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Participant not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
