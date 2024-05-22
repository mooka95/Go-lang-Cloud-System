package models
import (
	"CloudSystem/database"
	"CloudSystem/queries"
	"fmt"

	"github.com/google/uuid"
)
type Firewall struct{
	Name string `json:"name" binding:"required"`
	Identifier string 
	UserIdentifier string
}
func NewFirewall(name string) *Firewall{
	return &Firewall{
		Name: name,
	}

}

func (firewall *Firewall) InsertFirewall() (*string, error) {

	// Prepare the SQL statement.
	var id string
	errQuery := database.DB.QueryRow(queries.QueryFirewallMap["insertFirewall"],firewall.Name,3, uuid.New()).Scan(&id)
		if errQuery != nil {
		fmt.Println(errQuery)
		return nil, errQuery
	}
	return &id, nil
}