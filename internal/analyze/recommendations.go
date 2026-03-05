package analyze

/*
GenerateRecommendations creates resume improvement suggestions
based on missing skills and general resume best practices.
*/
func GenerateRecommendations(
	missingSkills []string,
	roleTitle string,
) []string {

	var recommendations []string

	// Suggest adding missing skills
	for _, skill := range missingSkills {

		recommendations = append(
			recommendations,
			"Consider adding experience with "+skill+" if applicable.",
		)

	}

	// General resume advice
	recommendations = append(
		recommendations,
		"Include measurable achievements (latency improvements, system scale, etc).",
	)

	recommendations = append(
		recommendations,
		"Highlight backend architecture or system design experience.",
	)

	recommendations = append(
		recommendations,
		"Ensure your resume clearly lists core technologies used in production.",
	)

	return recommendations
}