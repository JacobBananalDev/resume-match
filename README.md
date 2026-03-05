# 🚀 ResumeMatch

ResumeMatch is a production-style Go backend API that analyzes resumes against job descriptions and returns a match score, matched skills, missing skills, and improvement recommendations.

The project demonstrates backend architecture, skill extraction, semantic matching design, containerization, and CI/CD automation — similar to systems used in modern ATS and AI-assisted hiring tools.

## 🧠 Key Features
- 📄 Resume vs Job Description analysis
- 🧩 Skill extraction engine
- 🔁 Skill normalization (golang → go, k8s → kubernetes)
- 🧠 Semantic skill expansion (postgres → sql → database)
- 📊 Match scoring (0–100)
- ⚠️ Missing skill detection
- 💡 Resume improvement recommendations
- 🤖 AI embedding architecture (OpenAI-ready)
- 🚦 Rate limiting middleware
- 📜 Structured request logging
- 🐳 Docker containerization
- 🧱 Docker Compose orchestration
- ⚙️ CI/CD with GitHub Actions
- 📦 Automated Docker image publishing (GHCR)

## 🏗 Architecture

The project follows a clean layered backend architecture.
```bash 
Client
   │
   ▼
HTTP API (Chi Router)
   │
   ▼
Service Layer
   │
   ▼
Analyzer Engine
   │
   ├─ Skill Extraction
   ├─ Skill Normalization
   ├─ Semantic Expansion
   │
   ▼
Scoring Engine
```
## Repository structure:
```bash
resume-match
│
├── cmd/
│   └── api/
│
├── internal/
│   ├── analyze/       # Skill extraction + matching logic
│   ├── ai/            # Embedding + similarity engine
│   ├── service/       # Business logic orchestration
│   ├── httpapi/       # API handlers
│   └── middleware/    # Logging + rate limiting
│
├── Dockerfile
├── docker-compose.yml
└── .github/workflows  # CI/CD pipelines
```
## ⚙️ Tech Stack
| Category |	Technology |
| -------- | ----------- |
| Language |🐹 Go |
| HTTP Framework |	Chi Router |
| Containerization |	Docker |
| Local Dev |	Docker Compose |
| CI/CD |	GitHub Actions |
| Container Registry |	GitHub Container Registry |
| AI Integration |	OpenAI Embeddings (architecture ready) |
| Testing |	Go testing |
## 🚦 Running the Project
### Option 1 — Run with Docker Compose (Recommended)
```bash
docker compose up --build
```
API will start at:
```bash
http://localhost:8080
```
### Option 2 — Run Locally
```bash
go run ./cmd/api
```
## 📡 API Endpoints
Health Check
```bash
GET /health
```
Response:
```bash
{
  "status": "ok"
}
```
Resume Analysis
```bash
POST /analyze
```
Request
```bash
{
  "companyName": "Stripe",
  "roleTitle": "Backend Engineer",
  "resumeText": "Built scalable APIs using Go, Docker and PostgreSQL",
  "jobDescriptionText": "Backend engineer experienced with Go, SQL, Kubernetes and AWS"
}
```
Response
```bash
{
  "score": 75,
  "matchedSkills": [
    "go",
    "docker",
    "postgres"
  ],
  "missingSkills": [
    "kubernetes",
    "aws"
  ],
  "recommendations": [
    "Add Kubernetes experience",
    "Highlight cloud infrastructure work"
  ]
}
```
## 🧠 Skill Matching Engine

ResumeMatch performs multiple layers of analysis:

### 1️⃣ Tokenization

Text is normalized and converted to tokens.

### 2️⃣ Skill Extraction

Skills are detected using a curated skill taxonomy.

### 3️⃣ Skill Normalization

Synonyms are mapped to canonical skills.

Example:

golang → go
postgresql → postgres
k8s → kubernetes
### 4️⃣ Semantic Expansion

Related technologies expand into broader concepts.

Example:
```bash
postgres → sql → database
grpc → api
microservices → distributed
```
### 5️⃣ Match Scoring

Score calculation:
```bash
score = matched_skills / job_required_skills
```
## 🤖 AI Semantic Matching (Architecture Ready)

The project includes an AI embedding layer designed for semantic similarity.

Planned capability:
```bash
Resume Text
      │
      ▼
Embedding Model
      │
      ▼
Vector Similarity
      │
      ▼
Semantic Match Score
```

Embedding providers supported:
- OpenAI Embeddings
- Ollama (local models)
- Sentence Transformers

## 🐳 Docker

Build the container manually:
```bash
docker build -t resumematch .
```
Run the container:
```bash
docker run -p 8080:8080 resumematch
```
## ⚙️ CI/CD Pipeline

GitHub Actions automatically:
### 1️⃣ Runs tests
### 2️⃣ Builds the Go service
### 3️⃣ Builds the Docker image
### 4️⃣ Publishes the image to GitHub Container Registry

### Example pull command:
```bash
docker pull ghcr.io/jacobbananaldev/resume-match:latest
```
### Run the container:
```bash
docker run -p 8080:8080 ghcr.io/jacobbananaldev/resume-match:latest
```
## 🧪 Testing

Run all tests:
```bash
go test ./...
```
## 📈 Future Improvements
- Resume section parsing (skills / experience / projects)
- AI-powered skill extraction from job descriptions
- Vector search for semantic resume matching
- Resume PDF parsing
- Frontend dashboard for resume feedback
- PostgreSQL persistence for resume scoring history

## 👨‍💻 Author

Jacob Bananal

Software Engineer

GitHub:
https://github.com/JacobBananalDev
