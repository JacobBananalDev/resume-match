package analyze

import (
	"regexp"
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
	"javascript",
	"typescript",
	"sql",
	"api",
	"distributed",
	"database",
}

/*
skillSynonyms maps common alternate names to canonical skills.

Example:
golang → go
k8s → kubernetes
*/
var skillSynonyms = map[string]string{
	"golang": "go",
	"k8s": "kubernetes",
	"js": "javascript",
	"ts": "typescript",
}

/*
semanticSkills maps related technologies
to broader concepts.
*/
var semanticSkills = map[string][]string{
	"postgres": {"sql", "database"},
	"mysql": {"sql", "database"},
	"redis": {"database", "cache"},
	"grpc": {"api"},
	"rest": {"api"},
	"microservices": {"distributed", "distributed systems"},
	"docker": {"containers"},
	"kubernetes": {"containers", "orchestration"},
}

/*
normalizeToken converts synonyms to canonical skill names.
*/
func normalizeToken(token string) string {

	if normalized, exists := skillSynonyms[token]; exists {
		return normalized
	}

	return token
}

/*
tokenize converts text into normalized words.

Steps:
1) lowercase text
2) remove punctuation
3) split into tokens
*/
func tokenize(text string) []string {

	text = strings.ToLower(text)

	// Remove punctuation
	reg := regexp.MustCompile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	rawTokens := strings.Fields(text)

	var tokens []string

	for _, token := range rawTokens {

		normalized := normalizeToken(token)

		tokens = append(tokens, normalized)
	}

	return tokens
}

/*
extractSkills finds which skills exist in text.

It returns a map for fast lookup.
*/
func extractSkills(text string) map[string]bool {

	tokens := tokenize(text)

	tokenSet := make(map[string]bool)

	for _, token := range tokens {
		tokenSet[token] = true

		// semantic expansions
		if related, exists := semanticSkills[token]; exists {

			for _, r := range related {
				tokenSet[r] = true
			}

		}
	}

	foundSkills := make(map[string]bool)

	for _, skill := range skillBank {
		if tokenSet[skill] {
			foundSkills[skill] = true
		}
	}

	return foundSkills
}

/*
Analyze performs resume vs job comparison.

Steps:
1) extract job skills
2) extract resume skills
3) compare
4) compute score
*/
func Analyze(resumeText string, jobDescription string) AnalysisResult {

	jobSkills := extractSkills(jobDescription)
	resumeSkills := extractSkills(resumeText)

	var matched []string
	var missing []string

	for skill := range jobSkills {

		if resumeSkills[skill] {
			matched = append(matched, skill)
		} else {
			missing = append(missing, skill)
		}
	}

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