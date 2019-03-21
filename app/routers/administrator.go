package routers

import (
	"voting_system/app/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoutesAdministrator(r *gin.RouterGroup) {
	// 선거 관리
	r.GET("/admin/elections", controllers.AdminGetElectionsList)
	r.GET("/admin/elections/info/:electionid", controllers.AdminGetElectionInfo)
	r.GET("/admin/elections/info/:electionid/candidates", controllers.AdminGetElectionCandidatesList)
	r.GET("/admin/elections/result", controllers.AdminElectionResult)

	r.POST("/admin/elections/create", controllers.AdminCreateElection)
	r.PUT("/admin/elections/start/:electionid", controllers.AdminStartElection)
	r.PUT("/admin/elections/end/:electionid", controllers.AdminEndElection)
	r.PUT("/admin/elections/edit/:electionid", controllers.AdminEditElection)

	r.GET("/admin/candidates", controllers.AdminGetCandidatesList)
	r.GET("/admin/candidates/:candidateid", controllers.AdminCandidateInfo)
	r.POST("/admin/candidates/create", controllers.AdminCreateCandidate)
	r.PUT("/admin/candidates/edit/:candidateid", controllers.AdminEditCandidate)
}
