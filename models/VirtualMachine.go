package models
import(
	"CloudSystem/database"
	"fmt"
)

type VirtualMachine struct{
	hostName string
	isActive bool
	identifier string
	operatingSystem string
	userId string
}
func NewVirtualMachine(hostName,operatingSystem,userId string, isActive bool) *VirtualMachine{
	return &VirtualMachine{
		hostName: hostName,
		isActive: isActive,
		operatingSystem: operatingSystem,
		userId:userId,
	}

}
func (vm *VirtualMachine)SetIdentifier(identifier string) {
	vm.identifier =identifier
}
func (vm *VirtualMachine)InsertVirtualMachine() (int,error){

	sqlStatement := `INSERT INTO virtualmachines (hostname, is_active,operating_system) VALUES ($1, $2,$3) RETURNING id`
	// Prepare the SQL statement.
	var id int
	errQuery := database.DB.QueryRow(sqlStatement, vm.hostName, vm.isActive,vm.operatingSystem).Scan(&id)
	if errQuery != nil {
		fmt.Println(errQuery)
		return id, errQuery
	}

	fmt.Println("Data inserted successfully!")

	return id ,nil
} 