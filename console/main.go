package main

import (
        "bufio"
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
        "sort"
        "strings"
        "time"
)

const (
        NvidiaBaseURL     = "https://integrate.api.nvidia.com/v1"
        NvidiaAPIKey      = "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"
        NovaBaseURL       = "https://api.nova.amazon.com/v1"
        NovaAPIKey        = "fdbdcf6a-a2f3-4201-9488-89f94ea528a3"
        GatewayAPIKey     = "connect"
        MaxToolIterations = 10
)

var modelAliases = map[string]string{
        "Bielik-11b":          "speakleash/bielik-11b-v2.6-instruct",
        "Mistral-Small-4":     "mistralai/mistral-small-4-119b-2603",
        "DeepSeek-V3.1":       "deepseek-ai/deepseek-v3.1",
        "Kimi-K2":             "moonshotai/kimi-k2-instruct",
        "Amazon-Nova-2-lite-v1": "nova-2-lite-v1",
        "Minimax-m2.5":        "minimaxai/minimax-m2.5",
        "GLM-4.7":             "z-ai/glm4.7",
        "GPT-OSS-120b":        "openai/gpt-oss-120b",
        "Step-3.5-Flash":      "stepfun-ai/step-3.5-flash",
        "Qwen-3.5":            "qwen/qwen3.5-122b-a10b",
        "Kimi-K2.5":           "moonshotai/kimi-k2.5",
}

// Modele korzystające z Amazon Nova API zamiast NVIDIA
var novaModels = map[string]bool{
        "nova-2-lite-v1": true,
}

// Modele z wyłączonym thinking
var noThinkingModels = map[string]bool{
        "deepseek-ai/deepseek-v3.1": true,
}

func getProviderConfig(modelID string) (baseURL, apiKey string) {
        if novaModels[modelID] {
                return NovaBaseURL, NovaAPIKey
        }
        return NvidiaBaseURL, NvidiaAPIKey
}

// --- STRUKTURY ---

type Message struct {
        Role       string      `json:"role"`
        Content    interface{} `json:"content"`
        ToolCallID string      `json:"tool_call_id,omitempty"`
        ToolCalls  interface{} `json:"tool_calls,omitempty"`
        Name       string      `json:"name,omitempty"`
}

type ToolFunction struct {
        Name        string                 `json:"name"`
        Description string                 `json:"description,omitempty"`
        Parameters  map[string]interface{} `json:"parameters,omitempty"`
        Endpoint    string                 `json:"x-endpoint,omitempty"`
}

type Tool struct {
        Type     string       `json:"type"`
        Function ToolFunction `json:"function"`
}

type ChatRequest struct {
        Model       string        `json:"model"`
        Messages    []Message     `json:"messages"`
        Stream      *bool         `json:"stream,omitempty"`
        Tools       []Tool        `json:"tools,omitempty"`
        ToolChoice  interface{}   `json:"tool_choice,omitempty"`
        Temperature *float64      `json:"temperature,omitempty"`
        MaxTokens   *int          `json:"max_tokens,omitempty"`
}

type AccumToolCall struct {
        Index int
        ID    string
        Name  string
        Args  string
}

// --- POMOCNICZE ---

func resolveModel(requested string) string {
        if full, ok := modelAliases[requested]; ok {
                return full
        }
        return requested
}

func findTool(tools []Tool, name string) *Tool {
        for _, t := range tools {
                if t.Function.Name == name {
                        return &t
                }
        }
        return nil
}

// executeToolCall wykonuje HTTP POST do x-endpoint narzędzia
func executeToolCall(tool *Tool, argsJSON string) string {
        if tool == nil || tool.Function.Endpoint == "" {
                return fmt.Sprintf(`{"error":"brak x-endpoint dla narzędzia %s"}`, tool.Function.Name)
        }

        var args interface{}
        json.Unmarshal([]byte(argsJSON), &args)
        body, _ := json.Marshal(args)

        client := &http.Client{Timeout: 30 * time.Second}
        resp, err := client.Post(tool.Function.Endpoint, "application/json", bytes.NewReader(body))
        if err != nil {
                return fmt.Sprintf(`{"error":"%s"}`, err.Error())
        }
        defer resp.Body.Close()
        result, _ := io.ReadAll(resp.Body)
        return string(result)
}

