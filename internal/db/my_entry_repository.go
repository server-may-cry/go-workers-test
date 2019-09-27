package db

import "github.com/server-may-cry/go-workers-test/internal/entity"

type MyEntryRepository struct {
	connection *Connection
}

func NewMyEntryRepository(conn *Connection) *MyEntryRepository {
	return &MyEntryRepository{
		connection: conn,
	}
}

func (r *MyEntryRepository) Insert(entry *entity.MyEntry) error {
	return r.connection.db.Create(entry).Error
}

func (r *MyEntryRepository) SearchLastInserted() (entity.MyEntry, bool) {
	var entry entity.MyEntry
	r.connection.db.Order("id DESC").First(&entry)

	return entry, entry.ID != 0
}
