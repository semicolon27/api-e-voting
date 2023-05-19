package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/semicolon27/api-e-voting/api/models"
)

var admins = []models.Admin{
	models.Admin{
		Username: "admin1",
		Password: "admin1",
	},
	models.Admin{
		Username: "admin2",
		Password: "admin2",
	},
}

var candidates = []models.Candidate{
	models.Candidate{
		Option: 1,
		Name:   "Alfi",
		Image:  "Gambar Alfi",
	},
	models.Candidate{
		Option: 2,
		Name:   "George W Bush",
		Image:  "Gambar Bush",
	},
}

var classes = []models.Class{
	models.Class{
		Name: "Ilkom O 4",
	},
	models.Class{
		Name: "Farmasi O 4",
	},
}

var missions = []models.Mission{
	models.Mission{
		CandidateId: 1,
		Mission:     "Memberikan fasilitas gratis",
	},
	models.Mission{
		CandidateId: 1,
		Mission:     "Mengurangi gaji DPR",
	},
	models.Mission{
		CandidateId: 2,
		Mission:     "Memperbaiki seluruh jalan rusak tanpa di korupsi",
	},
	models.Mission{
		CandidateId: 2,
		Mission:     "Mengurangi durasi lampu merah hingga 50%",
	},
}

var participants = []models.Participant{
	models.Participant{
		RegNumber: "0001",
		Name:      "participant1",
		Password:  "participant1",
		ClassId:   1,
	},
	models.Participant{
		RegNumber: "0002",
		Name:      "participant2",
		Password:  "participant2",
		ClassId:   1,
	},
	models.Participant{
		RegNumber: "0003",
		Name:      "participant3",
		Password:  "participant3",
		ClassId:   1,
	},
	models.Participant{
		RegNumber: "0004",
		Name:      "participant4",
		Password:  "participant4",
		ClassId:   1,
	},
	models.Participant{
		RegNumber: "0005",
		Name:      "participant5",
		Password:  "participant5",
		ClassId:   1,
	},
}

var visions = []models.Vision{
	models.Vision{
		CandidateId: 1,
		Vision:      "Menyatukan Indonesia dan Malaysia",
	},
	models.Vision{
		CandidateId: 1,
		Vision:      "Menaikan IQ rata2 orang wakanda",
	},
	models.Vision{
		CandidateId: 2,
		Vision:      "Menjadikan Rupiah sebagai mata uang dunia",
	},
	models.Vision{
		CandidateId: 2,
		Vision:      "Menyatukan Korea Utara dan Korea Selatan",
	},
}

var votes = []models.Vote{
	models.Vote{
		CandidateId:   1,
		ParticipantId: 1,
	},
	models.Vote{
		CandidateId:   1,
		ParticipantId: 2,
	},
	models.Vote{
		CandidateId:   1,
		ParticipantId: 3,
	},
	models.Vote{
		CandidateId:   2,
		ParticipantId: 4,
	},
	models.Vote{
		CandidateId:   1,
		ParticipantId: 5,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(
		&models.Admin{},
		&models.Candidate{},
		&models.Class{},
		&models.Mission{},
		&models.Participant{},
		&models.Vision{},
		&models.Vote{},
	).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(
		&models.Admin{},
		&models.Candidate{},
		&models.Class{},
		&models.Mission{},
		&models.Participant{},
		&models.Vision{},
		&models.Vote{},
	).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "admin(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i, _ := range admins {
		err = db.Debug().Model(&models.Admin{}).Create(&admins[i]).Error
		if err != nil {
			log.Fatalf("cannot seed admins table: %v", err)
		}
	}

	for i, _ := range candidates {
		err = db.Debug().Model(&models.Candidate{}).Create(&candidates[i]).Error
		if err != nil {
			log.Fatalf("cannot seed candidates table: %v", err)
		}
	}

	for i, _ := range classes {
		err = db.Debug().Model(&models.Class{}).Create(&classes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed classes table: %v", err)
		}
	}

	for i, _ := range missions {
		err = db.Debug().Model(&models.Mission{}).Create(&missions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed mission table: %v", err)
		}
	}

	for i, _ := range participants {
		err = db.Debug().Model(&models.Participant{}).Create(&participants[i]).Error
		if err != nil {
			log.Fatalf("cannot seed participants table: %v", err)
		}
	}

	for i, _ := range visions {
		err = db.Debug().Model(&models.Vision{}).Create(&visions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed visions table: %v", err)
		}
	}

	for i, _ := range votes {
		err = db.Debug().Model(&models.Vote{}).Create(&votes[i]).Error
		if err != nil {
			log.Fatalf("cannot seed votes table: %v", err)
		}
	}
}
