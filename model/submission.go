package model

import (
	"ce-boostup-backend/db"
	"ce-boostup-backend/judge0"
	"fmt"
	"log"
	"strconv"
)

//Submission a model for submission
type Submission struct {
	SubmissionID int     `json:"submission_id" form:"submission_id"`
	UserID       int     `json:"user_id" form:"user_id"`
	ProblemID    int     `json:"problem_id" form:"problem_id"`
	LanguageID   int     `json:"language_id" form:"language_id"`
	Src          string  `json:"src" form:"src"`
	SubmittedAt  string  `json:"submitted_at" form:"submitted_at"`
	Score        int     `json:"score" form:"score"`
	Runtime      int     `json:"runtime" form:"runtime"`
	MemoryUsage  float32 `json:"memory_usage" form:"memory_usage"`
}

//NewSubmission create a new submission
func NewSubmission(userID int, problemID int, languageID int, src string) error {

	score := 0
	runtime := 0.0
	memory := 0

	testcase, err := SpecificTestcaseWithID(userID)
	if err != nil {
		return err
	}

	for i := range testcase {
		result := judge0.Submit(src, testcase[i].Input, testcase[i].Output) //empty string is for testcase in the future
		memory += result.Memory
		runtime += stringToFloat(result.Time)
	}

	length := len(testcase)
	runtime = runtime / float64(length)
	memory = memory / length

	statement := `INSERT INTO submission (usr_id,problem_id,lang_id,src,score,runtime,memory_usage) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err = db.DB.Exec(statement, userID, problemID, languageID, src, score, runtime, memory)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func stringToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
		return 0.0
	}
	return value
}
