package frac_math

import "math/big"

// MyInt represents a number as a product of primes:
// n = p_1^e_1 * p_2^e_2 * ... * p_k^e_k
// where p_i are identified by integer IDs.
type MyInt struct {
	factors map[int]int // primeID -> exponent
}

// New creates an empty integer (value = 1)
func New() *MyInt {
	return &MyInt{
		factors: make(map[int]int),
	}
}

// FromFactors creates a MyInt from an existing map (copied)
func FromFactors(f map[int]int) *MyInt {
	m := New()
	for k, v := range f {
		if v != 0 {
			m.factors[k] = v
		}
	}
	return m
}
func (m *MyInt) Factors() map[int]int {
	// return a copy to prevent external mutation
	copy := make(map[int]int)
	for k, v := range m.factors {
		copy[k] = v
	}
	return copy
}

// Clone returns a deep copy
func (m *MyInt) Clone() *MyInt {
	return FromFactors(m.factors)
}

// Multiply: m *= n
func (m *MyInt) Multiply(n *MyInt) {
	for prime, exp := range n.factors {
		m.factors[prime] += exp
		if m.factors[prime] == 0 {
			delete(m.factors, prime)
		}
	}
}
func (m *MyInt) Power(exp int) {
	//fast exponentiation by squaring and fast exit on 0 and 1
	if exp == 0 {
		m.factors = make(map[int]int)
		return
	}
	if exp == 1 {
		return
	}
	if m.IsOne() {
		return
	}
	for prime, e := range m.factors {
		m.factors[prime] = e * exp
	}
}
func (m *MyInt) IsPowerOfTwo() bool {
	if m.IsOne() {
		return false
	}
	if len(m.factors) == 1 { // only one prime factor and it must be 2
		for prime := range m.factors {
			return prime == 2
		}
	}
	return false
}

// CanDivide checks if m is divisible by n
func (m *MyInt) CanDivide(n *MyInt) bool {
	for prime, exp := range n.factors {
		if m.factors[prime] < exp {
			return false
		}
	}
	return true
}

// Divide: m /= n, returns false if not divisible
func (m *MyInt) Divide(n *MyInt) bool {
	if !m.CanDivide(n) {
		return false
	}

	for prime, exp := range n.factors {
		m.factors[prime] -= exp
		if m.factors[prime] == 0 {
			delete(m.factors, prime)
		}
	}
	return true
}

// Equals checks structural equality
func (m *MyInt) Equals(n *MyInt) bool {
	if len(m.factors) != len(n.factors) {
		return false
	}
	for k, v := range m.factors {
		if n.factors[k] != v {
			return false
		}
	}
	return true
}

// IsOne checks if the value is 1
func (m *MyInt) IsOne() bool {
	return len(m.factors) == 0
}
func (m *MyInt) Set(other *MyInt) {
	m.factors = make(map[int]int)
	for k, v := range other.factors {
		m.factors[k] = v
	}
}

// String returns a human-readable representation of the number
func (m *MyInt) String() string {
	//create a big.Int and populate it by multiplying the prime factors
	result := big.NewInt(1)
	for prime, exp := range m.factors {
		primeBig := big.NewInt(int64(prime))
		primeBig.Exp(primeBig, big.NewInt(int64(exp)), nil)
		result.Mul(result, primeBig)
	}
	return result.String()

}

func formatFactor(p, e int) string {
	if e == 1 {
		return "p" + itoa(p)
	}
	return "p" + itoa(p) + "^" + itoa(e)
}

// minimal int to string (avoid fmt for speed if you care)
func itoa(x int) string {
	if x == 0 {
		return "0"
	}

	neg := false
	if x < 0 {
		neg = true
		x = -x
	}

	buf := [20]byte{}
	i := len(buf)

	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}

	if neg {
		i--
		buf[i] = '-'
	}

	return string(buf[i:])
}

//func (m *MyInt) fromBigInt(n *big.Int) {

func MyIntFromBigInt(n *big.Int) *MyInt {
	m := New()
	// factor n into primes and populate m.factors
	// This is a naive implementation and can be optimized with better factoring algorithms
	nCopy := new(big.Int).Set(n)
	for p := 2; nCopy.Cmp(big.NewInt(1)) > 0; p++ {
		count := 0
		for {
			mod := new(big.Int)
			mod.Mod(nCopy, big.NewInt(int64(p)))
			if mod.Cmp(big.NewInt(0)) != 0 {
				break
			}
			count++
			nCopy.Div(nCopy, big.NewInt(int64(p)))
		}
		if count > 0 {
			m.factors[p] = count
		}
	}
	return m
}

// string to bigint to myint
func MyIntFromString(s string) *MyInt {
	n := new(big.Int)
	n.SetString(s, 10)
	return MyIntFromBigInt(n)
}

func FromInt(n int) *MyInt {
	tab := make(map[int]int)
	for p := 2; n > 1; p++ {
		count := 0
		for n%p == 0 {
			count++
			n /= p
		}
		if count > 0 {
			tab[p] = count
		}
	}
	return FromFactors(tab)
}
