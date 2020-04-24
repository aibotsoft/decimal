package decimal

func (d *Decimal) Size() int {
	if d == nil {
		return 0
	}
	return 1
}
