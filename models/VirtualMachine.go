package models

import (
	"CloudSystem/database"
	"fmt"
)

type VirtualMachine struct {
	HostName        string `json:"hostname" binding:"required"`
	IsActive        bool   `json:"isActive" binding:"required"`
	Identifier      string
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
func (vm *VirtualMachine) InsertVirtualMachine() (int, error) {

	sqlStatement := `INSERT INTO virtualmachines (hostname, is_active,operating_system) VALUES ($1, $2,$3) RETURNING identifier`
	// Prepare the SQL statement.
	var id int
	errQuery := database.DB.QueryRow(sqlStatement, vm.HostName, vm.IsActive, vm.OperatingSystem).Scan(&id)
	if errQuery != nil {
		fmt.Println(errQuery)
		return 0, errQuery
	}

	fmt.Println("Data inserted successfully!")

	return id, nil
}
func GetAllVirtualMachines() ([]VirtualMachine, error) {

	query := "SELECT virtualmachines.hostname,virtualmachines.is_active,virtualmachines.operating_system, virtualmachines.identifier,users.identifier FROM virtualmachines LEFT JOIN users ON users.id = virtualmachines.user_id"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var virtualMachines []VirtualMachine

	for rows.Next() {
		var vm VirtualMachine
		err := rows.Scan(&vm.HostName, &vm.IsActive, &vm.OperatingSystem, &vm.Identifier,&vm.UserIdentifier)

		if err != nil {
			return nil, err
		}

		virtualMachines = append(virtualMachines, vm)
	}

	return virtualMachines, nil
}
