package services

import (
	"time"
	"voting_system/app/helpers"
	"voting_system/app/services/models"
)

func VoterGetElectionsList(pagination helpers.Pagination) models.Elections {
	var elections models.Elections
	db := votingdb
	db.Order("id desc").
		Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Find(&elections)

	return elections
}

func VoterGetElectionInfo(electionid int) models.Election {
	var election models.Election

	db := votingdb
	db.Model(&models.Election{}).
		Where("id=?", electionid).
		First(&election)

	return election
}

func VoterGetElectionCandidatesList(electionid int) models.Candidates {
	var candidates models.Candidates

	db := votingdb
	db.Model(&models.Candidate{}).
		Where("election_id=?", electionid).
		Find(&candidates)

	return candidates
}

func VoterGetCandidateInfo(candidateid int) models.Candidate {
	var candidate models.Candidate
	db := votingdb
	db.Model(&models.Candidate{}).
		Where("id=?", candidateid).
		First(&candidate)

	return candidate
}

func VoterVoting(electionid, candidateid int) error {
	record := models.Voting{
		ElectionId:  electionid,
		CandidateId: candidateid,
		// Auto_hash: "",
	}

	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	return err
}

func VoterGetElectionResult() models.Elections {
	var endElections models.Elections

	db := votingdb
	db.Where("state=?", 3).
		Find(&endElections)

	return endElections
}

func VoterGetElectionResultCandidate(electionid int) models.EndElectionResult {
	var candidates models.EndElectionResult

	db := votingdb
	db.Model(&models.EndElectionCandidateInfo{}).
		Where("election_id=?", electionid).
		Order("poll desc").
		Find(&candidates)

	return candidates
}

func GetServerTime() int64 {
	currentTime := time.Now().Unix()
	return currentTime
}

func GetVoterInfo(studentId, password string) models.Voter {
	var voter models.Voter

	db := votingdb
	db.Model(&models.Voter{}).
		Where("student_id=? AND password=?", studentId, password).
		First(&voter)

	return voter
}
