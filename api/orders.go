package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arshabbir/utils/model"
)

func (s *server) handleOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order service"))

	seassonId := r.Header.Get("sessionid")
	if seassonId == "" {
		err := &model.ApiError{ErrorCode: http.StatusUnauthorized, Message: "Invalied Session ID"}
		json.NewEncoder(w).Encode(err)
		return
	}

	req := &model.OrderRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		s.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "orderservice", Message: "Error decoding the POST body"})
		return
	}
	s.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "orderservice", Message: fmt.Sprintf("%v", req)})

	resp := &model.OrderResponse{StatusCode: http.StatusOK}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		s.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "orderservice", Message: "Error encoding the order response "})
		return
	}

	// To be continued
}
