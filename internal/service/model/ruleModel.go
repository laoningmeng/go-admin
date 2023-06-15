package model

const (
	RuleStatusDisable = 0
	RuleStatusActive  = 1
)

type RuleStatus int
type Rule struct {
	BaseItem
	Name   string     `json:"name" gorm:"not null;"`
	Title  string     `json:"title" gorm:"not null;"`
	Status RuleStatus `json:"status" gorm:"default:0"`
}
