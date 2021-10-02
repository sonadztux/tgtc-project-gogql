package controller

import (
	"encoding/json"
	"log"
	"net/http"

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
