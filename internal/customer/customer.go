package customer

import (
	"database/sql"
	"fmt"
)

type CustomerService struct {
	db *sql.DB
}

type Customer struct {
	id    int
	name  string
	phone string
}

type CustomerInterface interface {
	GetCustomers() ([]Customer, error)
}

func NewCustomerService(db *sql.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (c *CustomerService) GetCustomers() ([]Customer, error) {

	rows, _ := c.db.Query("SELECT id, name, phone FROM  customer")
	for rows.Next() {
		var id int
		var name string
		var phone string
		rows.Scan(&id, &name, &phone)
		fmt.Println(id, name, phone)
	}

	customers := []Customer{
		{id: 1, name: "John", phone: "0712345678"},
		{id: 2, name: "Jane", phone: "0712345678"},
		{id: 3, name: "Jack", phone: "0712345678"},
	}

	return customers, nil

}

/*func filter() {
	//By Country
	//By Validity of Phone Number
	//Pagination
}

func resolveCountry() {

}

func validateNumber() {

}
*/
