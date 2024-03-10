package vending

import (
	"fmt"
)

type Ready struct {
	machine *VendingMachine
}

func SetReady(vm *VendingMachine) *Ready {
	return &Ready{
		machine: vm,
	}
}

func (state *Ready) CollectCash() error {
	fmt.Println("Please input you cash...")

	var cash float64
	_, err := fmt.Scan(&cash)
	if err != nil {
		return err
	}

	state.machine.Cash = cash

	fmt.Printf("Vending cash left : %.2f\n", state.machine.Cash)
	return nil
}
func (state *Ready) DispenseItem() error {
	return fmt.Errorf("no cash collected")
}
func (state *Ready) DispenseChange() (float64, error) {
	return 0.0, fmt.Errorf("no cash collected")
}
func (state *Ready) CancellingTransaction() error {
	return fmt.Errorf("machine are not doing any transaction")
}
