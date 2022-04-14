package reverse_bits

// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	var reversed uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (num >> i) & 1
		//fmt.Print(bit)
		reversed += bit << (31 - i)
	}

	return reversed
}
