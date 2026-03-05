package service

import (
	"fmt"
	"github.com/JacobBananalDev/resume-match/internal/analyze"
	"github.com/JacobBananalDev/resume-match/internal/ai"
)

/*
AnalyzeService is responsible for coordinating
resume analysis operations.

Services sit between HTTP handlers and the
business logic layer.

Why?

Handlers should not contain business logic.
Services orchestrate application behavior.
*/
type AnalyzeService struct {
	embeddingService *ai.EmbeddingService
}

/*
NewAnalyzeService creates a new service instance.

Later this could inject dependencies like:
- AI models
- databases
- caching layers
*/
func NewAnalyzeService() *AnalyzeService {
	return &AnalyzeService{
		embeddingService: ai.NewEmbeddingService(),
	}
}

/*
AnalyzeResume performs resume analysis.

This method calls the analysis engine.
Later we can extend this with:
- AI recommendations
- database storage
- analytics
*/
func (s *AnalyzeService) AnalyzeResume(
	resumeText string,
	jobDescription string,
	roleTitle string,
) (analyze.AnalysisResult, []string) {

	// Run keyword-based analysis
	result := analyze.Analyze(resumeText, jobDescription)

	   // Run semantic AI similarity
    semanticScore := s.ComputeSemanticScore(resumeText, jobDescription)

	fmt.Println("Semantic similarity:", semanticScore)

    // Convert semantic score (0–1) to percentage
    semanticPercent := int(semanticScore * 100)

    // Combine scores
    finalScore := int(float64(result.Score)*0.6 + float64(semanticPercent)*0.4)

    // Override score with combined score
    result.Score = finalScore

	recommendations := analyze.GenerateRecommendations(
		result.MissingSkills,
		roleTitle,
	)

	return result, recommendations
}

/*
ComputeSemanticScore measures how semantically similar
a resume is to a job description using AI embeddings.

How it works:

1) Convert the resume text into an embedding vector
2) Convert the job description into an embedding vector
3) Compare the two vectors using cosine similarity

Embeddings represent the *meaning* of text rather than
just the exact words used.

This allows ResumeMatch to detect relationships like:

"built scalable APIs" ≈ "backend distributed systems"
"PostgreSQL experience" ≈ "SQL database skills"

The cosine similarity score ranges from:

1.0 → nearly identical meaning
0.8 → strong match
0.5 → moderate match
0.0 → unrelated

This semantic score can later be combined with
the keyword-based score to produce a final
resume-job compatibility score.

If the embedding provider is unavailable (ex: no API key),
we return a fallback value so the API continues to function.

This keeps the architecture ready for AI while allowing
development without external dependencies.
*/
func (s *AnalyzeService) ComputeSemanticScore(
	resumeText string,
	jobDescription string,
) float64 {

	resumeEmbedding, err := s.embeddingService.GenerateEmbedding(resumeText)
	
	if err != nil {
		fmt.Println("Embedding unavailable, using fallback semantic score")
		return 0.7
	}

	jobEmbedding, err := s.embeddingService.GenerateEmbedding(jobDescription)
	if err != nil {
		fmt.Println("Embedding unavailable, using fallback semantic score")
		return 0.7
	}

	score := ai.CosineSimilarity(resumeEmbedding, jobEmbedding)

	return score
}