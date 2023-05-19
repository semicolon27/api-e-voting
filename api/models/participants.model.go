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

type Participant struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:100;not null;" json:"name"`
	RegNumber string `gorm:"size:100;not null;unique" json:"regnumber"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	ClassId   int    `sql:"type:int REFERENCES classes(id)" json:"classid"`
	Class     Class  `json:"class"`
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
	u.Class = Class{}
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
		if u.ClassId < 1 {
			return errors.New("Required ClassId")
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
		if u.ClassId < 1 {
			return errors.New("Required ClassId")
		}
		return nil
	}
}

func (u *Participant) SaveParticipant(db *gorm.DB) (*Participant, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Participant{}, err
	}
	if u.Id != 0 {
		err = db.Debug().Model(&Class{}).Where("id = ?", u.ClassId).Take(&u.Class).Error
		if err != nil {
			return &Participant{}, err
		}
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
	if len(participants) > 0 {
		for i, _ := range participants {
			err := db.Debug().Model(&Class{}).Where("id = ?", participants[i].ClassId).Take(&participants[i].Class).Error
			if err != nil {
				return &[]Participant{}, err
			}
		}
	}
	return &participants, err
}

func (u *Participant) FindParticipantByID(db *gorm.DB, uid uint32) (*Participant, error) {
	var err error
	err = db.Debug().Model(Participant{}).Where("id = ?", uid).Take(&u).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Participant{}, errors.New("Participant Not Found")
	}
	if err != nil {
		return &Participant{}, err
	}
	if u.Id != 0 {
		err = db.Debug().Model(&Class{}).Where("id = ?", u.ClassId).Take(&u.Class).Error
		if err != nil {
			return &Participant{}, err
		}
	}
	return u, err
}

func (u *Participant) UpdateParticipant(db *gorm.DB, uid uint32) (*Participant, error) {

	// To hash the password
	err := u.BeforeSaveParticipant()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Participant{}).Where("id = ?", uid).Take(&Participant{}).UpdateColumns(
		map[string]interface{}{
			"name":       u.Name,
			"regnumber":  u.RegNumber,
			"password":   u.Password,
			"classid":    u.ClassId,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Participant{}, db.Error
	}
	// This is the display the updated participant
	err = db.Debug().Model(&Participant{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Participant{}, err
	}
	if u.Id != 0 {
		err = db.Debug().Model(&Class{}).Where("id = ?", u.ClassId).Take(&u.Class).Error
		if err != nil {
			return &Participant{}, err
		}
	}
	return u, nil
}

func (u *Participant) DeleteParticipant(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Participant{}).Where("id = ?", uid).Take(&Participant{}).Delete(&Participant{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
