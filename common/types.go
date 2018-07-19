package common

const (
	HashLength    = 32
	AddressLength = 20
)

type Hash [HashLength]byte

// set b to Hash
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

// set the hash to the value of b
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

// get the byte representation of the underlying hash
func (h Hash) Bytes() []byte { return h[:] }

type Address [AddressLength]byte

// set byte representation of s to Address
func HexToAddress(s string) Address { return BytesToAddress(FromHex(s)) }

// set b to Address
func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

// set the address to the value of b
func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-HashLength:]
	}

	copy(a[AddressLength-len(b):], b)
}
