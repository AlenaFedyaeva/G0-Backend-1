package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

const username = "name"

func (s *Server) UserPrintHello(w http.ResponseWriter, r *http.Request) {
	requestName, ok := r.URL.Query()[username]

	if !ok {
		sendError(w, http.StatusBadRequest, "invalid format")
		s.log.Error().Msg("invalid format")
	}
	if len(requestName) > 0 {
		str := fmt.Sprintln("Hi,", requestName[0])
		w.Write([]byte(str))
	} else {
		w.Write([]byte("hi,nobody"))
	}

}

type ErrCode int

type ErrMsg struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func sendError(w http.ResponseWriter, errorCode ErrCode, message string) {
	userErr := ErrMsg{
		Code:    int(errorCode),
		Message: message,
	}
	w.WriteHeader(int(errorCode))

	err := json.NewEncoder(w).Encode(userErr)
	if err != nil {
		log.Error().Err(err).Msg("Error encode error when sending error")
	}
}
