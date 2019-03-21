package models

type Administrator struct {
	Id        string `json:"id" gorm:"primary_key, auto_increment, default=0"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	Ssn       string `json:"ssn"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Sex       string `json:"sex"`
	Birth     string `json:"birth"`
	Authority int    `json:"authority"`
}

/*
Authority
1: 선거, 후보자 둘 다 관리
2: 선거만 관리
3: 후보자만 관리
*/
