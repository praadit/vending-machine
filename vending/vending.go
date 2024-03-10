package vending

type VendingMachine struct {
	State        MachineState
	Cash         float64
	Products     map[string]float64
	ChoosenPrice float64
}

func NewVendingMachine(products map[string]float64) *VendingMachine {
	vm := &VendingMachine{
		Products: products,
	}
	vm.State = SetReady(vm)

	return vm
}

func (vm *VendingMachine) CollectCash() error {
	return vm.State.CollectCash()
}
func (vm *VendingMachine) DispenseItem() error {
	return vm.State.DispenseItem()
}

func (vm *VendingMachine) DispenseChange() error {
	_, err := vm.State.DispenseChange()
	return err
}
func (vm *VendingMachine) CancelTransaction() error {
	return vm.State.CancellingTransaction()
}
