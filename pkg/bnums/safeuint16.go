package bnums

func SafeUint16(val int) uint16 {
	if val > 65535 {
		return 65535
	}

	if val < 0 {
		return 0
	}

	return uint16(val)
}
