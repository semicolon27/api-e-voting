package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Vote struct {
	Id            int         `gorm:"primary_key;auto_increment" json:"id"`
	ParticipantId int         `sql:"type:int REFERENCES participants(id)" json:"participantid"`
	Participant   Participant `json:"participant"`
	CandidateId   int         `sql:"type:int REFERENCES candidates(id)" json:"candidateid"`
	Candidate     Candidate   `json:"candidate"`
	CreatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Vote) Prepare() {
	p.Id = 0
	p.Candidate = Candidate{}
	p.Participant = Participant{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Vote) SaveVote(db *gorm.DB) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Create(&p).Error
	if err != nil {
		return &Vote{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

// TODO: nanti bikin hitung vote

func (p *Vote) FindAllVotes(db *gorm.DB) (*[]Vote, error) {
	var err error
	missions := []Vote{}
	err = db.Debug().Model(&Vote{}).Limit(100).Find(&missions).Error
	if err != nil {
		return &[]Vote{}, err
	}
	if len(missions) > 0 {
		for i, _ := range missions {
			err := db.Debug().Model(&Candidate{}).Where("id = ?", missions[i].CandidateId).Take(&missions[i].Candidate).Error
			if err != nil {
				return &[]Vote{}, err
			}
		}
	}
	return &missions, nil
}

func (p *Vote) FindVoteByID(db *gorm.DB, pid uint64) (*Vote, error) {
	var err error
	err = db.Debug().Model(&Vote{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Vote{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Candidate{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

// ? ga pake update

// ? ga pake delete juga
