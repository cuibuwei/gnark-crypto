// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package twistededwards

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bls12-381"
	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/prop"
)

// ------------------------------------------------------------
// tests

func TestReceiverIsOperand(t *testing.T) {

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)

	// affine
	properties.Property("Equal affine: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {
			params := GetEdwardsCurve()
			var p1 PointAffine
			p1.Set(&params.Base)

			return p1.Equal(&p1) && p1.Equal(&params.Base)
		},
	))

	properties.Property("Add affine: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2, p3 PointAffine
			p1.Set(&params.Base)
			p2.Set(&params.Base)
			p3.Set(&params.Base)

			res := true

			p3.Add(&p1, &p2)
			p1.Add(&p1, &p2)
			res = res && p3.Equal(&p1)

			p1.Set(&params.Base)
			p2.Add(&p1, &p2)
			res = res && p2.Equal(&p3)

			return res
		},
	))

	properties.Property("Double affine: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.Set(&params.Base)
			p2.Set(&params.Base)

			p2.Double(&p1)
			p1.Double(&p1)

			return p2.Equal(&p1)
		},
	))

	properties.Property("Neg affine: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.Set(&params.Base)
			p2.Set(&params.Base)

			p2.Neg(&p1)
			p1.Neg(&p1)

			return p2.Equal(&p1)
		},
	))

	properties.Property("Neg affine: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.Set(&params.Base)
			p2.Set(&params.Base)

			var s big.Int
			s.SetUint64(10)

			p2.ScalarMul(&p1, &s)
			p1.ScalarMul(&p1, &s)

			return p2.Equal(&p1)
		},
	))
	properties.TestingRun(t, gopter.ConsoleReporter(false))

	// proj
	properties.Property("Equal projective: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {
			params := GetEdwardsCurve()
			var p1, baseProj PointProj
			p1.FromAffine(&params.Base)
			baseProj.FromAffine(&params.Base)

			return p1.Equal(&p1) && p1.Equal(&baseProj)
		},
	))

	properties.Property("Add projective: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2, p3 PointProj
			p1.FromAffine(&params.Base)
			p2.FromAffine(&params.Base)
			p3.FromAffine(&params.Base)

			res := true

			p3.Add(&p1, &p2)
			p1.Add(&p1, &p2)
			res = res && p3.Equal(&p1)

			p1.FromAffine(&params.Base)
			p2.Add(&p1, &p2)
			res = res && p2.Equal(&p3)

			return res
		},
	))

	properties.Property("Double projective: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2 PointProj
			p1.FromAffine(&params.Base)
			p2.FromAffine(&params.Base)

			p2.Double(&p1)
			p1.Double(&p1)

			return p2.Equal(&p1)
		},
	))

	properties.Property("Neg projective: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {

			params := GetEdwardsCurve()

			var p1, p2 PointProj
			p1.FromAffine(&params.Base)
			p2.FromAffine(&params.Base)

			p2.Neg(&p1)
			p1.Neg(&p1)

			return p2.Equal(&p1)
		},
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))

}

func TestOps(t *testing.T) {

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)
	genS1 := bls12381.GenBigInt()
	genS2 := bls12381.GenBigInt()

	// affine
	properties.Property("(affine) P+(-P)=O", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.ScalarMul(&params.Base, &s1)
			p2.Neg(&p1)

			p1.Add(&p1, &p2)

			var one fr.Element
			one.SetOne()

			return p1.X.IsZero() && p1.Y.Equal(&one)
		},
		genS1,
	))

	properties.Property("(affine) P+P=2*P", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2, inf PointAffine
			p1.ScalarMul(&params.Base, &s)
			p2.ScalarMul(&params.Base, &s)

			p1.Add(&p1, &p2)
			p2.Double(&p2)

			return p1.Equal(&p2) && !p1.Equal(&inf)
		},
		genS1,
	))

	properties.Property("(affine) [a]P+[b]P = [a+b]P", prop.ForAll(
		func(s1, s2 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2, p3, inf PointAffine
			inf.X.SetZero()
			inf.Y.SetZero()
			p1.ScalarMul(&params.Base, &s1)
			p2.ScalarMul(&params.Base, &s2)
			p3.Set(&params.Base)

			p2.Add(&p1, &p2)

			s1.Add(&s1, &s2)
			p3.ScalarMul(&params.Base, &s1)

			return p3.Equal(&p2) && !p3.Equal(&inf)
		},
		genS1,
		genS2,
	))

	properties.Property("(affine) [a]P+[-a]P = O", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2, inf PointAffine
			inf.X.SetZero()
			inf.Y.SetOne()
			p1.ScalarMul(&params.Base, &s1)
			s1.Neg(&s1)
			p2.ScalarMul(&params.Base, &s1)

			p2.Add(&p1, &p2)

			return p2.Equal(&inf)
		},
		genS1,
	))

	properties.Property("[5]P=[2][2]P+P", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.ScalarMul(&params.Base, &s1)

			five := big.NewInt(5)
			p2.Double(&p1).Double(&p2).Add(&p2, &p1)
			p1.ScalarMul(&p1, five)

			return p2.Equal(&p1)
		},
		genS1,
	))

	// proj
	properties.Property("(projective) P+(-P)=O", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.ScalarMul(&params.Base, &s1)
			p2.Neg(&p1)

			var _p1, _p2 PointProj
			_p1.FromAffine(&p1)
			_p2.FromAffine(&p2)
			_p1.Add(&_p1, &_p2)

			var p PointAffine
			p.FromProj(&_p1)

			var one fr.Element
			one.SetOne()

			return p.X.IsZero() && p.Y.Equal(&one)
		},
		genS1,
	))

	properties.Property("(projective) P+P=2*P", prop.ForAll(

		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2, inf PointAffine
			p1.ScalarMul(&params.Base, &s)
			p2.ScalarMul(&params.Base, &s)

			var _p1, _p2 PointProj
			_p1.FromAffine(&p1)
			_p2.FromAffine(&p2)
			_p1.Add(&_p1, &_p2)
			_p2.Double(&_p2)

			p1.FromProj(&_p1)
			p2.FromProj(&_p2)

			return p1.Equal(&p2) && !p1.Equal(&inf)
		},
		genS1,
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))

}

func TestMarshal(t *testing.T) {

	var point, unmarshalPoint PointAffine
	point.Set(&edwards.Base)
	for i := 0; i < 20; i++ {
		b := point.Marshal()
		unmarshalPoint.Unmarshal(b)
		if !point.Equal(&unmarshalPoint) {
			t.Fatal("error unmarshal(marshal(point))")
		}
		point.Add(&point, &edwards.Base)
	}
}
