package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/solrac97gr/api-golang-jwt/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var pays = []models.Pay{
	models.Pay{
		Title:         "Title 1",
		Content:       "Hello world 1",
		Amount:        230.00,
		NParticipants: 5,
	},
	models.Pay{
		Title:         "Title 2",
		Content:       "Hello world 2",
		Amount:        230.00,
		NParticipants: 2,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Pay{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Pay{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Pay{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		pays[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Pay{}).Create(&pays[i]).Error
		if err != nil {
			log.Fatalf("cannot seed pays table: %v", err)
		}
	}
}
