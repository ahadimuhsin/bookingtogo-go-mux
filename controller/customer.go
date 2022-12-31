package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "fmt"

	// "fmt"
	"go-backend-test/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    []models.Customer `json:"data"`
}

func SaveCustomer(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf(dd.Dump(r))
	var customer models.Customer
	// dd.Dump(r.Body)
	err := r.ParseForm();
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// name := 
	// fmt.Println(name)
	errs := json.NewDecoder(r.Body).Decode(&customer)

	if errs != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	fmt.Println(len(customer.Family))
	//tambah customer dulu
	insertID := models.SaveCustomer(customer)

	//tambah family
	var family models.FamilyList

	family.CustID = insertID

	for i:=0;i<len(customer.Family); i++{
		customer.Family[i].CustID = insertID
		
		//save
		_ = models.SaveFamily(customer.Family[i])
	}

	res := response{
		ID:      int(insertID),
		Message: "Data Customer Telah Ditambahkan",
	}

	json.NewEncoder(w).Encode(res)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	customers, err := models.GetAllCustomers()

	if err != nil {
		log.Fatalf("Tidak Bisa mengambil data. %v", err)
	}

	var response Response

	response.Status = 1
	response.Message = "Success"
	response.Data = customers

	//return responsenya
	json.NewEncoder(w).Encode(response)
}

func GetDetailCustomer(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	//konversi id dari string ke int
	id, err := strconv.Atoi(params["cst_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	//ambil data dengan parameter id
	customer, err := models.GetCustomer(int64(id))

	if err != nil {
		log.Fatalf("Tidak Bisa mengambil data. %v", err)
	}

	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	//konversi id dari string ke int
	id, err := strconv.Atoi(params["cst_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	//ambil data dari modelnya
	var customer models.Customer

	err = json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	//panggil method update Customer
	updatedCustomer := models.UpdateCustomer(int64(id), customer)

	msg := fmt.Sprintf("Data Customer Berhasil Diperbarui. Jumlah Yang Diupdate %v baris", updatedCustomer)

	res := response{
		ID: id,
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	//konversi id dari string ke int
	id, err := strconv.Atoi(params["cst_id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	deletedCustomer := models.DeleteCustomer(int64(id))

	msg := fmt.Sprintf("Data Customer Berhasil Dihapus. Total Data yang dihapus %v baris", deletedCustomer)

	res := response{
		ID: id,
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}