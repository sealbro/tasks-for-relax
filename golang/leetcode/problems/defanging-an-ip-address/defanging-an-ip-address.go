package defanging_an_ip_address

// https://leetcode.com/problems/defanging-an-ip-address/

func defangIPaddr(address string) string {
	buf := make([]byte, len(address)+6)

	j := 0

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			buf[j] = '['
			buf[j+1] = address[i]
			buf[j+2] = ']'
			j += 3
		} else {
			buf[j] = address[i]
			j++
		}
	}

	return string(buf)
}
