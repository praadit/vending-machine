package main

import (
	"fmt"

	vending "github.com/praadit/vending-machine/states"
)

var productsList map[string]float64 = map[string]float64{
	"1": 10,
	"2": 15,
	"3": 5,
}

func main() {
	vm := vending.NewVendingMachine(productsList)

	var err error = nil
	fmt.Println("Welcome to vending machine, please choose one option below:")
	for err == nil {
		fmt.Println("1. Input cash")
		fmt.Println("2. Choose product")
		fmt.Println("3. Collect change")
		fmt.Println("4. Cancel transaction")
		fmt.Println("5. Exit")
		fmt.Println("What your option?")
		var option int
		if _, err := fmt.Scan(&option); err != nil {
			fmt.Println("exit due to error %s ", err.Error())
			break
		}

		var msg error
		switch option {
		case 1:
			msg = vm.CollectCash()
			if msg == nil {
				vm.State = vending.SetDispensingItem(vm)
			}
		case 2:
			msg = vm.DispenseItem()
			if msg == nil {
				vm.State = vending.SetDispensingChange(vm)
			}
		case 3:
			msg = vm.DispenseChange()
			if msg == nil {
				vm.State = vending.SetCancelTransaction(vm)
			}
		case 4:
			msg = vm.CancelTransaction()
			if msg == nil {
				vm.State = vending.SetReady(vm)
			}
		case 5:
			err = fmt.Errorf("exit by user")
		default:
			err = fmt.Errorf("exit due to error")
		}

		if msg != nil {
			fmt.Println(msg)
		}
		fmt.Printf("\n----------------------------------------\n\n")
	}
	if err != nil {
		fmt.Println(err)
	}
}
