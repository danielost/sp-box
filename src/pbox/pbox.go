package pbox

import (
	"math/rand"
	"strconv"
)

// Permute the bits of an 8-bit unsigned integer.
// Returns a permuted 8-bit unsigned integer (byte) and a function to get the original value.
// Each call to the function randomly permutes the bits sequence
// and remembers the original position for each bit. Then the
// remembered positions are used to restore the original value.
func Permute(b uint8) (uint8, func(uint8) uint8) {
	binary := strconv.FormatInt(int64(b), 2)
	for len(binary) < 8 {
		binary = "0" + binary
	}
	bArr := []rune(binary)
	positions := make([]int, 8)
	for i := 0; i < len(positions); i++ {
		positions[i] = i
	}

	// Randomly shuffle the sequence using Fisher–Yates shuffle algorithm.
	for i := len(binary) - 1; i > 0; i-- {
		randomIndex := rand.Intn(i + 1)
		bArr[i], bArr[randomIndex] = bArr[randomIndex], bArr[i]

		// Remember the original positions of the bits.
		positions[i], positions[randomIndex] = positions[randomIndex], positions[i]
	}

	decimal, _ := strconv.ParseInt(string(bArr), 2, 64)

	// Return the permuted value and a reverse function.
	return uint8(decimal), getReverse(positions)
}

// Return a function that is used to restore the original value
// from a permuted sequence of bits. Uses remembered positions
// to restore the original value.
func getReverse(positions []int) func(uint8) uint8 {
	return func(b uint8) uint8 {
		binary := strconv.FormatInt(int64(b), 2)
		for len(binary) < 8 {
			binary = "0" + binary
		}
		bArr := []rune(binary)
		result := make([]rune, 8)
		for i := 0; i < len(positions); i++ {
			result[positions[i]] = bArr[i]
		}
		decimal, _ := strconv.ParseInt(string(result), 2, 64)
		return uint8(decimal)
	}
}
