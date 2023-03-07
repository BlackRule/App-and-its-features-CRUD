package app_feature

type App struct {
	ID       uint       `gorm:"primary_key"`
	Name     string     `gorm:"not null"`
	Features []*Feature `gorm:"many2many:app_features;"`
}