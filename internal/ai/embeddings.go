package ai

import (
	"fmt"
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

/*
EmbeddingService generates vector embeddings
for text using OpenAI models.
*/
type EmbeddingService struct {
	client *openai.Client
}

func NewEmbeddingService() *EmbeddingService {

	apiKey := os.Getenv("OPENAI_API_KEY")

	fmt.Println("Loaded API key:", apiKey)
	
	client := openai.NewClient(apiKey)

	return &EmbeddingService{
		client: client,
	}
}

/*
GenerateEmbedding converts text into a vector representation.
*/
func (s *EmbeddingService) GenerateEmbedding(text string) ([]float32, error) {

	req := openai.EmbeddingRequest{
		Input: []string{text},
		Model: "text-embedding-3-small",
	}

	resp, err := s.client.CreateEmbeddings(context.Background(), req)
	if err != nil {
		fmt.Println("OpenAI embedding error:", err)
		return nil, err
	}

	// Make sure response actually contains embeddings
	if len(resp.Data) == 0 {
		fmt.Println("No embedding returned from OpenAI")
		return nil, fmt.Errorf("empty embedding response")
	}

	embedding := resp.Data[0].Embedding

	fmt.Println("Embedding length:", len(embedding))

	return embedding, nil
}