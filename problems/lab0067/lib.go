package lab0067

// ascii码 0->48 1->49
func AddBinary(a string, b string) string {
	l := len(a)
	if l < len(b) {
		l = len(b)
	}

	res := make([]uint8, l)

	carry := false
	for i := 1; i <= l; i++ {
		var ia uint8 = '0'
		var ib uint8 = '0'

		if i-1 < len(a) {
			ia = a[len(a)-i]
		}
		if i-1 < len(b) {
			ib = b[len(b)-i]
		}

		if carry {
			carry = false

			if ia+ib >= 48+49 {
				carry = true
			}

			if ia+ib == 48+49 {
				res[l-i] = '0'
			} else {
				res[l-i] = '1'
			}
		} else {
			if ia+ib == 49+49 {
				carry = true
			}

			if ia+ib == 48+49 {
				res[l-i] = '1'
			} else {
				res[l-i] = '0'
			}
		}
	}

	if carry {
		res = append([]uint8{'1'}, res...)
	}

	return string(res)
}

func AddBinary2(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	if a[len(a)-1] == '1' && b[len(b)-1] == '1' {
		return AddBinary2(AddBinary2(a[:len(a)-1], b[:len(b)-1]), "1") + "0"
	} else if a[len(a)-1] == '0' && b[len(b)-1] == '0' {
		return AddBinary2(a[:len(a)-1], b[:len(b)-1]) + "0"
	} else {
		return AddBinary2(a[:len(a)-1], b[:len(b)-1]) + "1"
	}
}
