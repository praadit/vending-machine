package vending

type VendingMachine struct {
	state        MachineState
	cash         float64
	products     map[string]float64
	choosenPrice float64
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{}

	return vm
}

func (vm *VendingMachine) CollectCash() error {
	return vm.state.CollectCash()
}
func (vm *VendingMachine) DispenseItem() error {
	return vm.state.DispenseItem()
}

func (vm *VendingMachine) DispenseChange() error {
	_, err := vm.state.DispenseChange()
	return err
}
func (vm *VendingMachine) CancelTransaction() error {
	return vm.state.CancellingTransaction()
}

func (vm *VendingMachine) GetCash() float64                { return vm.cash }
func (vm *VendingMachine) GetState() MachineState          { return vm.state }
func (vm *VendingMachine) GetProducts() map[string]float64 { return vm.products }
func (vm *VendingMachine) GetChoosenPrice() float64        { return vm.choosenPrice }

func (vm *VendingMachine) SetCash(input float64)                { vm.cash = input }
func (vm *VendingMachine) SetState(input MachineState)          { vm.state = input }
func (vm *VendingMachine) SetProducts(input map[string]float64) { vm.products = input }
func (vm *VendingMachine) SetChoosenPrice(input float64)        { vm.choosenPrice = input }
