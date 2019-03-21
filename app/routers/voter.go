package routers

import (
	"voting_system/app/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutesVoter(r *gin.RouterGroup) {
	r.GET("/voter/elections", controllers.VoterGetElectionsList)
	r.GET("/voter/election/info/:electionid", controllers.VoterGetElectionInfo)
	r.GET("/voter/election/info/:electionid/candidates", controllers.VoterGetElectionCandidatesList)
	r.GET("/voter/elections/result", controllers.VoterGetElectionResult)
	r.GET("/voter/candidates/:candidateid", controllers.VoterGetCandidateInfo)
	r.POST("/voter/elections/voting", controllers.VoterVoting)

	r.GET("/servertime", controllers.GetServerTime)
}
