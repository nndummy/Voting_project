package models

type Election struct {
	Title             string `json:"title"`
	Major             string `json:"major"`
	College           string `json:"college"`
	Content           string `json:"content"`
	ElectionStartTime string `json:"start_time"`
	ElectionEndTime   string `json:"end_time"`
	State             int    `json:"state" gorm:"default:1"`
	Id                int    `json:"election_id" gorm:"primary_key, auto_increment"`
	AdminId           string `json:"admin"`
}

/*
[state] 1: 투표 전 2: 투표 중 3: 투표 완료
*/
type EndElectionCandidateInfo struct {
	ElectionId  int    `json:"election_id" gorm:"primary_key, not_null"`
	All_vote    int    `json:"all_vote" gorm:"not_null"`
	CandidateId int    `json:"candidate_id"`
	Poll        int    `json:"poll"`
	StudentId   string `json:"student_id"`
	Name        string `json:"name"`
	Major       string `json:"major"`
	College     string `json:"college"`
	Thumbnail   string `json:"thumbnail"`
	Resume      string `json:"resume"`
}

type Elections []Election
type EndElectionResult []EndElectionCandidateInfo

type EndElection struct {
	ElectionId        int               `json:"election_id"`
	Title             string            `json:"title"`
	ElectionStartTime string            `json:"start_time"`
	ElectionEndTime   string            `json:"end_time"`
	State             int               `json:"state"`
	Candidate         EndElectionResult `json:"candidate"`
}
