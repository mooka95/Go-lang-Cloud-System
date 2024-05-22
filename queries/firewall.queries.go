package queries

var QueryFirewallMap = map[string]string{
	"insertFirewall": "INSERT INTO firewalls (name,user_id,identifier) VALUES ($1, $2,$3) RETURNING identifier ",

}
