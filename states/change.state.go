package vending

import (
	"fmt"
)

type DispensingChange struct {
	machine *VendingMachine
}

func SetDispensingChange(vm *VendingMachine) *DispensingChange {
	return &DispensingChange{
		machine: vm,
	}
}

func (state *DispensingChange) CollectCash() error {
	return fmt.Errorf("machine are dispensing change")
}
func (state *DispensingChange) DispenseItem() error {
	return fmt.Errorf("one item per transaction")
}
func (state *DispensingChange) DispenseChange() (float64, error) {
	fmt.Printf("Dispense change : %.2f\n", state.machine.Cash-state.machine.ChoosenPrice)
	return state.machine.Cash - state.machine.ChoosenPrice, nil
}
func (state *DispensingChange) CancellingTransaction() error {
	return fmt.Errorf("cannot cancel transaction")
}
