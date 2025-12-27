# useful-genkit

A collection of utility packages for building AI applications using Firebase Genkit for Go. This project serves as a practical example integrating various AI models (Ollama, Google AI) and implementing workflows like translation.

## Project Structure

### 📁 `models/`
AI Model Management and Abstraction Layer

#### Intent
- Unified management of models from various AI providers (Ollama, Google AI)
- Consistent API for model initialization and retrieval
- Encapsulation of model-specific characteristics and constraints

#### Contents
- **`ollama.go`**: Ollama Local Model Management
  - Supports GPT-OSS 20B, Gemma3, Qwen3 models
  - Provides model definition and retrieval functions
  - Configured for multi-turn conversation and tool usage
- **`gemini.go`**: Google AI Model Management
  - Supports Gemini 2.5 Pro/Flash series
  - Supports Gemma 3 series (4B, 12B, 27B)
  - Unified interface for model access

### 📁 `flows/`
Workflow Definitions containing Business Logic

#### Intent
- Structuring complex AI tasks into reusable flows
- Ensuring input/output type safety
- Including error handling and validation logic

#### Contents
- **`translation.go`**: Translation Flow Implementation
  - Translation with configurable source/target languages and domain
  - Structured input (`TranslationInput`) and output (`TranslationOutput`)
  - Integration of prompt rendering and model invocation

### 📁 `prompts/`
Prompt Templates for AI Model Interaction

#### Intent
- Management of reusable and parameterized prompts
- Designing prompts with context-specific expertise
- Separation of prompt rendering and message generation

#### Contents
- **`translation.go`**: Translation-specific Prompts
  - Sets up the role of a domain-specific expert translator
  - Emphasizes preservation of original nuances
  - Dynamic prompt generation via template variables

### 📁 `tools/`
External Tools available for AI Models

#### Intent
- Extending model capabilities with real-world data or tasks
- Ensuring safe tool invocation with structured input/output
- Providing reusable utility functions

#### Contents
- **`get_current_time.go`**: Current Time Retrieval Tool
  - Supports AI tasks requiring time information
  - JSON-serializable structured output
- **`find_usage.go`**: Code Usage Finder Tool
  - Finds usages of functions, methods, or types in the codebase using [`ast-grep`](https://ast-grep.github.io)
  - Returns the full code definition of the matched symbol including file path and line number
  - **Note**: Requires `ast-grep` (`sg` command) to be installed on the system

### 📁 `logic/`
Advanced Generation Logic and Helper Functions

#### Intent
- Abstraction of complex AI interaction patterns into reusable functions
- Advanced workflows combining tool usage and data generation
- Conversation history management and context maintenance

#### Contents
- **`generate_data_with_tool.go`**: Data Generation with Tool Usage
  - Generates final response based on conversation history after tool invocation
  - Supports various output formats via generic types
  - Maintains context through pre-processed conversation history

## Key Features

- **Multi-Model Support**: Integration of Ollama local models and Google AI cloud models
- **Type Safety**: Safe AI interaction leveraging Go's strong type system
- **Modularity**: Improved maintainability with separated package structure for each function
- **Extensibility**: Structure designed for easy addition of new models, flows, and tools
- **Error Handling**: Robust error handling and validation logic included

## Usage Examples

This project provides practical patterns for developing AI applications using the Firebase Genkit Go SDK. It establishes a foundation that can be extended from translation services to various AI-based applications.
