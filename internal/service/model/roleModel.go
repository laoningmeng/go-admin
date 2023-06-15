package model

type RoleStatus int

const (
	RoleStatusDisable = 0
	RoleStatusActive  = 1
)

type Role struct {
	BaseItem
	Name   string     `json:"name" gorm:"not null;"`
	Title  string     `json:"title" gorm:"not null;"`
	Status RoleStatus `json:"status" gorm:"default:0"`
	Rules  []Rule     `json:"rules" gorm:"many2many:role_rule"`
}
