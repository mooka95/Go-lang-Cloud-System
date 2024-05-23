package queries

var QueryUserMap = map[string]string{
	"getUserByEmail": "SELECT identifier, password FROM users WHERE email = $1",
	"getUserIdFromIdentifier": "SELECT id FROM users WHERE identifier = $1",
}
