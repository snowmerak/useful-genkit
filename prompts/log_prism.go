package prompts

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

const LogPrismPromptName = "LogPrismPrompt"

type LogPrismInput struct {
	Code string `json:"code"`
}

type LogPrismOutput struct {
	Code string `json:"code"`
}

func LogPrismPrompt(g *genkit.Genkit) ai.Prompt {
	return genkit.DefinePrompt(g, LogPrismPromptName, ai.WithPrompt(`You are a Go expert. Your task is to add "Prism" style logging (Span and State) to the given Go code based on the concept below.

`+logPrismConcept+`

## Instructions

1.  **Analyze the Code**: Understand the flow of the provided Go code.
2.  **Add Span Logging**:
    *   Identify functions representing units of work.
    *   Initialize `+"`Span`"+` at the start.
    *   Handle errors with `+"`span.Fail(err)`"+`.
3.  **Add State Logging**:
    *   Log state changes with `+"`StateLogger.Transition`"+`.
    *   Log data snapshots with `+"`StateLogger.Snapshot`"+`.
4.  **Context**: Ensure `+"`context.Context`"+` is used to propagate RequestID.

## Input Code

{{code}}

Return the FULL source code with logging added.`, ai.WithInputType(LogPrismInput{}), ai.WithConfig(&ai.GenerationCommonConfig{
		Temperature: 0.1,
	})))
}

const logPrismConcept = `## Prism

I designed three log types:
- **Span**: Declared at the start of a request, method, or API call, and output with ` + "`execution time`" + ` at the end.
- **State**: Outputs ` + "`change details`" + ` when a state change occurs within a request, method, or API call.
- **Context**: Adds ` + "`ServiceName`" + ` and a random ` + "`RequestID`" + ` as mandatory fields to link these logs into a single flow, providing loose association.

### Span

Similar to OpenTelemetry. It tracks the lifecycle of a single request, method, or API call. This structure allows tracking how long the entire request took and which methods or APIs took the longest during execution, enabling preemptive service improvement or infrastructure expansion.

` + "```go" + `
type Span struct {
    ServiceName string
    Scope       string
    Method      string
    RequestID   string
    StartTime   time.Time
    Payload     map[string]interface{}
}

func NewSpan(service, scope, method, requestID string) *Span {
    return &Span{
        ServiceName: service,
        Scope:       scope,
        Method:      method,
        RequestID:   requestID,
        Payload:     make(map[string]interface{}),
    }
}

// Start records the start time and returns the Span object
func (s *Span) Start() *Span {
    s.StartTime = time.Now()
    return s
}

// Complete is called when work is done: automatically calculates and records Duration
func (s *Span) Complete() {
    duration := time.Since(s.StartTime)
    // Log format: {"level":"INFO", "service":"{Service}", "scope":"{Scope}", "method":"{Method}", "request_id":"{RequestID}", "status":"completed", "duration_ms":{Duration}}
}

// Fail is called when work fails: records error cause and duration together
func (s *Span) Fail(err error) {
    duration := time.Since(s.StartTime)
    // Log format: {"level":"ERROR", "service":"{Service}", "scope":"{Scope}", "method":"{Method}", "request_id":"{RequestID}", "status":"failed", "error":"{Error}", "duration_ms":{Duration}}
}
` + "```" + `

### State

It might look a bit unique, but it exists to log the result when a state changes due to an action. These logs might not be very useful in normal times but can be valuable when an incident occurs. The code below allows logging two types of state logs:
1. **Transition**: Records state transitions. Writes about content modified while reading from DB, deleting, or filtering during logic execution.
2. **Snapshot**: Records the current state itself. Used to log the initial and final state when first executed or terminated.

` + "```go" + `
type StateLogger struct {
    ServiceName string
    Scope       string
    RequestID   string
}

func NewStateLogger(service, scope, requestID string) *StateLogger {
    return &StateLogger{
        ServiceName: service,
        Scope:       scope,
        RequestID:   requestID,
    }
}

// Transition: Explicitly records the previous state (From) and the next state (To)
// Key for identifying "where the state came from" during bug tracking
func (l *StateLogger) Transition(entity string, from, to string, reason string) {
    // Log format: {"type":"transition", "service":"{Service}", "request_id":"{RequestID}", "entity":"{Entity}", "from":"{From}", "to":"{To}", "reason":"{Reason}"}
}

// Snapshot: Dumps the entire object state at a specific point in time
// For checking data consistency or verifying data at the time of debugging
func (l *StateLogger) Snapshot(entity string, data interface{}) {
    // Log format: {"type":"snapshot", "service":"{Service}", "request_id":"{RequestID}", "entity":"{Entity}", "data":{JSON_Dump}}
}
` + "```" + `

### Context

To ensure log correlation, Go's ` + "`context.Context`" + ` is actively used. It is used not only for cancellation signals or deadline propagation but also as a container to carry the request's unique identifier (Request ID) and service identifier (Service Name).

Key features:
1. **Creation and Injection**: A UUID v7-based unique Request ID is generated and injected into the Context at the entry point of the request (e.g., HTTP Middleware).
2. **Propagation**:
    - **In-Process**: ` + "`ctx`" + ` is passed as the first argument when calling functions to share the ID between goroutines.
    - **Cross-Process**: During communication between microservices, the ID is propagated via HTTP Header (` + "`X-Request-ID`" + `), enabling Distributed Tracing.
3. **Utilization**: All loggers (Span, State, Trend) extract this ID from the Context and include it in the logs. This allows searching the entire flow of a specific request among tens of thousands of logs at once.
`
