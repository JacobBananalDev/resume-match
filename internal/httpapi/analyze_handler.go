package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/JacobBananalDev/resume-match/internal/service"
)

/*
Handler struct allows us to attach dependencies
like services.

This keeps handlers clean and testable.
*/
type Handler struct {
	analyzeService *service.AnalyzeService
}

/*
NewHandler constructs a handler with dependencies.
*/
func NewHandler() *Handler {
	return &Handler{
		analyzeService: service.NewAnalyzeService(),
	}
}

/*
AnalyzeRequest represents the JSON body sent to POST /analyze.

We include:
- companyName: which company the job is for
- roleTitle: what role the candidate is applying to
- resumeText: full resume content
- jobDescriptionText: full job posting text

Example request:

{
  "companyName": "Google",
  "roleTitle": "Backend Engineer",
  "resumeText": "...",
  "jobDescriptionText": "..."
}
*/
type AnalyzeRequest struct {
	CompanyName 		string `json:"companyName"`
	RoleTitle			string `json:"roleTitle"`
	ResumeText        	string `json:"resumeText"`
	JobDescriptionText 	string `json:"jobDescriptionText"`
}

/*
AnalyzeResponse represents what our API will return.

Later this will include real scoring logic.
For now we just return a placeholder response.
*/
type AnalyzeResponse struct {
	Score         int      `json:"score"`
	MatchedSkills []string `json:"matchedSkills"`
	MissingSkills []string `json:"missingSkills"`
	Recommendations []string `json:"recommendations"`
}

/*
AnalyzeHandler handles POST /analyze requests.

Responsibilities:
1) Read JSON request body
2) Validate input
3) Call analysis logic (later)
4) Return JSON response
*/
func (h *Handler) Analyze(w http.ResponseWriter, r *http.Request) {

	// Create a variable to hold the decoded request
	var req AnalyzeRequest

	// Decode JSON request body into struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Very basic validation
	if req.CompanyName == "" || 
		req.RoleTitle == "" ||
		req.ResumeText == "" ||
		req.JobDescriptionText == "" {

		http.Error(w, "companyName, roleTitle, resumeText and jobDescriptionText are required", http.StatusBadRequest)
		return
	}

	// Call service layer
	result, recommendations := h.analyzeService.AnalyzeResume(
		req.ResumeText, 
		req.JobDescriptionText,
		req.RoleTitle)
	
	resp := AnalyzeResponse{
		Score:         result.Score,
		MatchedSkills: result.MatchedSkills,
		MissingSkills: result.MissingSkills,
		Recommendations: recommendations,
	}

	// Tell the client we are returning JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode struct → JSON response
	json.NewEncoder(w).Encode(resp)
}