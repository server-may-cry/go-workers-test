package entity

type MyEntry struct {
	ID      uint64 `gorm:"primary_key"`
	Message string
	Source  string
}
