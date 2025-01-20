package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type GraphQL struct {
	Schema *graphql.Schema
	Server *SSEServer
}

const subscriptionTimeout = 30 * time.Minute

type params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func (h *GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	operationName := r.URL.Query().Get("operationName")
	variables := make(map[string]interface{})
	if vars := r.URL.Query().Get("variables"); vars != "" {
		if err := json.Unmarshal([]byte(vars), &variables); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	params := params{
		Query:         query,
		OperationName: operationName,
		Variables:     variables,
	}

	if query != "" && strings.HasPrefix(query, "subscription") {
		h.serveSubscription(w, r, &params)
	} else {
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		h.serveDefault(w, r, &params)
	}
}

func (h *GraphQL) serveDefault(w http.ResponseWriter, r *http.Request, params *params) {
	response := h.Schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, statusCode := checkError(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

func (h *GraphQL) serveSubscription(w http.ResponseWriter, r *http.Request, params *params) {

	clientID := len(h.Server.Clients) + 1
	client := &SSEClient{
		ID:     clientID,
		Stream: make(chan string),
	}
	h.Server.AddClient <- client

	ctx, cancel := context.WithTimeout(r.Context(), subscriptionTimeout)
	h.callOnClose(w, cancel, client)

	c, err := h.Schema.Subscribe(ctx, params.Query, params.OperationName, params.Variables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/event-stream")
	w.Header().Set("cache-control", "no-cache")

	go func() {
		flusher, ok := w.(http.Flusher)
		if !ok {
			log.Printf("error: not a flusher\n")
		}
		for message := range client.Stream {
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		}
	}()

	var eventCount int
	for r := range c {
		response := r.(*graphql.Response)
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if hasError, statusCode := checkError(response); hasError {
			w.WriteHeader(statusCode)
		}
		client.Stream <- string(responseJSON)
		eventCount++
	}

	if eventCount == 0 {
		w.WriteHeader(http.StatusOK)
	}
}

func checkError(response *graphql.Response) (hasError bool, statusCode int) {
	statusCode = http.StatusOK

	if len(response.Errors) > 0 {
		hasError = true

		for _, e := range response.Errors {
			if e.Message == "unauthorized" {
				statusCode = http.StatusUnauthorized
				break
			}
		}
	}

	return hasError, statusCode
}

func (h *GraphQL) callOnClose(w http.ResponseWriter, cb func(), client *SSEClient) {
	if cn, ok := w.(http.CloseNotifier); !ok {
		log.Printf("error: not a close notifier\n")
	} else {
		closeChan := cn.CloseNotify()
		go func(closeChan <-chan bool, cn http.CloseNotifier, cb func()) {
			<-closeChan
			cb()
			h.Server.RemoveClient <- client.ID
		}(closeChan, cn, cb)
	}
}
