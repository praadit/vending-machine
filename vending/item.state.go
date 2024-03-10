package vending

import (
	"fmt"
)

type DispensingItem struct {
	machine *VendingMachine
}

func SetDispensingItem(vm *VendingMachine) *DispensingItem {
	return &DispensingItem{
		machine: vm,
	}
}

func (state *DispensingItem) CollectCash() error {
	return fmt.Errorf("machine are dispensing item")
}
func (state *DispensingItem) DispenseItem() error {
	fmt.Println("Choose your product...")

	var productId string
	_, err := fmt.Scan(&productId)
	if err != nil {
		return err
	}

	price, ok := state.machine.Products[productId]
	if !ok {
		return fmt.Errorf("product not found")
	}

	if state.machine.Cash < price {
		return fmt.Errorf("insufficient cash amount")
	}

	state.machine.ChoosenPrice = price
	fmt.Printf("Vending cash left : %.2f\n", state.machine.Cash-state.machine.ChoosenPrice)
	return nil
}
func (state *DispensingItem) DispenseChange() (float64, error) {
	return 0.0, fmt.Errorf("dispensing Item")
}
func (state *DispensingItem) CancellingTransaction() error {
	return fmt.Errorf("cannot cancel transaction")
}
