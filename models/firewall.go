package models
type Firewall struct{
	name string
	identifier string
}
func NewFirewall(name string) *Firewall{
	return &Firewall{
		name: name,
	}

}

func (firewall *Firewall)SetIdentifier(identifier string){
	firewall.identifier =identifier
}