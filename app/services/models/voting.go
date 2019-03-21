package models

type Voting struct {
	Id          int    `json:"id" gorm:"not null default:0, auto_increment, primary_key"`
	CandidateId int    `json:"candidate_id" gorm:"not_null"`
	ElectionId  int    `json:"election_id" gorm:"not_null"`
	Auto_hash   string `json:"auto_hash"`
}

type Votings []Voting
