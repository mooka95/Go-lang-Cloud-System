package models

import (
	"CloudSystem/database"
	"CloudSystem/queries"
	"fmt"

	"github.com/google/uuid"
)

type Firewall struct {
	Name           string `json:"name" binding:"required"`
	Identifier     string
	UserIdentifier string
	UserId         int64  `json:"userId,omitempty"`
	Id             int64 `json:"id,omitempty"`
}

func NewFirewall(name string) *Firewall {
	return &Firewall{
		Name: name,
	}

}

func (firewall *Firewall) InsertFirewall() (*string, error) {

	// Prepare the SQL statement.
	var id string
	errQuery := database.DB.QueryRow(queries.QueryFirewallMap["insertFirewall"], firewall.Name, firewall.UserId, uuid.New()).Scan(&id)
	if errQuery != nil {
		fmt.Println(errQuery)
		return nil, errQuery
	}
	return &id, nil
}
func GetAllFirewalls() ([]Firewall, error) {
	rows, err := database.DB.Query(queries.QueryFirewallMap["getAllFirewalls"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var firewalls []Firewall

	for rows.Next() {
		var firewall Firewall
		err := rows.Scan(&firewall.Id, &firewall.Name, &firewall.Identifier, &firewall.UserIdentifier)

		if err != nil {
			return nil, err
		}
	
		firewalls = append(firewalls, firewall)
	}

	return firewalls, nil
}
func GetFirewallByID(identifier string) (*Firewall, error) {
	row := database.DB.QueryRow(queries.QueryFirewallMap["getFirewallById"], identifier)
	var firewall Firewall
	err := row.Scan(&firewall.Id, &firewall.Name, &firewall.Identifier, &firewall.UserIdentifier, &firewall.UserId)
	if err != nil {
		return nil, err
	}

	return &firewall, nil
}
func (firewall *Firewall) DeleteFirewall() error {
	stmt, err := database.DB.Prepare(queries.QueryFirewallMap["deleteFirewall"])

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(firewall.Identifier)
	return err
}
func GetFirewallByNameAndUserId(userId int64, firewallName string) (*Firewall, error) {
	row := database.DB.QueryRow(queries.QueryFirewallMap["getFirewallByNameAndUserId"], firewallName, userId)
	var firewall Firewall
	err := row.Scan(&firewall.Name)
	if err != nil {
		return nil, err
	}

	return &firewall, nil
}
func (firewall *Firewall) CheckIfFirewallAttachedToVirtualMachine(vmId string) bool {
	row := database.DB.QueryRow(queries.QueryFirewallMap["getFirewallAttachedVirtualMachine"], firewall.Identifier, vmId)

	err := row.Scan(&firewall.Name)
	if err != nil {
		return false
	}

	return true
}

// func GetFirewallByIdAndName(identifier string) (*VirtualMachine, error) {
// 	row := database.DB.QueryRow(queries.QueryMap["getVirtualMachineById"], identifier)
// 	var virtualMachine VirtualMachine
// 	err := row.Scan(&virtualMachine.Identifier, &virtualMachine.IsActive, &virtualMachine.OperatingSystem, &virtualMachine.HostName, &virtualMachine.UserIdentifier)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return nil, err
// 	}

// 	return &virtualMachine, nil
// }