// --- UPSTREAM CALL (non-streaming, zbiera pełną odpowiedź) ---

func callUpstream(modelID string, messages []Message, tools []Tool, toolChoice interface{}, temperature *float64, maxTokens *int) (map[string]interface{}, error) {
        payload := map[string]interface{}{
                "model":    modelID,
                "messages": messages,
                "stream":   false,
        }
        if noThinkingModels[modelID] {
                payload["thinking"] = false
        }
        if temperature != nil {
                payload["temperature"] = *temperature
        }
        if maxTokens != nil {
                payload["max_tokens"] = *maxTokens
        }
        if len(tools) > 0 {
                payload["tools"] = tools
                if toolChoice != nil {
                        payload["tool_choice"] = toolChoice
                } else {
                        payload["tool_choice"] = "auto"
                }
        }

        baseURL, apiKey := getProviderConfig(modelID)
        body, _ := json.Marshal(payload)
        req, _ := http.NewRequest("POST", baseURL+"/chat/completions", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+apiKey)

        client := &http.Client{Timeout: 120 * time.Second}
        resp, err := client.Do(req)
        if err != nil {
                return nil, err
        }
        defer resp.Body.Close()

        var result map[string]interface{}
        if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
                return nil, err
        }
        return result, nil
}

// --- STREAMING UPSTREAM (ostatnia odpowiedź) ---

func streamUpstream(w http.ResponseWriter, modelID string, messages []Message, tools []Tool, toolChoice interface{}, temperature *float64, maxTokens *int, clientModel string) {
        payload := map[string]interface{}{
                "model":    modelID,
                "messages": messages,
                "stream":   true,
        }
        if noThinkingModels[modelID] {
                payload["thinking"] = false
        }
        if temperature != nil {
                payload["temperature"] = *temperature
        }
        if maxTokens != nil {
                payload["max_tokens"] = *maxTokens
        }
        if len(tools) > 0 {
                payload["tools"] = tools
                if toolChoice != nil {
                        payload["tool_choice"] = toolChoice
                } else {
                        payload["tool_choice"] = "auto"
                }
        }

        baseURL, apiKey := getProviderConfig(modelID)
        body, _ := json.Marshal(payload)
        req, _ := http.NewRequest("POST", baseURL+"/chat/completions", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+apiKey)

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
                http.Error(w, err.Error(), 502)
                return
        }
        defer resp.Body.Close()

        flusher, _ := w.(http.Flusher)
        scanner := bufio.NewScanner(resp.Body)
        accum := make(map[int]*AccumToolCall)

        for scanner.Scan() {
                line := scanner.Text()
                if !strings.HasPrefix(line, "data: ") || line == "data: [DONE]" {
                        fmt.Fprint(w, line+"\n\n")
                        if flusher != nil {
                                flusher.Flush()
                        }
                        continue
                }

                var chunk map[string]interface{}
                if err := json.Unmarshal([]byte(strings.TrimPrefix(line, "data: ")), &chunk); err != nil {
                        continue
                }

                choices, ok := chunk["choices"].([]interface{})
                if !ok || len(choices) == 0 {
                        continue
                }

                choice := choices[0].(map[string]interface{})
                delta, _ := choice["delta"].(map[string]interface{})
                if delta == nil {
                        continue
                }
                finishReason := choice["finish_reason"]

                if tcs, ok := delta["tool_calls"].([]interface{}); ok {
                        for _, tcVal := range tcs {
                                tc := tcVal.(map[string]interface{})
                                idx := int(tc["index"].(float64))
                                acc, exists := accum[idx]
                                if !exists {
                                        acc = &AccumToolCall{Index: idx}
                                        if id, ok := tc["id"].(string); ok {
                                                acc.ID = id
                                        }
                                        accum[idx] = acc
                                }
                                if fn, ok := tc["function"].(map[string]interface{}); ok {
                                        if name, ok := fn["name"].(string); ok {
                                                acc.Name += name
                                        }
                                        if args, ok := fn["arguments"].(string); ok {
                                                acc.Args += args
                                        }
                                }
                        }
                        continue
                }

                if (finishReason == "tool_calls" || finishReason == "function_call") && len(accum) > 0 {
                        var keys []int
                        for k := range accum {
                                keys = append(keys, k)
                        }
                        sort.Ints(keys)

                        finalTools := []map[string]interface{}{}
                        for _, k := range keys {
                                a := accum[k]
                                finalTools = append(finalTools, map[string]interface{}{
                                        "index": a.Index, "id": a.ID, "type": "function",
                                        "function": map[string]interface{}{"name": a.Name, "arguments": a.Args},
                                })
                        }

                        response := map[string]interface{}{
                                "id": chunk["id"], "object": "chat.completion.chunk", "created": chunk["created"],
                                "model": clientModel,
                                "choices": []map[string]interface{}{{
                                        "index":         0,
                                        "delta":         map[string]interface{}{"role": "assistant", "tool_calls": finalTools},
                                        "finish_reason": "tool_calls",
                                }},
                        }
                        jsonBytes, _ := json.Marshal(response)
                        fmt.Fprintf(w, "data: %s\n\n", string(jsonBytes))
                        if flusher != nil {
                                flusher.Flush()
                        }
                        accum = make(map[int]*AccumToolCall)
                        continue
                }

                // podmień model na alias klienta
                chunk["model"] = clientModel
                out, _ := json.Marshal(chunk)
                fmt.Fprintf(w, "data: %s\n\n", string(out))
                if flusher != nil {
                        flusher.Flush()
                }
        }

        fmt.Fprint(w, "data: [DONE]\n\n")
        if flusher != nil {
                flusher.Flush()
        }
}

