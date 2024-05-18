package models

type VirtualMachine struct{
	hostName string
	ramSize int
	isActive bool
	identifier string
	operatingSystem string
	userId string
}
func NewVirtualMachine(hostName,operatingSystem,userId string, isActive bool, ramSize int) *VirtualMachine{
	return &VirtualMachine{
		hostName: hostName,
		ramSize: ramSize,
		isActive: isActive,
		operatingSystem: operatingSystem,
		userId:userId,
	}

}
func (vm *VirtualMachine)SetIdentifier(identifier string){
	vm.identifier =identifier
}