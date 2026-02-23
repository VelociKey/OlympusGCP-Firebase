package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	firebasev1 "OlympusGCP-Firebase/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/firebase/v1"
	"OlympusGCP-Firebase/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/firebase/v1/firebasev1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type FirebaseServer struct{}

func (s *FirebaseServer) CreateUser(ctx context.Context, req *connect.Request[firebasev1.CreateUserRequest]) (*connect.Response[firebasev1.CreateUserResponse], error) {
	slog.Info("CreateUser", "email", req.Msg.Email)
	return connect.NewResponse(&firebasev1.CreateUserResponse{Uid: "user-abc"}), nil
}

func (s *FirebaseServer) SetDocument(ctx context.Context, req *connect.Request[firebasev1.SetDocumentRequest]) (*connect.Response[firebasev1.SetDocumentResponse], error) {
	slog.Info("SetDocument", "collection", req.Msg.Collection, "id", req.Msg.DocId)
	return connect.NewResponse(&firebasev1.SetDocumentResponse{Uid: req.Msg.DocId}), nil
}

func main() {
	server := &FirebaseServer{}
	mux := http.NewServeMux()
	path, handler := firebasev1connect.NewFirebaseServiceHandler(server)
	mux.Handle(path, handler)

	port := "8099" // From genesis.json
	slog.Info("FirebaseManager starting", "port", port)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}
