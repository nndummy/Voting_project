package database

import (
	"strconv"
	"time"
	"voting_system/app/services/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/votingDB?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)

	DropTablesIfExists(db)
	AutoMigrate(db)
	AutoPopulate(db)
	AddForeignKeys(db)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Administrator{},
		&models.Candidate{},
		&models.Election{},
		&models.EndElectionCandidateInfo{},
		&models.Voter{},
		&models.Voting{},
	)
}

func DropTablesIfExists(db *gorm.DB) {
	//db.Exec("DROP TABLE Administrators, Candidates, Elections, EndElectionCandidateInfos, Voters, Votings CASCADE;")
	db.DropTable(&models.Administrator{})
	db.DropTable(&models.Candidate{})
	db.DropTable(&models.Election{})
	db.DropTable(&models.EndElectionCandidateInfo{})
	db.DropTable(&models.Voter{})
	db.DropTable(&models.Voting{})
}

func AddForeignKeys(db *gorm.DB) {

}

func AutoPopulate(db *gorm.DB) {
	PopulateAdmins(db)
	PopulateVoters(db)
	PopulateElection(db)
	PopulateEndElectionsCandidateInfo(db)
	PopulateCandidates(db)
}

func PopulateAdmins(db *gorm.DB) {
	db.Create(&models.Administrator{
		Id:        "201202274",
		Password:  "201202274",
		Name:      "윤인배",
		Mobile:    "010-2786-2455",
		Ssn:       "",
		Address:   "서초구 내곡동",
		Email:     "iby2455@gmail.com",
		Sex:       "남",
		Birth:     "10/19/1993",
		Authority: 1,
	})
	db.Create(&models.Administrator{
		Id:        "201200000",
		Password:  "000001",
		Name:      "김철수",
		Mobile:    "010-1423-2455",
		Ssn:       "",
		Address:   "서초구 잠실동",
		Email:     "chulsu@gmail.com",
		Sex:       "남",
		Birth:     "10/10/1993",
		Authority: 2,
	})
	db.Create(&models.Administrator{
		Id:        "201211111",
		Password:  "000002",
		Name:      "안현주",
		Mobile:    "010-1423-3942",
		Ssn:       "",
		Address:   "강남구 역삼동",
		Email:     "djdj@gmail.com",
		Sex:       "여",
		Birth:     "12/11/1993",
		Authority: 3,
	})
}

