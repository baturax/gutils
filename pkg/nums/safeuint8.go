package nums

func SafeUint8(val int) uint8 {
	if val > 255 {
		return 255
	}
	if val < 0 {
		return 0
	}

	return uint8(val)
}
