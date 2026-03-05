package analyze

import "testing"

/*
TestBasicSkillMatching verifies that
the analyzer correctly matches simple skills.
*/
func TestBasicSkillMatching(t *testing.T) {

	resume := "Go developer using Docker and AWS"
	job := "Looking for Go developer with Docker and AWS"

	result := Analyze(resume, job)

	if result.Score != 100 {
		t.Errorf("expected score 100, got %d", result.Score)
	}

	if len(result.MatchedSkills) != 3 {
		t.Errorf("expected 3 matched skills, got %d", len(result.MatchedSkills))
	}
}

/*
TestMissingSkills verifies missing skills are detected.
*/
func TestMissingSkills(t *testing.T) {

	resume := "Go developer with AWS"
	job := "Go developer with AWS and Kubernetes"

	result := Analyze(resume, job)

	if result.Score >= 100 {
		t.Errorf("expected score less than 100")
	}

	if len(result.MissingSkills) == 0 {
		t.Errorf("expected missing skills but found none")
	}
}

/*
TestSynonymNormalization verifies
golang is normalized to go.
*/
func TestSynonymNormalization(t *testing.T) {

	resume := "Experienced Golang engineer"
	job := "Looking for Go developer"

	result := Analyze(resume, job)

	if result.Score != 100 {
		t.Errorf("expected golang to match go")
	}
}

/*
TestSemanticMatching verifies
semantic skills like postgres → sql work.
*/
func TestSemanticMatching(t *testing.T) {

	resume := "Built backend systems using PostgreSQL"
	job := "Looking for engineer with SQL experience"

	result := Analyze(resume, job)

	if result.Score == 0 {
		t.Errorf("expected semantic matching between postgres and sql")
	}
}