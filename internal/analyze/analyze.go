package analyze

import (
	"strings"
)

/*
AnalysisResult represents the output of our analysis engine.

This is separate from HTTP response models because
business logic should not depend on HTTP structures.
*/
type AnalysisResult struct {
	Score         int
	MatchedSkills []string
	MissingSkills []string
}

/*
skillBank is a small list of skills we recognize.

Later this can be replaced by:
- a database
- an AI embedding model
- a skills taxonomy
*/
var skillBank = []string{
	"go",
	"docker",
	"kubernetes",
	"aws",
	"postgres",
	"react",
	"python",
	"terraform",
}

/*
Analyze compares the resume text against the job description.

Algorithm (very simple for now):

1) Detect which skills appear in the job description
2) Check which of those appear in the resume
3) Calculate score
*/
func Analyze(resumeText string, jobDescription string) AnalysisResult {

	// Normalize text to lowercase for easier comparison
	resume := strings.ToLower(resumeText)
	job := strings.ToLower(jobDescription)

	var jobSkills []string
	var matched []string
	var missing []string

	// Find which skills appear in the job description
	for _, skill := range skillBank {
		if strings.Contains(job, skill) {
			jobSkills = append(jobSkills, skill)
		}
	}

	// Check which required skills appear in the resume
	for _, skill := range jobSkills {

		if strings.Contains(resume, skill) {
			matched = append(matched, skill)
		} else {
			missing = append(missing, skill)
		}

	}

	// Calculate score
	score := 0

	if len(jobSkills) > 0 {
		score = int(float64(len(matched)) / float64(len(jobSkills)) * 100)
	}

	return AnalysisResult{
		Score:         score,
		MatchedSkills: matched,
		MissingSkills: missing,
	}
}