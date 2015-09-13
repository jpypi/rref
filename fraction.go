package main

import (
	"strconv"
)

type Fraction struct {
	numerator   int
	denominator int
}

func (self *Fraction) simplify() {

	if self.denominator != 0 && self.numerator%self.denominator == 0 {
		self.numerator /= self.denominator
		self.denominator = 1
	}
}

// Multiplication of fraction objects
func (self Fraction) mul(other Fraction) Fraction {
	return Fraction{
		numerator:   self.numerator * other.numerator,
		denominator: self.denominator * other.denominator,
	}
}

// Multiply a fraction by an integer
func (self Fraction) mulInt(integer int) Fraction {
	return Fraction{
		numerator:   self.numerator * integer,
		denominator: self.denominator * integer,
	}
}

// Adding two fractions together
func (self Fraction) add(other Fraction) Fraction {
	// TODO: could potentiallly use a math.Min(denominators) and % operator
	// to only need to multiply one fraction before adding (maybe faster?)
	new_self := self.mulInt(other.denominator)
	new_other := other.mulInt(self.denominator)

	return Fraction{
		numerator:   new_self.numerator + new_other.numerator,
		denominator: new_self.denominator,
	}
}

// Adding an integer to a fraction
func (self Fraction) addInt(integer int) Fraction {
	return Fraction{
		numerator:   self.numerator + self.denominator*integer,
		denominator: self.denominator,
	}
}

func (self Fraction) intVal() int {
	return self.numerator / self.denominator
}

func (self Fraction) floatVal() float32 {
	return float32(self.numerator) / float32(self.denominator)
}

// To implement the Stringy interface so it can be printed out
func (self Fraction) String() string {
	return strconv.FormatInt(int64(self.numerator), 10) + "/" +
		strconv.FormatInt(int64(self.denominator), 10)
}
