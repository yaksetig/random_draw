package cryptops

import "math/big"

// TODO: Change this to receive a min parameter
func RandInInterval(max *big.Int, seed []byte)(*big.Int, error){

	zero 	:= big.NewInt(0)
	one 	:= big.NewInt(1)

	// Check if max value is valid
	if max.Sign() <= 0 {
		panic("crypto/rand: argument to Int is <= 0")
	}

	// create new variable n containing the
	n := big.NewInt(0).Sub(max, big.NewInt(1))

	// bitLen is the maximum bit length needed to encode a value < max.
	maxBitLen := n.BitLen()

	// the only valid result is 0
	if maxBitLen == 0 {
		return zero, nil
	}

	// reqBytes is the maximum byte length needed to encode a value < max.
	reqBytes := (maxBitLen + 7) / 8

	// b is the number of bits in the most significant byte of max-1.
	b := uint(maxBitLen % 8)

	if b == 0 {
		b = 8
	}

	//fmt.Println("n bits:", b)

	// randMax is the max possible number that the PRF can generate within its byte output
	randMax := big.NewInt(0)
	randMax.Exp(big.NewInt(2), big.NewInt(int64(8*reqBytes)), nil)

	// rand_excess = (RAND_MAX % n) + 1;
	randExcess := big.NewInt(0)
	randExcess.Mod(randMax, max)
	randExcess.Add(randExcess, one)

	// randLimit is the max value that the RNG can generate.
	// randLimit = randMax - randExcess
	randLimit := big.NewInt(0)
	randLimit.Sub(randMax, randExcess)

	// Create the random value variable (type []byte)
	randomByteValue := make([]byte, reqBytes)

	// Create the random value variable (type big.Int)
	randomIntValue := big.NewInt(0)

	//
	s := make([]byte, 32)

	randomIntValue.SetBytes(randomByteValue)

	for {
		s = PRF(seed)
		// et the reqBytes that contain a random value from the PRF
		copy(randomByteValue, s[0:reqBytes])

		// Clear bits in the first byte to increase the probability that the candidate is < max.
		randomByteValue[0] &= uint8(int(1<<b) - 1)

		randomIntValue.SetBytes(randomByteValue)

		// if smaller than randLimit, return
		if randomIntValue.Cmp(randLimit) < 0  {
			return randomIntValue.Mod(randomIntValue, max), nil
		}
	}
}

// ShuffleList returns a (cryptographically seeded) Fisher-Yates shuffle of a list of integers
func ShuffleList(seed []byte, list []int)([]int){

	n := len(list)
	r := big.NewInt(0)
	data := seed

	for i:= n-1 ; i > 0 ; i-- {

		// Randomly generate a value in [0, i)
		r, _ = RandInInterval(big.NewInt(int64(i)), data)

		// convert value to int
		j := int(r.Int64())

		// Exchange list[j] and list [i]
		list[j], list[i] = list [i], list[j]

		// Hash the data value so that next iteration of the loop does not return the same value
		// TODO: Change this var update!
		data = PRF(data)
	}
	return list
}
