package models
type Ip struct{
	ipAddress string
	identifier string
	isReserved bool
}
func NewIp(ipAddress string, isReserved bool) *Ip{
	return &Ip{
		ipAddress: ipAddress,
		isReserved: isReserved,
	}

}

func (ip *Ip)SetIdentifier(identifier string){
	ip.identifier =identifier
}