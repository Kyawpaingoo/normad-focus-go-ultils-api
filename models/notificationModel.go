package models

import "time"

type Notification struct {
	Id          uint      `gorm:"primary_key" json:"id,omitempty"`
	User_Id     uint      `json:"user___id,omitempty"`
	Source_Type string    `json:"source___type,omitempty"`
	Source_Id   uint      `json:"source___id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Message     string    `json:"message,omitempty"`
	Notify_At   time.Time `json:"notify___at"`
	Sent_At     time.Time `json:"sent___at"`
	Sent        bool      `json:"sent,omitempty"`
}

func (Notification) TableName() string {
	return "Notification" // <-- This fixes the issue
}
