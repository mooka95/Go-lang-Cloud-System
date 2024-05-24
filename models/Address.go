package models

import (
	"CloudSystem/database"
	"CloudSystem/queries"

	"github.com/google/uuid"
)

type Address struct {
	City       string `json:"city" binding:"required"`
	Street     string `json:"street" binding:"required"`
	Country    string `json:"country" binding:"required"`
	Identifier string
	UserId     int64
}

func NewAddress(city, street, country string) *Address {
	return &Address{
		City:    city,
		Street:  country,
		Country: country,
	}

}

func (address *Address) CreateAddress() (*string, error) {
	var id string
	err := database.DB.QueryRow(queries.QueryAddressMap["insertAddress"], address.City, address.Street, address.Country, uuid.New(), address.UserId).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
// func GetAllAddresses() ([]Address, error) {
// 	rows, err := database.DB.Query(queries.QueryFirewallMap["getAllFirewalls"])
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var firewalls []Firewall

// 	for rows.Next() {
// 		var firewall Firewall
// 		err := rows.Scan(&firewall.Name, &firewall.Identifier, &firewall.UserIdentifier)

// 		if err != nil {
// 			return nil, err
// 		}

// 		firewalls = append(firewalls, firewall)
// 	}

// 	return firewalls, nil
// }