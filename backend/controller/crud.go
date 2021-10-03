package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sonadztux/tgtc-project-gogql/backend/model"
	"github.com/sonadztux/tgtc-project-gogql/backend/views"
)

func GetAllCoupons(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		resp views.APIResponse
		err  error
	)
	// some input validations here
	//
	//
	//

	// proceed to the main service
	resp.Data, err = model.GetAllCoupons()

	// construct api response json
	if err != nil {
		resp.Error = err.Error()
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonResponse)
}

func GetCouponByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// variable declarations
	var (
		resp   views.APIResponse
		err    error
		userID int
	)
	// some input validations here
	//
	//
	//

	r.ParseForm()
	vars := mux.Vars(r)
	userID, err = strconv.Atoi(vars["id"])

	// input validated
	if err == nil {
		// proceed to the main service
		resp.Data, err = model.GetCouponByUserID(userID)
	}

	// construct api response json
	if err != nil {
		resp.Error = err.Error()
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}

	w.Write(jsonResponse)
}
