package queries

var QueryMap = map[string]string{
	"getAllVirtualMachines": `
	SELECT virtualmachines.hostname,virtualmachines.is_active,virtualmachines.operating_system, virtualmachines.identifier,users.identifier FROM virtualmachines LEFT JOIN users ON users.id = virtualmachines.user_id;
    `,
	"getVirtualMachineById": "SELECT virtualmachines.identifier As vm_identifier,is_active,operating_system,hostname,users.identifier As user_identifier FROM virtualmachines LEFT JOIN users ON users.id = virtualmachines.user_id  WHERE virtualmachines.identifier = $1 ",
	"deleteVirtualMachine":  "DELETE FROM virtualmachines WHERE identifier = $1",
	"updateIsActive":"UPDATE virtualmachines SET is_active = $1 WHERE identifier = $2",
}