// --- GŁÓWNY HANDLER ---

func handleChat(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodOptions {
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")
                w.WriteHeader(http.StatusNoContent)
                return
        }

        providedKey := authHeader(r)
        if !isValidAPIKey(providedKey) {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
        }

        var req ChatRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                http.Error(w, "Bad Request", http.StatusBadRequest)
                return
        }

        clientModel := req.Model
        modelID := resolveModel(req.Model)
        messages := req.Messages
        tools := req.Tools
        toolChoice := req.ToolChoice

        wantStream := req.Stream == nil || *req.Stream

        // --- PĘTLA AGENTYCZNA ---
        // Jeśli są narzędzia z x-endpoint, automatycznie wykonujemy pętle tool calls.
        // Każda iteracja: non-streaming call → sprawdź tool_calls → wykonaj → dodaj wyniki → powtórz.
        // Ostatnia odpowiedź (bez tool_calls) jest streamowana/zwracana do klienta.

        hasAutoExec := false
        if len(tools) > 0 {
                for _, t := range tools {
                        if t.Function.Endpoint != "" {
                                hasAutoExec = true
                                break
                        }
                }
        }

        if hasAutoExec {
                for i := 0; i < MaxToolIterations; i++ {
                        result, err := callUpstream(modelID, messages, tools, toolChoice, req.Temperature, req.MaxTokens)
                        if err != nil {
                                http.Error(w, err.Error(), 502)
                                return
                        }

                        choices, ok := result["choices"].([]interface{})
                        if !ok || len(choices) == 0 {
                                break
                        }

                        choice := choices[0].(map[string]interface{})
                        message, _ := choice["message"].(map[string]interface{})
                        finishReason, _ := choice["finish_reason"].(string)

                        // dodaj wiadomość asystenta do historii
                        assistantMsg := Message{Role: "assistant"}
                        if content, ok := message["content"]; ok && content != nil {
                                assistantMsg.Content = content
                        }
                        if tcs, ok := message["tool_calls"]; ok && tcs != nil {
                                assistantMsg.ToolCalls = tcs
                        }
                        messages = append(messages, assistantMsg)

                        if finishReason != "tool_calls" && finishReason != "function_call" {
                                // brak tool calls — zwróć wynik klientowi
                                w.Header().Set("Content-Type", "application/json")
                                w.Header().Set("Access-Control-Allow-Origin", "*")
                                result["model"] = clientModel
                                json.NewEncoder(w).Encode(result)
                                return
                        }

                        // wykonaj wszystkie tool calls
                        tcList, _ := message["tool_calls"].([]interface{})
                        for _, tcVal := range tcList {
                                tc, _ := tcVal.(map[string]interface{})
                                if tc == nil {
                                        continue
                                }
                                tcID, _ := tc["id"].(string)
                                fn, _ := tc["function"].(map[string]interface{})
                                if fn == nil {
                                        continue
                                }
                                fnName, _ := fn["name"].(string)
                                fnArgs, _ := fn["arguments"].(string)

                                tool := findTool(tools, fnName)
                                toolResult := executeToolCall(tool, fnArgs)

                                messages = append(messages, Message{
                                        Role:       "tool",
                                        Content:    toolResult,
                                        ToolCallID: tcID,
                                        Name:       fnName,
                                })
                        }
                }

                // max iteracji osiągnięte — ostatnia próba bez narzędzi
                result, err := callUpstream(modelID, messages, nil, nil, req.Temperature, req.MaxTokens)
                if err != nil {
                        http.Error(w, err.Error(), 502)
                        return
                }
                w.Header().Set("Content-Type", "application/json")
                w.Header().Set("Access-Control-Allow-Origin", "*")
                result["model"] = clientModel
                json.NewEncoder(w).Encode(result)
                return
        }

        // --- NORMALNY TRYB (bez auto-exec): stream do klienta ---
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("X-Accel-Buffering", "no")
        w.Header().Set("Cache-Control", "no-cache")

        if !wantStream {
                // klient nie chce streamu — zbierz odpowiedź i zwróć JSON
                result, err := callUpstream(modelID, messages, tools, toolChoice, req.Temperature, req.MaxTokens)
                if err != nil {
                        http.Error(w, err.Error(), 502)
                        return
                }
                w.Header().Set("Content-Type", "application/json")
                result["model"] = clientModel
                json.NewEncoder(w).Encode(result)
                return
        }

        streamUpstream(w, modelID, messages, tools, toolChoice, req.Temperature, req.MaxTokens, clientModel)
}

