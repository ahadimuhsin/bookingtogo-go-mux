package models

import (
	"database/sql"
	"fmt"
	"go-backend-test/config"
	"log"
)

type Customer struct {
	CustID        int64  `json:"id"`
	NationalityID int64  `json:"nationality_id"`
	CstName       string `json:"name"`
	CstDob        string `json:"date_of_birth"`
	CstPhone      string `json:"phone_number"`
	CstEmail      string `json:"email"`
	Family        []FamilyList
}

type FamilyList struct {
	FlID		int64  `json:"id"`
	CustID     int64  `json:"cst_id"`
	FlRelation string `json:"relation"`
	FlName     string `json:"name"`
	FlDob      string `json:"date_of_birth"`
}

func SaveCustomer(customer Customer) int64 {
	db := config.CreateConnection()

	//untuk tutup koneksi di akhir proses
	defer db.Close()

	queryInsert := `INSERT INTO customer (nationality_id, 
		cst_name, cst_dob, cst_phone, cst_email)
		VALUES ($1, $2, $3, $4, $5) RETURNING cst_id`

	var cust_id int64
	//Scan funtion akan menyimpan hasil id ke dalam variabel ID
	err := db.QueryRow(queryInsert, customer.NationalityID, customer.CstName, customer.CstDob, customer.CstPhone, customer.CstEmail).Scan(&cust_id)

	if err != nil {
		log.Fatalf("Failed execute query. %v", err)
	}

	fmt.Printf("Insert Data Single Record %v", cust_id)

	return cust_id
}

func SaveFamily(family FamilyList) int64 {
	db := config.CreateConnection()
	//untuk tutup koneksi di akhir proses
	defer db.Close()

	queryInsert := `INSERT INTO family_list 
	(cst_id, fl_relation, fl_name, fl_dob)
		VALUES ($1, $2, $3, $4) RETURNING fl_id`

	var fl_id int64
	//Scan funtion akan menyimpan hasil id ke dalam variabel ID
	err := db.QueryRow(queryInsert, family.CustID, family.FlRelation, family.FlName, family.FlDob).Scan(&fl_id)

	if err != nil {
		log.Fatalf("Failed execute query. %v", err)
	}

	fmt.Printf("Insert Data Single Record %v", fl_id)

	return fl_id
}

// ambil satu customer
func GetAllCustomers() ([]Customer, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var customers []Customer

	// kita buat select query
	sqlStatement := `SELECT * FROM customer`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var customer Customer

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&customer.CustID, &customer.NationalityID,
			&customer.CstName, &customer.CstDob, &customer.CstPhone, &customer.CstEmail)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		// masukkan kedalam slice customers
		customers = append(customers, customer)

	}

	// return empty customer atau jika error
	return customers, err
}

// mengambil satu customer
func GetCustomer(id int64) (Customer, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var customer Customer
	var family FamilyList

	// buat sql query
	sqlStatement := `SELECT * FROM customer JOIN family_list ON customer.cst_id = family_list.cst_id WHERE customer.cst_id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&customer.CustID, &customer.NationalityID,
		&customer.CstName, &customer.CstDob, &customer.CstPhone, &customer.CstEmail, &family.FlID, &family.CustID, &family.FlRelation, &family.FlName, &family.FlDob)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return customer, nil
	case nil:
		return customer, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return customer, err
}

// update user in the DB
func UpdateCustomer(id int64, customer Customer) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE customer SET nationality_id=$2, cst_name=$3, cst_dob=$4, cst_phone=$5, cst_email=$6 WHERE cst_id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, customer.NationalityID, customer.CstName, customer.CstDob, customer.CstPhone, customer.CstEmail)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func DeleteCustomer(id int64) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM customer WHERE cst_id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
