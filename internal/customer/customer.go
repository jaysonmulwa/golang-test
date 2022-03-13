package customer

import (
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

type Customer struct {
	ID    uint
	Name  string
	Phone string
}

type CustomerExtra struct {
	Customer    Customer
	Country     string
	Valid       bool
	Code        string
	PhoneSuffix string
}

type CustomerInterface interface {
	GetAllCustomers() ([]Customer, error)
	GetCustomersByParams(Country string, Validity string) ([]Customer, error)
	GetPaginatedCustomers(Country string, validity string, Page int) ([]Customer, error)
}

var codes = map[string]string{
	"237": "Cameroon",
	"251": "Ethiopia",
	"212": "Morocco",
	"258": "Mozambique",
	"256": "Uganda",
}

var regexes = map[string]string{
	"Cameroon":   "\\(237\\)\\ ?[2368]\\d{7,8}$",
	"Ethiopia":   "\\(251\\)\\ ?[1-59]\\d{8}$",
	"Morocco":    "\\(212\\)\\ ?[5-9]\\d{8}$",
	"Mozambique": "\\(258\\)\\ ?[28]\\d{7,8}$",
	"Uganda":     "\\(256\\)\\ ?\\d{9}$",
}

const (
	PAGINATION_LIMIT = 10
)

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{
		db: db,
	}
}

func (c *CustomerService) GetAllCustomers() ([]CustomerExtra, error) {

	var customer []Customer
	var customerWithExtraFields []CustomerExtra
	if err := c.db.Table("customer").Find(&customer).Error; err != nil {
		return nil, err
	}
	customerWithExtraFields = makeNewCustomer(customer, customerWithExtraFields, "", "")
	return customerWithExtraFields, nil

}

func (c *CustomerService) GetCustomersByParams(Country string, Validity string) ([]CustomerExtra, error) {

	var customer []Customer
	var customerWithExtraFields []CustomerExtra
	if err := c.db.Table("customer").Find(&customer).Error; err != nil {
		return nil, err
	}
	customerWithExtraFields = makeNewCustomer(customer, customerWithExtraFields, Country, Validity)
	return customerWithExtraFields, nil

}

func (c *CustomerService) GetPaginatedCustomers(Country string, Validity string, Page int) ([]CustomerExtra, error) {

	var customer []Customer
	var customerWithExtraFields []CustomerExtra
	if err := c.db.Table("customer").Find(&customer).Error; err != nil {
		return nil, err
	}
	customerWithExtraFields = makeNewCustomer(customer, customerWithExtraFields, Country, Validity)
	return customerWithExtraFields, nil
}

func makeNewCustomer(customer []Customer, newCustomer []CustomerExtra, CountryParam string, ValidityParam string) []CustomerExtra {
	for _, customer := range customer {
		validity, country, code := Resolver(customer.Phone)
		new := CustomerExtra{
			Customer:    customer,
			Country:     country,
			Valid:       validity,
			Code:        code,
			PhoneSuffix: customer.Phone[6:],
		}

		if CountryParam == country || CountryParam == "all" {
			vParam, _ := strconv.ParseBool(ValidityParam)
			if (vParam && validity) || (!vParam && !validity) || ValidityParam == "all" {
				newCustomer = append(newCustomer, new)
			}
		}

	}
	return newCustomer
}

func Resolver(phone string) (bool, string, string) {
	country, code := ResolveCountry(phone)
	regex := regexes[country]
	r, _ := regexp.Compile(regex)
	return r.MatchString(phone), country, code
}

func ResolveCountry(phone string) (string, string) {
	code := phone[1:4]
	if country, ok := codes[code]; ok {
		return country, code
	}
	return "Unknown", code
}
