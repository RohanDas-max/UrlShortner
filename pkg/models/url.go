package models

type Url struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Url      string `json:"goly" gorm:"unique;not null"`
	Clicked  uint   `json:"clicked"`
	Random   bool   `json:"random"`
}
