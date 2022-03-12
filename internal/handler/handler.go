package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jaysonmulwa/jumia/internal/customer"
)

type Handler struct {
	Router   *mux.Router
	Customer *customer.CustomerService
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(customer *customer.CustomerService) *Handler {
	return &Handler{
		Customer: customer,
	}
}

func (h *Handler) SetupRoutes() {

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/customers/{country}/{validity}", h.GetCustomers).Methods("GET")
	h.Router.HandleFunc("/customers/{country}/{validity}/{page}", h.GetPaginatedCustomers).Methods("GET")
}

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	country := vars["country"]
	validity := vars["validity"]

	customers, err := h.Customer.GetCustomers(country, validity)
	if err != nil {
		sendErrorResponse(w, "Error", err)
		return
	}

	if err = sendOkResponse(w, customers); err != nil {
		panic(err)
	}

}

func (h *Handler) GetPaginatedCustomers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	country := vars["country"]
	validity := vars["validity"]
	page := vars["page"]

	_page, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse requested Page", err)
	}

	customers, err := h.Customer.GetPaginatedCustomers(country, validity, int(_page))
	if err != nil {
		sendErrorResponse(w, "Error", err)
		return
	}

	if err = sendOkResponse(w, customers); err != nil {
		panic(err)
	}

}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
