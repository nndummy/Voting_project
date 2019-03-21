package models

type Candidate struct {
	StudentId  string `json:"student_id"`
	Name       string `json:"name"`
	Major      string `json:"major"`
	College    string `json:"college"`
	Thumbnail  string `json:"thumbnail"`
	Resume     string `json:"resume"`
	Id         int    `json:"candidate_id" gorm:"default=0, auto_increment, primary_key"`
	ElectionId int    `json:"election_id"`
	//Poll       int    `json:"poll"`
}

type Candidates []Candidate
