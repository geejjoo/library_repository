package responder

import (
	"encoding/json"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/infrastructure/response"
	"go.uber.org/zap"
	"net/http"
)

type Responder interface {
	OutputJSON(w http.ResponseWriter, responseData interface{})
	ErrorUnauthorized(w http.ResponseWriter, err error)
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorForbidden(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type Respond struct {
	logger *zap.Logger
}

func NewResponder(logger *zap.Logger) Responder {
	return &Respond{logger: logger}
}

func (r *Respond) OutputJSON(w http.ResponseWriter, responseData interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		r.logger.Error(errors.ResponderEncodeError, zap.Error(err))
		return
	}
}

func (r *Respond) ErrorUnauthorized(w http.ResponseWriter, err error) {
	r.logger.Error(errors.ResponderStatusUnauthorized, zap.Error(err))
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
	if err != nil {
		r.logger.Error(errors.ResponderErrorUnauthorized, zap.Error(err))
		return
	}
}

func (r *Respond) ErrorBadRequest(w http.ResponseWriter, err error) {
	r.logger.Error(errors.ResponderStatusBadRequest, zap.Error(err))
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
	if err != nil {
		r.logger.Error(errors.ResponderErrorBadRequest, zap.Error(err))
		return
	}
}

func (r *Respond) ErrorForbidden(w http.ResponseWriter, err error) {
	r.logger.Error(errors.ResponderStatusForbidden, zap.Error(err))
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
	if err != nil {
		r.logger.Error(errors.ResponderErrorForbidden, zap.Error(err))
		return
	}
}

func (r *Respond) ErrorInternal(w http.ResponseWriter, err error) {
	r.logger.Error(errors.ResponderStatusInternal, zap.Error(err))
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(response.Response{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
	if err != nil {
		r.logger.Error(errors.ResponderErrorInternal, zap.Error(err))
		return
	}
}
