package vending

type MachineState interface {
	CollectCash() error
	DispenseItem() error
	DispenseChange() (float64, error)
	CancellingTransaction() error
}
