package fracmath

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
// This is ATOMIC: no mutation if division fails
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

// ApplyDelta applies a FRACTRAN-style delta:
// negative exponents = division requirement
// positive exponents = multiplication
func (m *MyInt) ApplyDelta(delta map[int]int) bool {
	// check feasibility (for negative parts)
	for prime, change := range delta {
		if change < 0 && m.factors[prime] < -change {
			return false
		}
	}

	// apply all changes
	for prime, change := range delta {
		m.factors[prime] += change
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

// String returns a debug representation
func (m *MyInt) String() string {
	if len(m.factors) == 0 {
		return "1"
	}

	s := ""
	for p, e := range m.factors {
		s += formatFactor(p, e) + " "
	}
	return s
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
