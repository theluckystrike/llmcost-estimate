# llmcost-estimate

`llmcost` is a tiny Go package for estimating what an LLM conversation
costs, given token counts and per-million-token pricing for models like
Claude, GPT-4, Gemini, Llama and Mistral.

It exists because I kept re-doing the same arithmetic when I wanted to
[compare LLM pricing](https://kickllm.com/) across providers — the
KickLLM cost calculator does this interactively in the browser, and this
package is the same computation as a small, testable Go function.

## Usage

```go
gpt4 := llmcost.Pricing{Name: "gpt-4o", Input: 2.50, Output: 10.00}
c := gpt4.Cost(llmcost.Usage{InputTokens: 12000, OutputTokens: 800})
fmt.Printf("$%.4f\n", c)
```
