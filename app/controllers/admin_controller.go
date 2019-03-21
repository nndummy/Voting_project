package controllers

import (
	"strconv"
	"voting_system/app/helpers"
	"voting_system/app/services"
	"voting_system/app/services/models"

	"github.com/gin-gonic/gin"
)

// 선거 관리
func AdminGetElectionsList(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountElections := services.CountRowsElection()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountElections, currentPage, itemPerPage)

	var list models.Elections
	list = services.AdminGetElectionsList(pagination)

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

func AdminGetElectionInfo(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	election := services.AdminGetElectionInfo(electionId)

	if election == (models.Election{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 선거가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":   200,
			"election": election,
		})
	}
}
func AdminGetElectionCandidatesList(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	candidates := services.AdminGetElectionCandidateList(electionId)

	if len(candidates) <= 0 {
		c.JSON(200, gin.H{
			"status": 200,
			"error":  "해당 선거엔 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":    200,
			"candidate": candidates,
		})
	}
}
func AdminCreateElection(c *gin.Context) {
	var request models.Election
	c.ShouldBindJSON(&request)

	election, err := services.AdminCreateElection(request)

	if err == nil {
		c.JSON(201, gin.H{
			"status":         201,
			"message":        "새로운 선거를 생성했습니다",
			"election_state": election.State,
		})
	} else {
		c.JSON(500, gin.H{
			"status": 500,
			"error":  "서버에 오류가 발생했습니다",
		})
	}
}

func AdminStartElection(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	election := services.AdminGetElectionInfo(electionId)

	if election == (models.Election{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 선거가 없습니다",
		})
	} else {
		election, err := services.AdminStartElection(electionId)

		if err == nil {
			c.JSON(201, gin.H{
				"status":  201,
				"message": "해당 선거를 수정했습니다",
				"state":   election.State,
			})
		} else {
			c.JSON(500, gin.H{
				"status": 500,
				"error":  "서버에 오류가 발생했습니다",
			})
		}
	}
}

func AdminEndElection(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	election := services.AdminGetElectionInfo(electionId)

	if election == (models.Election{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 선거가 없습니다",
		})
	} else {
		election, err := services.AdminEndElection(electionId)

		if err == nil {
			c.JSON(201, gin.H{
				"status":  201,
				"message": "해당 선거를 수정했습니다",
				"state":   election.State,
			})
		} else {
			c.JSON(500, gin.H{
				"status": 500,
				"error":  "서버에 오류가 발생했습니다",
			})
		}
	}
}

func AdminEditElection(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	election := services.AdminGetElectionInfo(electionId)

	if election == (models.Election{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 선거가 없습니다",
		})
	} else {
		var request models.Election
		err := c.ShouldBindJSON(&request)

		if err == nil {
			election, err := services.AdminEditElection(electionId, request)

			if err == nil {
				c.JSON(201, gin.H{
					"status":   201,
					"message":  "해당 선거를 수정했습니다",
					"election": election,
				})
			} else {
				c.JSON(500, gin.H{
					"status": 500,
					"error":  "서버에 오류가 발생했습니다",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"status": 400,
				"error":  "잘못된 요청입니다",
			})
		}
	}
}

func AdminElectionResult(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountElections := services.CountRowsElection()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountElections, currentPage, itemPerPage)

	endElections := services.AdminElectionResult(pagination)

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
			result[i].Candidate = services.AdminElectionResultCandidates(endElections[i].Id)
		}

		c.JSON(200, gin.H{
			"status":       200,
			"current_page": currentPage,
			"election":     result,
		})
	}
}

// 후보자 관리
func AdminGetCandidatesList(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountCandidates := services.CountRowsCandidate()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountCandidates, currentPage, itemPerPage)

	var list models.Candidates
	list = services.AdminGetCandidatesList(pagination)

	if len(list) <= 0 {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "생성된 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":       200,
			"current_page": currentPage,
			"list":         list,
		})
	}
}

func AdminCandidateInfo(c *gin.Context) {
	candidateId, err := strconv.Atoi(c.Params.ByName("candidateid"))
	if err != nil {
		panic(err)
	}
	candidate := services.AdminGetCandidateInfo(candidateId)

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

func AdminCreateCandidate(c *gin.Context) {
	var request models.Candidate
	c.ShouldBindJSON(&request)

	candidate, err := services.AdminCreateCandidate(request)

	if err == nil {
		c.JSON(201, gin.H{
			"status":       201,
			"message":      "새로운 후보자를 생성했습니다",
			"candidate_id": candidate.Id,
		})
	} else {
		c.JSON(500, gin.H{
			"status": 500,
			"error":  "서버에 오류가 발생했습니다",
		})
	}
}

func AdminEditCandidate(c *gin.Context) {
	candidateId, err := strconv.Atoi(c.Params.ByName("candidateid"))
	if err != nil {
		panic(err)
	}
	candidate := services.AdminGetCandidateInfo(candidateId)

	if candidate == (models.Candidate{}) {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  "등록된 후보자가 없습니다",
		})
	} else {
		var request models.Candidate
		err := c.ShouldBindJSON(&request)

		if err == nil {
			candidate, err := services.AdminEditCandidate(candidateId, request)

			if err == nil {
				c.JSON(201, gin.H{
					"status":    201,
					"message":   "해당 후보자를 수정했습니다",
					"candidate": candidate,
				})
			} else {
				c.JSON(500, gin.H{
					"status": 500,
					"error":  "서버에 오류가 발생했습니다",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"status": 400,
				"error":  "잘못된 요청입니다",
			})
		}
	}
}
