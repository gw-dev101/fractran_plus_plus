package fracmath

import "testing"
func TestMultiplyCommutative(t *testing.T) {
	a := FromFactors(map[int]int{2: 1, 3: 2})
	b := FromFactors(map[int]int{3: 1, 5: 1})

	left := a.Clone()
	left.Multiply(b)

	right := b.Clone()
	right.Multiply(a)

	if !left.Equals(right) {
		t.Fatalf("expected commutativity: got %s and %s", left.String(), right.String())
	}
}

func TestMultiplyMutatesReciever(t *testing.T) {
	a := FromFactors(map[int]int{2: 1})
	b := FromFactors(map[int]int{3: 1})

	a.Multiply(b)

	expected := FromFactors(map[int]int{2: 1, 3: 1})
	if !a.Equals(expected) {
		t.Fatalf("expected mutation of receiver: got %s, expected %s", a.String(), expected.String())
	}
}
func TestMultiplyDoesntMutateArgument(t *testing.T) {
	a := FromFactors(map[int]int{2: 1})
	b := FromFactors(map[int]int{3: 1})
	bcopy := b.Clone()
	a.Multiply(b)

	if !b.Equals(bcopy) {
		t.Fatalf("expected no mutation of argument: got %s, expected %s", b.String(), bcopy.String())
	}}

func TestDivideCancel(t *testing.T) {
	a := FromFactors(map[int]int{2: 3, 3: 2})
	b := FromFactors(map[int]int{2: 1, 3: 1})
// (a multiply b) / b should equal a
c := a.Clone()
c.Multiply(b)
if !c.Divide(b) {
	t.Fatalf("expected division to succeed")
}
if !c.Equals(a) {
	t.Fatalf("expected (a * b) / b to equal a: got %s, expected %s", c.String(), a.String())
}
}

func TestDivideNotDivisible(t *testing.T) {
	a := FromFactors(map[int]int{2: 1})
	b := FromFactors(map[int]int{2: 2})

	if a.Divide(b) {
		t.Fatalf("expected division to fail")
	}
	if !a.Equals(FromFactors(map[int]int{2: 1})) {
		t.Fatalf("expected no mutation on failed division: got %s, expected %s", a.String(), FromFactors(map[int]int{2: 1}).String())
	}
}

func TestPower(t *testing.T) {
	a := FromFactors(map[int]int{2: 1, 3: 1})
	a.Power(3)

	expected := FromFactors(map[int]int{2: 3, 3: 3})
	if !a.Equals(expected) {
		t.Fatalf("expected power to work: got %s, expected %s", a.String(), expected.String())
	}
}

func TestPowerZero(t *testing.T) {
	a := FromFactors(map[int]int{2: 1, 3: 1})
	a.Power(0)

	expected := FromFactors(map[int]int{})
	if !a.Equals(expected) {
		t.Fatalf("expected power of zero to yield one: got %s, expected %s", a.String(), expected.String())
	}
}

func TestPowerOne(t *testing.T) {
	a := FromFactors(map[int]int{2: 1, 3: 1})
	acopy := a.Clone()
	a.Power(1)

	if !a.Equals(acopy) {
		t.Fatalf("expected power of one to yield same number: got %s, expected %s", a.String(), acopy.String())
	}
}

func TestPowerOneBase(t *testing.T) {
	a := FromFactors(map[int]int{})
	a.Power(5)

	expected := FromFactors(map[int]int{})
	if !a.Equals(expected) {
		t.Fatalf("expected power of one to yield one: got %s, expected %s", a.String(), expected.String())
	}
}