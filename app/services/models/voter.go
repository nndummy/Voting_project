package models

type Voter struct {
	StudentId string `json:"student_id" gorm:"not_null, primary_key"`
	Password  string `json:"password" gorm:"not_null"`
	Name      string `json:"name" gorm:"not_null"`
	Major     string `json:"major" gorm:"not_null"`
	College   string `json:"college" gorm:"not_null"` // ex) 공과 대학, 어문 대학, ...
	Mobile    string `json:"mobile" gorm:"not_null"`
	Address   string `json:"address" gorm:"not_null"`
	Email     string `json:"email" gorm:"not_null"`
	Sex       string `json:"sex" gorm:"not_null"`
	Birth     string `json:"birth" gorm:"not_null" `
}
