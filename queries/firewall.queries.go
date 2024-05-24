package queries

var QueryFirewallMap = map[string]string{
	"insertFirewall":                    "INSERT INTO firewalls (name,user_id,identifier) VALUES ($1, $2,$3) RETURNING identifier ",
	"getAllFirewalls":                   "SELECT firewalls.id,firewalls.name,firewalls.identifier,users.identifier As userIdentifier FROM firewalls LEFT JOIN users ON users.id = firewalls.user_id;",
	"getFirewallById":                   "SELECT firewalls.id,firewalls.name,firewalls.identifier,users.identifier As userIdentifier,users.id As userId FROM firewalls LEFT JOIN users ON users.id = firewalls.user_id  WHERE firewalls.identifier = $1",
	"deleteFirewall":                    "DELETE FROM firewalls WHERE identifier = $1",
	"getFirewallByNameAndUserId":        "SELECT firewalls.name FROM firewalls  WHERE firewalls.name = $1 AND firewalls.user_id=$2",
	"getFirewallAttachedVirtualMachine": "SELECT firewalls.name FROM firewalls INNER JOIN virtualmachines_firewalls ON virtualmachines_firewalls.firewall_id = firewalls.id WHERE firewalls.identifier = $1 AND virtualmachines_firewalls.virtualmachines_firewalls=$2 ",
}
