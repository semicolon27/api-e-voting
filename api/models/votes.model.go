package models

import (
	"errors"
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

	// Check if ParticipantId has voted before
	existingVote := Vote{}
	err = db.Debug().Model(&Vote{}).Where("participant_id = ?", p.ParticipantId).Take(&existingVote).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Vote{}, err
	}
	if existingVote.Id != 0 {
		return &Vote{}, errors.New("Participant has already voted")
	}

	// save the vote
	err = db.Debug().Model(&Vote{}).Create(&p).Error
	if err != nil {
		return &Vote{}, err
	}
	if p.Id != 0 {
		err = db.Debug().Model(&Vote{}).Where("id = ?", p.CandidateId).Take(&p.Candidate).Error
		if err != nil {
			return &Vote{}, err
		}
	}
	return p, nil
}

// TODO: nanti bikin hitung vote
func (p *VoteCount) GetVoteCountByParticipantID(db *gorm.DB) (*[]VoteCount, error) {
	var voteCount []VoteCount
	err := db.Debug().Model(&Vote{}).
		Select("candidate_id, COUNT(*) as count").
		Group("candidate_id").
		Scan(&voteCount).Error
	if err != nil {
		return nil, err
	}
	if len(voteCount) > 0 {
		for i, _ := range voteCount {
			err := db.Debug().Model(&Candidate{}).Where("id = ?", voteCount[i].CandidateId).Take(&voteCount[i].Candidate).Error
			if err != nil {
				return &[]VoteCount{}, err
			}
			// get mission and vision
			if voteCount[i].CandidateId != 0 {
				err := db.Debug().Model(&Mission{}).Where("candidate_id = ?", &voteCount[i].Candidate.Id).Find(&voteCount[i].Candidate.Mission).Error
				if err != nil {
					return &[]VoteCount{}, err
				}
				err1 := db.Debug().Model(&Vision{}).Where("candidate_id = ?", &voteCount[i].Candidate.Id).Find(&voteCount[i].Candidate.Vision).Error
				if err1 != nil {
					return &[]VoteCount{}, err1
				}
			}
		}
	}
	return &voteCount, nil
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
