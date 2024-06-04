package models

type Appointment struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Time     string
	Master   string
}
