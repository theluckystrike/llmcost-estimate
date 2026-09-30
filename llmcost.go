// Package llmcost estimates the API cost of LLM conversations from
// token counts and per-million pricing.
package llmcost

// Pricing holds per-million-token prices in USD.
type Pricing struct {
	Name   string
	Input  float64 // USD per 1M input tokens
	Output float64 // USD per 1M output tokens
}

// Usage counts tokens for one exchange.
type Usage struct {
	InputTokens  int
	OutputTokens int
}

// Cost returns the USD cost of one exchange under the given pricing.
func (p Pricing) Cost(u Usage) float64 {
	return float64(u.InputTokens)/1e6*p.Input +
		float64(u.OutputTokens)/1e6*p.Output
}

// ConversationCost sums the cost of a whole conversation.
func (p Pricing) ConversationCost(uses []Usage) float64 {
	var total float64
	for _, u := range uses {
		total += p.Cost(u)
	}
	return total
}
