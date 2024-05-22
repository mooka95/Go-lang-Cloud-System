package queries

var QueryFirewallMap = map[string]string{
	"insertFirewall": "INSERT INTO firewalls (name,user_id,identifier) VALUES ($1, $2,$3) RETURNING identifier ",
	"getAllFirewalls": "SELECT firewalls.name,firewalls.identifier,users.identifier As userIdentifier FROM firewalls LEFT JOIN users ON users.id = firewalls.user_id;",

}
