package models

import (
	"CloudSystem/database"
	"CloudSystem/queries"
	"fmt"

	"github.com/google/uuid"
)

type VirtualMachine struct {
	HostName        string `json:"hostname" binding:"required"`
	IsActive        bool   `json:"isActive" binding:"required"`
	Identifier      string `json:"identifier"`
	OperatingSystem string `json:"operatingSystem" binding:"required"`
	UserIdentifier  string
}

func NewVirtualMachine(hostName, operatingSystem, UserIdentifier string, isActive bool) *VirtualMachine {
	return &VirtualMachine{
		HostName:        hostName,
		IsActive:        isActive,
		OperatingSystem: operatingSystem,
		UserIdentifier:  UserIdentifier,
	}

}
func (vm *VirtualMachine) SetIdentifier(identifier string) {
	vm.Identifier = identifier
}
func (vm *VirtualMachine) InsertVirtualMachine() (*string, error) {

	sqlStatement := `INSERT INTO virtualmachines (hostname, is_active,operating_system,user_id,identifier) VALUES ($1, $2,$3,$4,$5) RETURNING identifier`
	// Prepare the SQL statement.
	var id string
	errQuery := database.DB.QueryRow(sqlStatement, vm.HostName, vm.IsActive, vm.OperatingSystem, 3, uuid.New()).Scan(&id)
	if errQuery != nil {
		fmt.Println(errQuery)
		return nil, errQuery
	}

	fmt.Println("Data inserted successfully!")

	return &id, nil
}
func GetAllVirtualMachines() ([]VirtualMachine, error) {
	rows, err := database.DB.Query(queries.QueryMap["getAllVirtualMachines"])
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var virtualMachines []VirtualMachine

	for rows.Next() {
		var vm VirtualMachine
		err := rows.Scan(&vm.HostName, &vm.IsActive, &vm.OperatingSystem, &vm.Identifier, &vm.UserIdentifier)

		if err != nil {
			return nil, err
		}

		virtualMachines = append(virtualMachines, vm)
	}

	return virtualMachines, nil
}
func GetVirtualMachineByID(identifier string) (*VirtualMachine, error) {
	row := database.DB.QueryRow(queries.QueryMap["getVirtualMachineById"], identifier)
	var virtualMachine VirtualMachine
	err := row.Scan(&virtualMachine.Identifier, &virtualMachine.IsActive, &virtualMachine.OperatingSystem, &virtualMachine.HostName, &virtualMachine.UserIdentifier)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return &virtualMachine, nil
}

// func (vm *VirtualMachine) UpdateVirtualMachineActiveState() error {

// 	query := `
// 	UPDATE virtualmachines
// 	SET is_active = ?
// 	WHERE identifier = ?
// 	`
// 	stmt, err := database.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	// _, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err
// }
