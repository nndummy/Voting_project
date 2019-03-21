package services

import (
	"time"
	"voting_system/app/services/models"

	"gopkg.in/robfig/cron.v2"
)

const createdFormat = "2006-01-02 15:04:05"

func init() {
	c := cron.New()
	//c.AddFunc("* * * * * *", CronStartVote) // 테스트용. 매 분 StartVote 펑션 실행
	//c.AddFunc("* * * * * *", CronEndVote)
	c.AddFunc("0 30 * * * *", CronStartVote)
	c.AddFunc("0 30 * * * *", CronEndVote)
	c.Start()
}

func CronStartVote() {
	currentTime := time.Now().Format(createdFormat)
	var startElections models.Elections //시작해야 하는 투표들 리스트

	db := votingdb
	db.Model(&models.Election{}).
		Where("election_start_time=?", currentTime).
		Find(&startElections)

	if len(startElections) <= 0 {
		return
	} else {
		for i := 0; i < len(startElections); i++ {
			err := db.Set("gorm:save_associations", false).
				Model(&models.Election{}).
				Where("id=?", startElections[i].Id).
				Update("state", 2).Error

			if err != nil {
				panic(err)
			}
		}
	}
}

func CronEndVote() {
	currentTime := time.Now().Format(createdFormat)
	var endElections models.Elections //시작해야 하는 투표들 리스트

	db := votingdb
	db.Model(&models.Election{}).
		Where("election_end_time=?", currentTime).
		Find(&endElections)

	if len(endElections) <= 0 {
		return
	} else {
		for i := 0; i < len(endElections); i++ {
			err := db.Set("gorm:save_associations", false).
				Model(&models.Election{}).
				Where("id=?", endElections[i].Id).
				Update("state", 3).Error

			if err != nil {
				panic(err)
			}

			var election models.Election
			err = db.Where("id=?", endElections[i].Id).
				Find(&election).Error

			var count int
			db.Model(&models.Votings{}).
				Where("election_id=?", endElections[i].Id).
				Count(&count)

			if err != nil {
				panic(err)
			}

			var candidatelist models.Candidates
			db.Model(&models.Candidate{}).
				Where("election_id=?", election.Id).
				Find(&candidatelist)

			for i := 0; i < len(candidatelist); i++ {
				var poll int

				db.Model(&models.Voting{}).
					Where("candidate_id=?", candidatelist[i].Id).
					Count(&poll)

				candidateInfo := AdminGetCandidateInfo(candidatelist[i].Id)
				endElectionCandidateInfo := models.EndElectionCandidateInfo{
					ElectionId:  election.Id,
					All_vote:    count,
					CandidateId: candidateInfo.Id,
					Poll:        poll,
					StudentId:   candidateInfo.StudentId,
					Name:        candidateInfo.Name,
					Major:       candidateInfo.Major,
					College:     candidateInfo.College,
					Thumbnail:   candidateInfo.Thumbnail,
					Resume:      candidateInfo.Resume,
				}

				err := db.Set("gorm:save_associations", false).
					Create(&endElectionCandidateInfo).Error

				if err != nil {
					panic(err)
				}
			}

		}
	}

}
