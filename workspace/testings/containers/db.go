package containers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	db *gorm.DB
}

type Member struct {
	gorm.Model
	Email    string `gorm:"column:email;unique_index"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password; not null"`
}

// ========================= Database

// NewDatabase returns a new db given connection arg with mysql dialect
func NewDatabase(conn string) (*DB, error) {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (d *DB) AutoMigrate() {
	d.db.AutoMigrate(&Member{})
}

func (d *DB) DropTables() {
	d.db.DropTableIfExists(&Member{})
}

func (d *DB) Close() error {
	return d.db.Close()
}

// ========================= Member repository

func (d *DB) Save(m *Member) error {
	return d.db.Create(m).Error
}

func (d *DB) FindByEmail(email string) (*Member, error) {
	var m Member
	err := d.db.Where(&Member{Email: email}).First(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (d *DB) FindAllByUsername(username string) ([]Member, error) {
	var members []Member
	err := d.db.Where(&Member{Username: username}).Find(&members).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return members, nil
}
