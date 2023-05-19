package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Class struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Class) Prepare() {
	p.Id = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Class) Validate() error {

	// if p.Option == "" {
	// 	return errors.New("Required Option")
	// }
	if p.Name == "" {
		return errors.New("Required Name")
	}
	return nil
}

func (p *Class) SaveClass(db *gorm.DB) (*Class, error) {
	var err error
	err = db.Debug().Model(&Class{}).Create(&p).Error
	if err != nil {
		return &Class{}, err
	}
	return p, nil
}

func (p *Class) FindAllClasses(db *gorm.DB) (*[]Class, error) {
	var err error
	Classes := []Class{}
	err = db.Debug().Model(&Class{}).Limit(100).Find(&Classes).Error
	if err != nil {
		return &[]Class{}, err
	}
	return &Classes, nil
}

func (p *Class) FindClassByID(db *gorm.DB, pid uint64) (*Class, error) {
	var err error
	err = db.Debug().Model(&Class{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Class{}, err
	}
	return p, nil
}

func (p *Class) UpdateClass(db *gorm.DB) (*Class, error) {

	var err error
	err = db.Debug().Model(&Class{}).Where("id = ?", p.Id).Updates(Class{Name: p.Name, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Class{}, err
	}
	return p, nil
}

func (p *Class) DeleteClass(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Class{}).Where("id = ?", pid).Take(&Class{}).Delete(&Class{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Class not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
