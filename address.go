package walletcore

type AddressType string

const (
	AddressTypeF1  AddressType = "f1"
	AddressTypeF4  AddressType = "f4"
	AddressTypeEth AddressType = "0x"
)

func (t AddressType) Valid() bool {
	switch t {
	case AddressTypeF1, AddressTypeF4, AddressTypeEth:
		return true
	}
	return false
}
