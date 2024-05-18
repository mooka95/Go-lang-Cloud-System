package models
type Address struct{
	city string
	street string
	country string
	identifier string
}
func NewAddress(city,street,country string) *Address{
	return &Address{
		city: city,
		street: country,
		country: country,
	}

}

func (address *Address)SetIdentifier(identifier string){
	address.identifier =identifier
}