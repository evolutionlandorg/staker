/*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 2 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, see <http://www.gnu.org/licenses/>.
*
* Copyright (C) echo, 2021
 */
package apr

import (
	"fmt"
	"math"
	"math/big"
)

type Fraction struct {
	numerator   *big.Int
	denominator *big.Int
}

func NewFraction(numerator, denominator *big.Int) *Fraction {
	return &Fraction{numerator, denominator}
}

func (z *Fraction) Div(x, y *Fraction) *Fraction {
	z.numerator = new(big.Int).Mul(x.numerator, y.denominator)
	z.denominator = new(big.Int).Mul(x.denominator, y.numerator)
	return z
}

func (z *Fraction) Mul(x, y *Fraction) *Fraction {
	z.numerator = new(big.Int).Mul(x.numerator, y.numerator)
	z.denominator = new(big.Int).Mul(x.denominator, y.denominator)
	return z
}

func (x *Fraction) toFixed(decimal int) float64 {
	if x.numerator.Cmp(Big0) == 0 {
		return float64(0)
	}
	if x.denominator.Cmp(Big0) == 0 {
		return math.MaxFloat64
	}
	tmp := math.Pow10(decimal)
	res := new(big.Int).Div(new(big.Int).Mul(x.numerator, big.NewInt(int64(tmp))), x.denominator)
	return float64(res.Int64()) / tmp
}

func (x *Fraction) String() string {
	return fmt.Sprintf("{numerator:%v, denominator:%v}", x.numerator, x.denominator)
}
