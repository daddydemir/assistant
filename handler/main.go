package handler

import (
	"github.com/daddydemir/assistant/pkg/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
)

func route() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(setLogging)

	r.HandleFunc("/stat", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK"))
	})

	r.HandleFunc("/api/v1/messages", messages).Methods(http.MethodPost)

	handler := cors.AllowAll().Handler(r)
	return handler
}

func StartApi() {
	server := &http.Server{
		Addr:    config.Get("PORT"),
		Handler: route(),
	}

	slog.Info("server started at", "http://localhost", config.Get("PORT"))

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
