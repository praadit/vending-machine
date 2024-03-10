package vending

import (
	"fmt"
)

type CancelTransaction struct {
	machine *VendingMachine
}

func SetCancelTransaction(vm *VendingMachine) *CancelTransaction {
	return &CancelTransaction{
		machine: vm,
	}
}

func (state *CancelTransaction) CollectCash() error {
	return fmt.Errorf("machine are dispensing change")
}
func (state *CancelTransaction) DispenseItem() error {
	return fmt.Errorf("one item per transaction")
}
func (state *CancelTransaction) DispenseChange() (float64, error) {
	return 0.0, fmt.Errorf("no product selected")
}
func (state *CancelTransaction) CancellingTransaction() error {
	fmt.Println("Canceling transaction")
	return nil
}