func UnixToTimestamp(unixtime string) time.Time {
	i, err := strconv.ParseInt(unixtime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}

func PopulateElection(db *gorm.DB) {
	db.Create(&models.Election{
		Title:             "제 101호 부학생회장",
		Major:             "",
		College:           "종합대학",
		Content:           "부학생회장선거",
		ElectionStartTime: "2018-11-12 12:23:00",
		ElectionEndTime:   "2018-11-12 12:23:00",
		State:             1,
		Id:                98,
		AdminId:           "201202274",
	})
	db.Create(&models.Election{
		Title:             "제 101호 학생회장",
		Major:             "",
		College:           "종합대학",
		Content:           "종합대학회장선거",
		ElectionStartTime: "2018-11-12 12:23:00",
		ElectionEndTime:   "2018-11-12 12:23:00",
		State:             2,
		Id:                99,
		AdminId:           "201202274",
	})
	db.Create(&models.Election{
		Title:             "제 101호 컴퓨터공학과 회장 선거",
		Major:             "컴퓨터공학과",
		College:           "공과대학",
		Content:           "컴퓨터공학과회장선거",
		ElectionStartTime: "2018-11-12 12:00:00",
		ElectionEndTime:   "2018-11-12 14:00:00",
		State:             3,
		Id:                100,
		AdminId:           "201202274",
	})
	db.Create(&models.Election{
		Title:             "제 101호 컴퓨터공학과 부회장 선거",
		Major:             "컴퓨터공학과",
		College:           "공과대학",
		Content:           "컴퓨터공학과부회장선거",
		ElectionStartTime: "2018-11-12 12:00:00",
		ElectionEndTime:   "2018-11-12 14:00:00",
		State:             3,
		Id:                101,
		AdminId:           "201202274",
	})
}

func PopulateEndElectionsCandidateInfo(db *gorm.DB) {
	db.Create(&models.EndElectionCandidateInfo{
		ElectionId:  100,
		All_vote:    100,
		CandidateId: 10,
		Poll:        70,
		StudentId:   "201301923",
		Name:        "도로링",
		Major:       "컴퓨터공학과",
		College:     "공과대학",
		Thumbnail:   "aaa",
		Resume:      "ddd",
	})
	db.Create(&models.EndElectionCandidateInfo{
		ElectionId:  100,
		All_vote:    100,
		CandidateId: 11,
		Poll:        30,
		StudentId:   "201201932",
		Name:        "아라링",
		Major:       "컴퓨터공학과",
		College:     "공과대학",
		Thumbnail:   "aaa",
		Resume:      "ddd",
	})
	db.Create(&models.EndElectionCandidateInfo{
		ElectionId:  101,
		All_vote:    100,
		CandidateId: 12,
		Poll:        65,
		StudentId:   "20133299",
		Name:        "소로링",
		Major:       "컴퓨터공학과",
		College:     "공과대학",
		Thumbnail:   "aaa",
		Resume:      "ddd",
	})
	db.Create(&models.EndElectionCandidateInfo{
		ElectionId:  101,
		All_vote:    100,
		CandidateId: 13,
		Poll:        35,
		StudentId:   "201492392",
		Name:        "기리링",
		Major:       "컴퓨터공학과",
		College:     "공과대학",
		Thumbnail:   "aaa",
		Resume:      "ddd",
	})

}
func PopulateCandidates(db *gorm.DB) {
	db.Create(&models.Candidate{
		StudentId:  "201301923",
		Name:       "도로링",
		Major:      "컴퓨터공학과",
		College:    "공과대학",
		Thumbnail:  "aaa",
		Resume:     "ddd",
		Id:         10,
		ElectionId: 100,
	})
	db.Create(&models.Candidate{
		StudentId:  "201201932",
		Name:       "아라링",
		Major:      "컴퓨터공학과",
		College:    "공과대학",
		Thumbnail:  "aaa",
		Resume:     "ddd",
		Id:         11,
		ElectionId: 100,
	})
	db.Create(&models.Candidate{
		StudentId:  "20133299",
		Name:       "소로링",
		Major:      "컴퓨터공학과",
		College:    "공과대학",
		Thumbnail:  "aaa",
		Resume:     "ddd",
		Id:         12,
		ElectionId: 101,
	})
	db.Create(&models.Candidate{
		StudentId:  "201492392",
		Name:       "기리링",
		Major:      "컴퓨터공학과",
		College:    "공과대학",
		Thumbnail:  "aaa",
		Resume:     "ddd",
		Id:         13,
		ElectionId: 101,
	})
}

func PopulateVoters(db *gorm.DB) {
	db.Create(&models.Voter{
		StudentId: "201202274",
		Password:  "201202274",
		Name:      "윤인배",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-2786-2455",
		Address:   "서초구 내곡동",
		Email:     "iby2455@naver.com",
		Sex:       "M",
		Birth:     "1993-10-19",
	})
	db.Create(&models.Voter{
		StudentId: "201200000",
		Password:  "201200000",
		Name:      "김영영",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0000",
		Address:   "강남구",
		Email:     "koo@gmail.com",
		Sex:       "M",
		Birth:     "1993-01-01",
	})
	db.Create(&models.Voter{
		StudentId: "201200001",
		Password:  "201200001",
		Name:      "김영일",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0001",
		Address:   "강남구",
		Email:     "asd@gmail.com",
		Sex:       "F",
		Birth:     "1993-01-02",
	})
	db.Create(&models.Voter{
		StudentId: "201200002",
		Password:  "201200002",
		Name:      "김영이",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0002",
		Address:   "강남구",
		Email:     "zxdc@gmail.com",
		Sex:       "F",
		Birth:     "1993-01-03",
	})
	db.Create(&models.Voter{
		StudentId: "201200003",
		Password:  "201200003",
		Name:      "김영삼",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0003",
		Address:   "강남구",
		Email:     "ekf@gmail.com",
		Sex:       "M",
		Birth:     "1993-10-03",
	})
	db.Create(&models.Voter{
		StudentId: "201200004",
		Password:  "201200004",
		Name:      "김영사",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0004",
		Address:   "강남구",
		Email:     "dnwi@gmail.com",
		Sex:       "M",
		Birth:     "1992-03-01",
	})
	db.Create(&models.Voter{
		StudentId: "201200005",
		Password:  "201200005",
		Name:      "김영오",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-0000-0005",
		Address:   "강남구",
		Email:     "efkj@gmail.com",
		Sex:       "M",
		Birth:     "1991-12-04",
	})
	db.Create(&models.Voter{
		StudentId: "201200006",
		Password:  "201200006",
		Name:      "김영육",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-3920-1294",
		Address:   "서초구",
		Email:     "jwql@gmail.com",
		Sex:       "F",
		Birth:     "1994-09-29",
	})
	db.Create(&models.Voter{
		StudentId: "201200007",
		Password:  "201200007",
		Name:      "안지수",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-1427-9434",
		Address:   "강남구",
		Email:     "jsl@gmail.com",
		Sex:       "M",
		Birth:     "1994-11-29",
	})
	db.Create(&models.Voter{
		StudentId: "201200008",
		Password:  "201200008",
		Name:      "강형욱",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-2132-4223",
		Address:   "은평구",
		Email:     "flfa@gmail.com",
		Sex:       "M",
		Birth:     "1993-05-14",
	})
	db.Create(&models.Voter{
		StudentId: "201200009",
		Password:  "201200009",
		Name:      "강형욱",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-2214-6323",
		Address:   "수지구",
		Email:     "kgwq@gmail.com",
		Sex:       "M",
		Birth:     "1993-03-18",
	})
	db.Create(&models.Voter{
		StudentId: "201200010",
		Password:  "201200010",
		Name:      "전하수",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-3342-6223",
		Address:   "도봉구",
		Email:     "gk@gmail.com",
		Sex:       "F",
		Birth:     "1995-11-24",
	})
}
