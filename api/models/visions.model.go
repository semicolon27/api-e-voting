package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Vision struct {
	Id          int       `gorm:"primary_key;auto_increment" json:"id"`
	CandidateId int       `sql:"type:int REFERENCES candidates(id)" json:"candidateid"`
	Candidate   Candidate `json:"candidate"`
	Vision      string    `gorm:"size:255;not null" json:"vision"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Vision) Prepare() {
	p.Id = 0
	p.Candidate = Candidate{}
	p.Vision = html.EscapeString(strings.TrimSpace(p.Vision))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Vision) Validate() error {

	// if p.Option == "" {
	// 	return errors.New("Required Option")
	// }
	if p.Vision == "" {
		return errors.New("Required Vision")
	}
	return nil
}

func (p *Vision) SaveVision(db *gorm.DB) (*Vision, error) {
	var err error
	err = db.Debug().Model(&Vision{}).Create(&p).Error
	if err != nil {
		return &Vision{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vision{}, err
		}
	}
	return p, nil
}

func (p *Vision) FindAllVisions(db *gorm.DB) (*[]Vision, error) {
	var err error
	missions := []Vision{}
	err = db.Debug().Model(&Vision{}).Limit(100).Find(&missions).Error
	if err != nil {
		return &[]Vision{}, err
	}
	if len(missions) > 0 {
		for i, _ := range missions {
			err := db.Debug().Model(&Candidate{}).Where("id = ?", missions[i].CandidateId).Take(&missions[i].Candidate).Error
			if err != nil {
				return &[]Vision{}, err
			}
		}
	}
	return &missions, nil
}

func (p *Vision) FindVisionByID(db *gorm.DB, pid uint64) (*Vision, error) {
	var err error
	err = db.Debug().Model(&Vision{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Vision{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vision{}, err
		}
	}
	return p, nil
}

func (p *Vision) UpdateVision(db *gorm.DB) (*Vision, error) {

	var err error
	err = db.Debug().Model(&Vision{}).Where("id = ?", p.Id).Updates(Vision{Vision: p.Vision, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Vision{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vision{}, err
		}
	}
	return p, nil
}

func (p *Vision) DeleteVision(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Vision{}).Where("id = ?", pid).Take(&Vision{}).Delete(&Vision{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Vision not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
