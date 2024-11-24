package handler

import (
	"encoding/json"
	"fmt"
	"github.com/daddydemir/assistant/pkg/broker"
	"log/slog"
	"net/http"
)

func messages(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	queueName := requestData["queueName"]
	content := requestData["content"]

	activeBroker := broker.GetActiveBroker()
	err = activeBroker.WriteMessage(fmt.Sprintf("%v", queueName), fmt.Sprintf("%v", content))
	if err != nil {
		slog.Error("messages", "error", err)
	}
}
