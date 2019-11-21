package operation

import (
	"encoding/json"
	"net/http"
	"strconv"

	hp "redcoins-api/internal"
)

// Create : get user handler
func Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	operationType := r.Form.Get("operation_type")
	amount, err := strconv.ParseFloat(r.Form.Get("amount"), 64)

	if operationType != "purchase" && operationType != "sale" {
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid operation."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	if err != nil {
		res := hp.JSONStandardResponse{Code: 406, Message: "Invalid amount."}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	err = CreateOperation(Operation{
		uuid:          ``,
		opertaionType: operationType,
		amount:        amount,
		userUUID:      "31c68ff7-0bc9-11ea-900c-9829a6e582d0"})

	if err != nil {
		res := hp.JSONStandardResponse{Code: 406, Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("res")
	return
}
