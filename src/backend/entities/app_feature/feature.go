package app_feature

type Feature struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Apps []*App
}
