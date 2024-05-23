package queries

var QueryFirewallMap = map[string]string{
	"insertFirewall": "INSERT INTO firewalls (name,user_id,identifier) VALUES ($1, $2,$3) RETURNING identifier ",
	"getAllFirewalls": "SELECT firewalls.name,firewalls.identifier,users.identifier As userIdentifier FROM firewalls LEFT JOIN users ON users.id = firewalls.user_id;",
	"getFirewallById": "SELECT firewalls.name,firewalls.identifier,users.identifier As userIdentifier FROM firewalls LEFT JOIN users ON users.id = firewalls.user_id  WHERE firewalls.identifier = $1",
	"deleteFirewall":  "DELETE FROM firewalls WHERE identifier = $1",
	"getFirewallByNameAndUserId":"SELECT firewalls.name FROM firewalls  WHERE firewalls.name = $1 AND firewalls.user_id=$2",

}
