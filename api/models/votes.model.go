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

type VoteCount struct {
	CandidateId int       `sql:"type:int REFERENCES candidates(id)" json:"candidateid"`
	Candidate   Candidate `json:"candidate"`
	Count       string    `json:"count"`
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
		err = db.Debug().Model(&Vote{}).Where("id = ?", p.Id).Take(&p.Candidate).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

// TODO: nanti bikin hitung vote
func (p *VoteCount) GetVoteCountByParticipantID(db *gorm.DB) ([]VoteCount, error) {
	var voteCount []VoteCount
	err := db.Debug().Model(&Vote{}).
		Select("participant_id, COUNT(*) as asdasd").Take(&p.Candidate).
		Group("participant_id").
		Scan(&voteCount).Error
	if err != nil {
		return nil, err
	}
	return voteCount, nil
}

func (p *Vote) FindAllVotes(db *gorm.DB) (*[]Vote, error) {
	var err error
	vote := []Vote{}
	err = db.Debug().Model(&Vote{}).Limit(100).Find(&vote).Error
	if err != nil {
		return &[]Vote{}, err
	}
	return &vote, nil
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
