package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/toky03/oracle-swap-demo/model"
	"github.com/toky03/oracle-swap-demo/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type oracleService interface {
	ReadWalletNames() ([]string, error)
	ReadFunds(walletId string) (model.WalletDto, error)
	AddFunds(walletId string, payload model.FundsDto) error
	ClosePool(walletId string, payload model.CloseDto) error
	CreatePool(string, model.FundsDto) error
	RemoveFunds(string, model.RemoveDto) error
	Swap(string, model.SwapDto) error
}

func main() {
	handler := &handler{
		oracleService: createService(),
	}
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))
	api.HandleFunc("/wallets", handler.ReadWallets).Methods("GET")
	api.HandleFunc("/{walletId}/add", handler.AddFunds).Methods("PUT")
	api.HandleFunc("/{walletId}/close", handler.ClosePool).Methods("PUT")
	api.HandleFunc("/{walletId}/create", handler.CreatePool).Methods("POST")
	api.HandleFunc("/{walletId}/funds", handler.ReadFunds).Methods("GET")
	api.HandleFunc("/{walletId}/remove", handler.RemoveFunds).Methods("PUT")
	api.HandleFunc("/{walletId}/swap", handler.Swap).Methods("PUT")
	api.HandleFunc("/{walletId}", handler.OptionsHandler).Methods("OPTIONS")
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})
	log.Println("Start Middleware on Port 3001")
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}

type handler struct {
	oracleService oracleService
}

func createService() oracleService {
	timeout := os.Getenv("STARTUP_TIMEOUT_SECONDS")
	if timeout == "" {
		timeout = "60"
	}
	timeoutSeconds, err := strconv.Atoi(timeout)
	if err != nil {
		log.Fatalf("Startuptime %s cannot be converted to int", timeout)
	}
	expirationTime := time.Now().Add(time.Duration(timeoutSeconds) * time.Second)
	for {
		if time.Now().After(expirationTime) {
			log.Fatalf("Timeout of %s seconds reached", timeout)
		}
		s, err := service.CreateOracleService()
		if err == nil {
			log.Println("Configuration files found")
			return s
		}
		log.Println("Files not found retry in 2 seconds")
		time.Sleep(2*time.Second)
	}

}

func (h *handler) OptionsHandler(_ http.ResponseWriter, _ *http.Request) {
	return
}

func (h *handler) ReadWallets(w http.ResponseWriter, r *http.Request) {
	wallets, err := h.oracleService.ReadWalletNames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, marshalErr := json.Marshal(wallets)
	writeContent(w, marshalErr, js)
}

func (h *handler) ReadFunds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	funds, err := h.oracleService.ReadFunds(walletId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, marshalErr := json.Marshal(funds)
	writeContent(w, marshalErr, js)

}


func writeContent(w http.ResponseWriter, marshalErr error, data []byte) {
	if marshalErr != nil {
		http.Error(w, marshalErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *handler) AddFunds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	var payload model.FundsDto
	decoderError := json.NewDecoder(r.Body).Decode(&payload)
	if decoderError != nil {
		http.Error(w, decoderError.Error(), http.StatusBadRequest)
		return
	}
	err := h.oracleService.AddFunds(walletId, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) ClosePool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	var payload model.CloseDto
	decoderError := json.NewDecoder(r.Body).Decode(&payload)
	if decoderError != nil {
		http.Error(w, decoderError.Error(), http.StatusBadRequest)
		return
	}
	err := h.oracleService.ClosePool(walletId, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) CreatePool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	var payload model.FundsDto
	decoderError := json.NewDecoder(r.Body).Decode(&payload)
	if decoderError != nil {
		http.Error(w, decoderError.Error(), http.StatusBadRequest)
		return
	}
	err := h.oracleService.CreatePool(walletId, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}


func (h *handler) RemoveFunds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	var payload model.RemoveDto
	decoderError := json.NewDecoder(r.Body).Decode(&payload)
	if decoderError != nil {
		http.Error(w, decoderError.Error(), http.StatusBadRequest)
		return
	}
	err := h.oracleService.RemoveFunds(walletId, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) Swap(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletId := vars["walletId"]
	var payload model.SwapDto
	decoderError := json.NewDecoder(r.Body).Decode(&payload)
	if decoderError != nil {
		http.Error(w, decoderError.Error(), http.StatusBadRequest)
		return
	}
	err := h.oracleService.Swap(walletId, payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
