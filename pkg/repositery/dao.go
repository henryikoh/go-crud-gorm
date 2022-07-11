package repositery

import (
	"github.com/henryikoh/book-management-go/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// interface for the repo/database there would be various types of queries based on your
// usecase I just implimented one.

// DAO stands for Data Access Object, This is what would grant access to the DB to servives and handlers
type DAO struct {
	db *gorm.DB
}

func InitDAO() *DAO {
	// open database connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// migerate new data
	db.AutoMigrate(&models.Book{})
	s := &DAO{
		db: db,
	}
	return s
}

func (d *DAO) NewMovieQuery() MovieQuery {
	return NewMovieQuery(d.db)
}
