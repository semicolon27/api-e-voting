package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Mission struct {
	Id          int       `gorm:"primary_key;auto_increment" json:"id"`
	CandidateId int       `sql:"type:int REFERENCES candidates(id)" json:"candidateid"`
	Candidate   Candidate `json:"candidate"`
	Mission     string    `gorm:"size:255;not null" json:"mission"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Mission) Prepare() {
	p.Id = 0
	p.Candidate = Candidate{}
	p.Mission = html.EscapeString(strings.TrimSpace(p.Mission))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Mission) Validate() error {

	// if p.Option == "" {
	// 	return errors.New("Required Option")
	// }
	if p.Mission == "" {
		return errors.New("Required Mission")
	}
	return nil
}

func (p *Mission) SaveMission(db *gorm.DB) (*Mission, error) {
	var err error
	err = db.Debug().Model(&Mission{}).Create(&p).Error
	if err != nil {
		return &Mission{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Mission{}, err
		}
	}
	return p, nil
}

func (p *Mission) FindAllMissions(db *gorm.DB) (*[]Mission, error) {
	var err error
	missions := []Mission{}
	err = db.Debug().Model(&Mission{}).Limit(100).Find(&missions).Error
	if err != nil {
		return &[]Mission{}, err
	}
	if len(missions) > 0 {
		for i, _ := range missions {
			err := db.Debug().Model(&Candidate{}).Where("id = ?", missions[i].CandidateId).Take(&missions[i].Candidate).Error
			if err != nil {
				return &[]Mission{}, err
			}
		}
	}
	return &missions, nil
}

func (p *Mission) FindMissionByID(db *gorm.DB, pid uint64) (*Mission, error) {
	var err error
	err = db.Debug().Model(&Mission{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Mission{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Mission{}, err
		}
	}
	return p, nil
}

func (p *Mission) UpdateMission(db *gorm.DB) (*Mission, error) {

	var err error
	err = db.Debug().Model(&Mission{}).Where("id = ?", p.Id).Updates(Mission{Mission: p.Mission, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Mission{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Mission{}, err
		}
	}
	return p, nil
}

func (p *Mission) DeleteMission(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Mission{}).Where("id = ?", pid).Take(&Mission{}).Delete(&Mission{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Mission not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
