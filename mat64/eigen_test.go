// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math"

	"gopkg.in/check.v1"
)

func (s *S) TestEigen(c *check.C) {
	for _, t := range []struct {
		a *Dense

		epsilon float64

		ev EigenValues
		v  *Dense
	}{
		{
			a: NewDense(3, 3, []float64{
				1, 2, 1,
				6, -1, 0,
				-1, -2, -1,
			}),

			epsilon: math.Pow(2, -52.0),

			ev: []complex128{3.0000000000000044, -4.000000000000003, -1.0980273383714707e-16},
			v: NewDense(3, 3, []float64{
				-0.48507125007266627, 0.41649656391752204, 0.11785113019775795,
				-0.7276068751089995, -0.8329931278350428, 0.7071067811865481,
				0.48507125007266627, -0.4164965639175216, -1.5320646925708532,
			}),
		},
		{
			a: NewDense(3, 3, []float64{
				1, 6, -1,
				6, -1, -2,
				-1, -2, -1,
			}),

			epsilon: math.Pow(2, -52.0),

			ev: []complex128{-6.240753470718579, -1.3995889142010132, 6.640342384919599},
			v: NewDense(3, 3, []float64{
				-0.6134279348516111, -0.31411097261113, -0.7245967607083111,
				0.7697297716508223, -0.03251534945303795, -0.6375412384185983,
				0.17669818159240022, -0.9488293044247931, 0.2617263908869383,
			}),
		},
		{ // Jama pvals
			a: NewDense(3, 3, []float64{
				4, 1, 1,
				1, 2, 3,
				1, 3, 6,
			}),

			ev: []complex128{0.34508918353562557, 3.5955906738074535, 8.059320142656922},

			epsilon: math.Pow(2, -52.0),
		},
		{ // Jama evals
			a: NewDense(4, 4, []float64{
				0, 1, 0, 0,
				1, 0, 2e-7, 0,
				0, -2e-7, 0, 1,
				0, 0, 1, 0,
			}),

			ev: []complex128{-0.9999999999999909 + 9.999999966061991e-08i, -0.9999999999999909 - 9.999999966061991e-08i, 0.9999999999999932 + 9.999999990731938e-08i, 0.9999999999999932 - 9.999999990731938e-08i},

			epsilon: math.Pow(2, -52.0),
		},
		{ // Jama badeigs
			a: NewDense(5, 5, []float64{
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 1,
				0, 0, 0, 1, 0,
				1, 1, 0, 0, 1,
				1, 0, 1, 0, 1,
			}),

			ev: []complex128{1.6180339887498956, 1.1102230246251565e-16, -0.6180339887498951, 0.9999999999999997i, -0.9999999999999997i},

			epsilon: math.Pow(2, -52.0),
		},
		{
			a: NewDense(3, 3, []float64{
				1, 6, 7,
				9, 1, 10,
				15, 6, 1,
			}),

			ev: []complex128{18.13437754618839, -7.567188773094204 + 1.0909493212308445i, -7.567188773094204 - 1.0909493212308445i},

			epsilon: math.Pow(2, -52.0),
		},
	} {
		ef := Eigen(DenseCopyOf(t.a), t.epsilon)
		if t.ev != nil {
			c.Check(ef.Values, check.DeepEquals, t.ev)
		}

		if t.v != nil {
			c.Check(ef.Vectors.Equals(t.v), check.Equals, true)
		}

		t.a.Mul(t.a, ef.Vectors)
		ef.Vectors.Mul(ef.Vectors, ef.Values.D())
		c.Check(t.a.EqualsApprox(ef.Vectors, 1e-12), check.Equals, true)
	}
}
