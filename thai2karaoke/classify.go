package thai2karaoke

import (
	"encoding/json"
	"math"

	_ "embed"
)

//go:embed weights.json
var rawWeights []byte

// net holds simple network weights.
type net struct {
	W1     [][]float64 `json:"w1"`
	B1     []float64   `json:"b1"`
	W2     [][]float64 `json:"w2"`
	B2     []float64   `json:"b2"`
	Labels []string    `json:"labels"`
}

var model net

func init() {
	_ = json.Unmarshal(rawWeights, &model)
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

// oneHotEncode converts runes to a simple numeric vector.
func oneHotEncode(word string) []float64 {
	rs := []rune(word)
	out := make([]float64, 3)
	for i := 0; i < len(out); i++ {
		if i < len(rs) {
			out[i] = float64(rs[i]) / 1000.0
		} else {
			out[i] = 0
		}
	}
	return out
}

// classify returns a karaoke tag for the word.
func classify(word string) string {
	input := oneHotEncode(word)
	h := make([]float64, len(model.W1))
	for i := range model.W1 {
		sum := model.B1[i]
		for j, w := range model.W1[i] {
			sum += w * input[j]
		}
		h[i] = sigmoid(sum)
	}
	out := make([]float64, len(model.W2))
	for i := range model.W2 {
		sum := model.B2[i]
		for j, w := range model.W2[i] {
			sum += w * h[j]
		}
		out[i] = sigmoid(sum)
	}
	idx := 0
	max := out[0]
	for i := 1; i < len(out); i++ {
		if out[i] > max {
			max = out[i]
			idx = i
		}
	}
	if idx < len(model.Labels) {
		return model.Labels[idx]
	}
	return ""
}
