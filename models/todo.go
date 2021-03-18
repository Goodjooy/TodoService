package models

import (
	"time"
)


type Todo struct {
	Model `form:"-" json:"model"`

	Title string `gorm:"type:varchar(32);not null" form:"title"`
	Body string `gorm:"type:text(255);not null" form:"body"`

	Status TodoStatus `gorm:"not null" form:"-"`

	User UserModel`form:"-" json:"-"`
	UserID uint `gorm:"not null" form:"-" json:"-"`

	DeathLine time.Time `gorm:"not null" form:"-"`

}
type  TodoStatus uint
const(
	UnDone TodoStatus=1+iota
	Done
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
