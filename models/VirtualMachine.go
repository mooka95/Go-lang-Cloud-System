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
	UserId          int64 `json:"userId,omitempty"`
	Id              int64 `json:"id,omitempty"`
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
	var id string
	errQuery := database.DB.QueryRow(queries.QueryMap["insertVirtualMachine"], vm.HostName, vm.IsActive, vm.OperatingSystem, vm.UserId, uuid.New()).Scan(&id)
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
	err := row.Scan(&virtualMachine.Id, &virtualMachine.Identifier, &virtualMachine.IsActive, &virtualMachine.OperatingSystem, &virtualMachine.HostName, &virtualMachine.UserIdentifier, &virtualMachine.UserId)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return &virtualMachine, nil
}
func (vm *VirtualMachine) DeleteVirtualMachine(identifier string) error {
	// query := "DELETE FROM events WHERE id = ?"
	stmt, err := database.DB.Prepare(queries.QueryMap["deleteVirtualMachine"])

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(identifier)
	return err
}

func (vm *VirtualMachine) UpdateVirtualMachineActiveState(isActive bool) error {

	stmt, err := database.DB.Prepare(queries.QueryMap["updateIsActive"])

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(isActive, vm.Identifier)
	return err
}
func (vm *VirtualMachine) AttachVirtualMachineToFirewall(firewallId int64) error {
	var id int
	errQuery := database.DB.QueryRow(queries.QueryMap["AttachVirtualMachineToFirewall"], vm.Id, firewallId, uuid.New()).Scan(&id)
	if errQuery != nil {
		return errQuery
	}

	fmt.Println("Data inserted successfully!")

	return nil
}
