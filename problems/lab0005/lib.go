package lab0005

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	max := 0
	index := 0

	start := 0

	p := make([]uint8, 0)
	ps := false //开始回文
	pl := 0
	pi := 0

	for i := 0; i < len(s); i++ {
		if len(p) == 0 {
			start = i
			p = append(p, s[i])
			pi = i
			pl = 0
			ps = false
			continue
		}

		if len(p) >= 2 {
			if s[i] == p[len(p)-2] {
				if i < len(s)-1 {
					if !ps {
						if s[i] == s[i+1] {
							p = append(p, s[i])
							continue
						} else {
							num := 0
							for j := i + 1; j < len(s) && len(p)-2+num+1 > 0; j++ {
								if s[j] != p[len(p)-2+num+1] {
									break
								}

								num++
							}

							if num > 0 {
								for j := 0; j < num; j++ {
									p = append(p, s[i+j+1])
								}
								i += num
								continue
							}
						}
					}
				}

				l := 0
				for j := len(p) - 3; j >= 0; j-- {
					if s[i] == p[j] {
						l += 1
					}
				}

				pl += 3 + l
				p = p[:len(p)-(2+l)]
				ps = true
			} else {
				if s[i] == p[len(p)-1] {
					if i < len(s)-1 {
						if s[i] == s[i+1] && !ps {
							p = append(p, s[i])
							continue
						}
					}

					pl += 2
					p = p[:len(p)-1]
					ps = true
				} else {
					if !ps {
						p = append(p, s[i])
					} else {
						//if p[len(p)-2] == p[len(p)-1] {
						//	pl += 2
						//	p = p[:len(p)-2]
						//	ps = true
						//	i -= 1
						//	continue
						//}

						p = make([]uint8, 0)
						i = start
						continue
					}
				}
			}
		} else {
			if s[i] == p[0] {
				if i < len(s)-1 {
					if s[i] == s[i+1] && !ps {
						p = append(p, s[i])
						continue
					}
				}

				pl += 2
				p = make([]uint8, 0)
			} else {
				if !ps {
					p = append(p, s[i])
				} else {
					p = make([]uint8, 0)
					i = start
					continue
				}
			}
		}

		if len(p) == 0 {
			if pl > max {
				max = pl
				index = pi
			}
			i = start
		}

		if i == len(s)-1 {
			if len(p) != 0 {
				p = make([]uint8, 0)
				i = start
			}
		}
	}

	if max == 0 {
		return s[:1]
	} else {
		return s[index : index+max]
	}
}
