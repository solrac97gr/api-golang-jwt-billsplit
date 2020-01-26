package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Pay struct {
	ID            uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title         string    `gorm:"size:255;not null;" json:"title"`
	Content       string    `gorm:"size:255;not null;" json:"content"`
	NParticipants uint64    `gorm:"not null" json:"nparticipants"`
	Amount        float64   `gorm:"not null" json:"amount"`
	Author        User      `json:"author"`
	AuthorID      uint32    `gorm:"not null" json:"author_id"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Pay) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Pay) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.Amount == 0 {
		return errors.New("Required amount")
	}

	if p.NParticipants == 0 {
		return errors.New("Required Number of participants")
	}

	if p.AuthorID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (p *Pay) SavePay(db *gorm.DB) (*Pay, error) {
	var err error
	err = db.Debug().Model(&Pay{}).Create(&p).Error
	if err != nil {
		return &Pay{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Pay{}, err
		}
	}
	return p, nil
}

func (p *Pay) FindAllPays(db *gorm.DB) (*[]Pay, error) {
	var err error
	pays := []Pay{}
	err = db.Debug().Model(&Pay{}).Limit(100).Find(&pays).Error
	if err != nil {
		return &[]Pay{}, err
	}
	if len(pays) > 0 {
		for i, _ := range pays {
			err := db.Debug().Model(&User{}).Where("id = ?", pays[i].AuthorID).Take(&pays[i].Author).Error
			if err != nil {
				return &[]Pay{}, err
			}
		}
	}
	return &pays, nil
}

func (p *Pay) FindPayByID(db *gorm.DB, pid uint64) (*Pay, error) {
	var err error
	err = db.Debug().Model(&Pay{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Pay{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Pay{}, err
		}
	}
	return p, nil
}
func (p *Pay) FindPaysByUserID(db *gorm.DB, pid uint64) (*[]Pay, error) {
	var err error
	pays := []Pay{}
	err = db.Debug().Model(&Pay{}).Where("author_id = ?", pid).Limit(100).Find(&pays).Error
	if err != nil {
		return &[]Pay{}, err
	}
	if len(pays) > 0 {
		for i, _ := range pays {
			err := db.Debug().Model(&User{}).Where("id = ?", pays[i].AuthorID).Take(&pays[i].Author).Error
			if err != nil {
				return &[]Pay{}, err
			}
		}
	}
	return &pays, nil
}

func (p *Pay) UpdateAPay(db *gorm.DB) (*Pay, error) {

	var err error

	err = db.Debug().Model(&Pay{}).Where("id = ?", p.ID).Updates(Pay{Title: p.Title, Content: p.Content, UpdatedAt: time.Now(), Amount: p.Amount, NParticipants: p.NParticipants}).Error
	if err != nil {
		return &Pay{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Pay{}, err
		}
	}
	return p, nil
}

func (p *Pay) DeleteAPay(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Pay{}).Where("id = ? and author_id = ?", pid, uid).Take(&Pay{}).Delete(&Pay{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Pay not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
