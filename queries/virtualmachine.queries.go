package queries

var QueryMap = map[string]string{
	"getAllVirtualMachines": `
	SELECT virtualmachines.hostname,virtualmachines.is_active,virtualmachines.operating_system, virtualmachines.identifier,users.identifier FROM virtualmachines LEFT JOIN users ON users.id = virtualmachines.user_id where users.id =$1;
    `,
	"getVirtualMachineById":          "SELECT virtualmachines.id,virtualmachines.identifier As vm_identifier,is_active,operating_system,hostname,users.identifier As user_identifier,users.id As userId FROM virtualmachines LEFT JOIN users ON users.id = virtualmachines.user_id  WHERE virtualmachines.identifier = $1 AND  users.id =$2",
	"deleteVirtualMachine":           "DELETE FROM virtualmachines WHERE identifier = $1",
	"updateIsActive":                 "UPDATE virtualmachines SET is_active = $1 WHERE identifier = $2",
	"insertVirtualMachine":           "INSERT INTO virtualmachines (hostname, is_active,operating_system,user_id,identifier) VALUES ($1, $2,$3,$4,$5) RETURNING identifier",
	"AttachVirtualMachineToFirewall": "INSERT INTO virtualmachines_firewalls (virtualmachine_id, firewall_id,identifier) VALUES ($1, $2,$3) RETURNING id",
}
