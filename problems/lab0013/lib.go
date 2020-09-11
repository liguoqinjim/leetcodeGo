package lab0013

//Symbol       Value	rune
//I             1		73
//V             5		86
//X             10		88
//L             50		76
//C             100		67
//D             500		68
//M             1000	77

func RomanToInt(s string) int {
	num := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i < len(s)-1 {
				vv := s[i+1]
				switch vv {
				case 'V':
					num += 4
					i++
				case 'X':
					num += 9
					i++
				default:
					num++
				}
			} else {
				num += 1
			}
		case 'V':
			num += 5
		case 'X':
			if i < len(s)-1 {
				vv := s[i+1]
				switch vv {
				case 'L':
					num += 40
					i++
				case 'C':
					num += 90
					i++
				default:
					num += 10
				}
			} else {
				num += 10
			}
		case 'L':
			num += 50
		case 'C':
			if i < len(s)-1 {
				vv := s[i+1]
				switch vv {
				case 'D':
					num += 400
					i++
				case 'M':
					num += 900
					i++
				default:
					num += 100
				}
			} else {
				num += 100
			}
		case 'D':
			num += 500
		case 'M':
			num += 1000
		}
	}

	return num
}

func RomanToInt2(s string) int {
	num := 0

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case 'I':
			if num >= 5 {
				num -= 1
			} else {
				num += 1
			}
		case 'V':
			num += 5
		case 'X':
			if num >= 50 {
				num -= 10
			} else {
				num += 10
			}
		case 'L':
			num += 50
		case 'C':
			if num >= 500 {
				num -= 100
			} else {
				num += 100
			}
		case 'D':
			num += 500
		case 'M':
			num += 1000
		}
	}

	return num
}
