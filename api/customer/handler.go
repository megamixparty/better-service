package customer

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"customer/api/middlewares"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var req CustomerDetail
	json.NewDecoder(r.Body).Decode(&req)
	id, err := CreateCustomer(middlewares.GetDB(r.Context()), req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CreateResponse{CustomerID: id})
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := GetCustomer(middlewares.GetDB(r.Context()), id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	var cp CustomerPagination

	cp.Offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))
	cp.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))

	if cp.Offset < 0 {
		cp.Offset = 0
	}

	if cp.Limit <= 0 {
		cp.Limit = 10
	} else if cp.Limit > 50 {
		cp.Limit = 50
	}

	res, err := ListCustomer(middlewares.GetDB(r.Context()), &cp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.Header().Set("Pagination-Offset", strconv.Itoa(cp.Offset))
	w.Header().Set("Pagination-Limit", strconv.Itoa(cp.Limit))
	w.Header().Set("Pagination-Total", strconv.Itoa(cp.Total))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := GetCustomer(middlewares.GetDB(r.Context()), id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	err = DeleteCustomer(middlewares.GetDB(r.Context()), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
