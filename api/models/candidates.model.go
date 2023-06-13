package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Candidate struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Option    int       `gorm:"type:int;not null;unique;" json:"option"`
	Name      string    `gorm:"size:100;not null;" json:"name"`
	Label     string    `gorm:"size:100;not null;" json:"label"`
	Image     string    `gorm:"type:text;not null;" json:"image"`
	Vision    []Vision  `json:"vision"`
	Mission   []Mission `json:"mission"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Candidate) Prepare() {
	p.Id = 0
	p.Option = p.Option
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Label = html.EscapeString(strings.TrimSpace(p.Label))
	p.Image = html.EscapeString(strings.TrimSpace(p.Image))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Candidate) Validate() error {

	// if p.Option == "" {
	// 	return errors.New("Required Option")
	// }
	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Image == "" {
		return errors.New("Required Image")
	}
	return nil
}

func (p *Candidate) SaveCandidate(db *gorm.DB) (*Candidate, error) {
	var err error
	err = db.Debug().Model(&Candidate{}).Create(&p).Error
	if err != nil {
		return &Candidate{}, err
	}
	return p, nil
}

func (p *Candidate) FindAllCandidates(db *gorm.DB) (*[]Candidate, error) {
	var err error
	candidates := []Candidate{}
	err = db.Debug().Model(&Candidate{}).Limit(100).Find(&candidates).Error
	if err != nil {
		return &[]Candidate{}, err
	}
	if len(candidates) > 0 {
		for i, _ := range candidates {
			err := db.Debug().Model(&Mission{}).Where("candidate_id = ?", candidates[i].Id).Find(&candidates[i].Mission).Error
			if err != nil {
				return &[]Candidate{}, err
			}
			err1 := db.Debug().Model(&Vision{}).Where("candidate_id = ?", candidates[i].Id).Find(&candidates[i].Vision).Error
			if err1 != nil {
				return &[]Candidate{}, err1
			}
		}
	}
	return &candidates, nil
}

func (p *Candidate) FindCandidateByID(db *gorm.DB, pid uint64) (*Candidate, error) {
	var err error
	err = db.Debug().Model(&Candidate{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Candidate{}, err
	}
	if p.Id != 0 {
		err := db.Debug().Model(&Mission{}).Where("candidate_id = ?", p.Id).Find(&p.Mission).Error
		if err != nil {
			return &Candidate{}, err
		}
		err1 := db.Debug().Model(&Vision{}).Where("candidate_id = ?", p.Id).Find(&p.Vision).Error
		if err1 != nil {
			return &Candidate{}, err1
		}
	}
	return p, nil
}

func (p *Candidate) UpdateCandidate(db *gorm.DB) (*Candidate, error) {

	var err error
	err = db.Debug().Model(&Candidate{}).Where("id = ?", p.Id).Updates(Candidate{Option: p.Option, Name: p.Name, Label: p.Label, Image: p.Image, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Candidate{}, err
	}
	if p.Id != 0 {
		err := db.Debug().Model(&Mission{}).Where("candidate_id = ?", p.Id).Find(&p.Mission).Error
		if err != nil {
			return &Candidate{}, err
		}
		err1 := db.Debug().Model(&Vision{}).Where("candidate_id = ?", p.Id).Find(&p.Vision).Error
		if err1 != nil {
			return &Candidate{}, err1
		}
	}
	return p, nil
}

func (p *Candidate) DeleteCandidate(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Candidate{}).Where("id = ?", pid).Take(&Candidate{}).Delete(&Candidate{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Candidate not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
