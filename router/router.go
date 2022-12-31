package router

import (
	"go-backend-test/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/customers", controller.GetAllCustomers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/customer/{cst_id}", controller.GetDetailCustomer).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/customer", controller.SaveCustomer).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/customer/{cst_id}", controller.UpdateCustomer).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/customer/{cst_id}", controller.DeleteCustomer).Methods("DELETE", "OPTIONS")

	return router
}
