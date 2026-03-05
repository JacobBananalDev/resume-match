package ai

import (
	"fmt"
	"math"
)
/*
CosineSimilarity measures similarity between two vectors.

Result:
1.0 = identical
0.0 = unrelated
*/
func CosineSimilarity(a, b []float32) float64 {

	if len(a) == 0 || len(b) == 0 {
		fmt.Println("CosineSimilarity: empty vector")
		return 0
	}

	if len(a) != len(b) {
		fmt.Println("CosineSimilarity: vector length mismatch")
		return 0
	}

	var dotProduct float64
	var normA float64
	var normB float64

	for i := range a {
		dotProduct += float64(a[i] * b[i])
		normA += float64(a[i] * a[i])
		normB += float64(b[i] * b[i])
	}

	if normA == 0 || normB == 0 {
		fmt.Println("CosineSimilarity: zero vector")
		return 0
	}

	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}