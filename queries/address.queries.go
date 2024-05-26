package queries

var QueryAddressMap = map[string]string{
	"insertAddress": " INSERT INTO addresses (city, street, country, identifier, user_id)  VALUES ($1, $2, $3, $4, $5) RETURNING identifier",
	"getAllAddress": "SELECT addresses.name,addresses.identifier,users.identifier As userIdentifier FROM addresses LEFT JOIN users ON users.id = addresses.user_id;",
	"getAddressById": "SELECT addresses.street,addresses.city,addresses.country,users.identifier As userIdentifier FROM addresses LEFT JOIN users ON users.id = addresses.user_id  WHERE addresses.identifier = $1",



}