func handleModels(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodOptions {
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")
                w.WriteHeader(http.StatusNoContent)
                return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        var data []map[string]interface{}
        now := time.Now().Unix()
        for alias := range modelAliases {
                data = append(data, map[string]interface{}{
                        "id":       alias,
                        "object":   "model",
                        "created":  now,
                        "owned_by": "nvidia",
                })
        }
        json.NewEncoder(w).Encode(map[string]interface{}{"object": "list", "data": data})
}

func detectBaseURL(r *http.Request) string {
        return "https://ai-gateway-lrv4.onrender.com/v1"
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
        setCORSHeaders(w, r)
        if r.Method == http.MethodOptions {
                w.WriteHeader(http.StatusNoContent)
                return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
                "base_url": detectBaseURL(r),
        })
}

func main() {
        port := os.Getenv("PORT")
        if port == "" {
                port = "10000"
        }

        mux := http.NewServeMux()
        mux.HandleFunc("/v1/chat/completions", handleChat)
        mux.HandleFunc("/v1/models", handleModels)
        mux.HandleFunc("/v1/config", handleRoot)
        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(`{"status":"ok"}`))
        })

        go func() {
                for {
                        refreshKeyCache()
                        time.Sleep(60 * time.Second)
                }
        }()
        log.Printf("API Gateway running on :%s", port)
        if err := http.ListenAndServe(":"+port, mux); err != nil {
                log.Fatalf("Server error: %v", err)
        }
}
