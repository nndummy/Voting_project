package controllers

import (
	"strconv"
	"voting_system/app/helpers"
	"voting_system/app/services"
	"voting_system/app/services/models"

	"github.com/gin-gonic/gin"
)

func VoterGetElectionsList(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountElections := services.CountRowsElection()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountElections, currentPage, itemPerPage)

	var list models.Elections
	list = services.VoterGetElectionsList(pagination)

	if len(list) <= 0 {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "생성된 선거가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":       200,
			"current_page": currentPage,
			"list":         list,
		})
	}
}

func VoterGetElectionInfo(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	election := services.VoterGetElectionInfo(electionId)

	if election == (models.Election{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":   200,
			"election": election,
		})
	}
}

func VoterGetElectionCandidatesList(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	list := services.VoterGetElectionCandidatesList(electionId)

	if len(list) <= 0 {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "해당 선거엔 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":    200,
			"candidate": list,
		})
	}
}

func VoterGetCandidateInfo(c *gin.Context) {
	candidateId, err := strconv.Atoi(c.Params.ByName("candidateid"))
	if err != nil {
		panic(err)
	}
	candidate := services.VoterGetCandidateInfo(candidateId)

	if candidate == (models.Candidate{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":   200,
			"election": candidate,
		})
	}
}

func VoterVoting(c *gin.Context) {
	var request models.Voting
	c.ShouldBindJSON(&request)

	err := services.VoterVoting(request.ElectionId, request.CandidateId)

	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "투표 과정에 오류가 발생했습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status": 200,
		})
	}
}

func VoterGetElectionResult(c *gin.Context) {
	endElections := services.VoterGetElectionResult()

	if len(endElections) <= 0 {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "완료된 선거가 없습니다",
		})
	} else {
		result := make([]models.EndElection, len(endElections))
		for i := 0; i < len(endElections); i++ {
			result[i].ElectionId = endElections[i].Id
			result[i].Title = endElections[i].Title
			result[i].ElectionStartTime = endElections[i].ElectionStartTime
			result[i].ElectionEndTime = endElections[i].ElectionEndTime
			result[i].State = endElections[i].State
			result[i].Candidate = services.VoterGetElectionResultCandidate(endElections[i].Id)
		}

		c.JSON(200, gin.H{
			"status":   200,
			"election": result,
		})
	}
}

func GetServerTime(c *gin.Context) {
	currentTime := services.GetServerTime()

	c.JSON(200, gin.H{
		"status":       200,
		"current_time": currentTime,
	})
}
