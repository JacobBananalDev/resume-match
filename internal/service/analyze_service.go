package service

import (
	"github.com/JacobBananalDev/resume-match/internal/analyze"
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
}

/*
NewAnalyzeService creates a new service instance.

Later this could inject dependencies like:
- AI models
- databases
- caching layers
*/
func NewAnalyzeService() *AnalyzeService {
	return &AnalyzeService{}
}

/*
AnalyzeResume performs resume analysis.

This method calls the analysis engine.
Later we can extend this with:
- AI recommendations
- database storage
- analytics
*/
func (s *AnalyzeService) AnalyzeResume(resumeText string, jobDescription string) analyze.AnalysisResult {

	result := analyze.Analyze(resumeText, jobDescription)

	return result
}