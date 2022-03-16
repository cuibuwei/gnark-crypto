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
	"math/rand"
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bw6-761/fr"
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

	// projective
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

	// extended
	properties.Property("Equal extended: having the receiver as operand should output the same result", prop.ForAll(
		func() bool {
			params := GetEdwardsCurve()
			var p1, baseProj PointProj
			p1.FromAffine(&params.Base)
			baseProj.FromAffine(&params.Base)

			return p1.Equal(&p1) && p1.Equal(&baseProj)
		},
	))

	properties.Property("Add extended: having the receiver as operand should output the same result", prop.ForAll(
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

	properties.Property("Double extended: having the receiver as operand should output the same result", prop.ForAll(
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

	properties.Property("Neg extended: having the receiver as operand should output the same result", prop.ForAll(
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

func TestField(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)
	genS := GenBigInt()

	properties.Property("MulByA(x) should match Mul(x, curve.A)", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var z1, z2 fr.Element
			z1.SetBigInt(&s)
			z2.Mul(&z1, &params.A)
			mulByA(&z1)

			return z1.Equal(&z2)
		},
		genS,
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))
}

func TestOps(t *testing.T) {

	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100

	properties := gopter.NewProperties(parameters)
	genS1 := GenBigInt()
	genS2 := GenBigInt()

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

			return p1.IsOnCurve() && p1.IsZero()
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

			return p1.IsOnCurve() && p1.Equal(&p2) && !p1.Equal(&inf)
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

			return p2.IsOnCurve() && p3.Equal(&p2) && !p3.Equal(&inf)
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

			return p2.IsOnCurve() && p2.Equal(&inf)
		},
		genS1,
	))

	properties.Property("(affine) [5]P=[2][2]P+P", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var p1, p2 PointAffine
			p1.ScalarMul(&params.Base, &s1)

			five := big.NewInt(5)
			p2.Double(&p1).Double(&p2).Add(&p2, &p1)
			p1.ScalarMul(&p1, five)

			return p2.IsOnCurve() && p2.Equal(&p1)
		},
		genS1,
	))

	// projective
	properties.Property("(projective) P+(-P)=O", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj, p1, p2, p PointProj
			baseProj.FromAffine(&params.Base)
			p1.ScalarMul(&baseProj, &s1)
			p2.Neg(&p1)

			p.Add(&p1, &p2)

			return p.IsZero()
		},
		genS1,
	))

	properties.Property("(projective) P+P=2*P", prop.ForAll(

		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj, p1, p2, p PointProj
			baseProj.FromAffine(&params.Base)
			p.ScalarMul(&baseProj, &s)

			p1.Add(&p, &p)
			p2.Double(&p)

			return p1.Equal(&p2)
		},
		genS1,
	))

	properties.Property("(projective) [5]P=[2][2]P+P", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj, p1, p2 PointProj
			baseProj.FromAffine(&params.Base)
			p1.ScalarMul(&baseProj, &s1)

			five := big.NewInt(5)
			p2.Double(&p1).Double(&p2).Add(&p2, &p1)
			p1.ScalarMul(&p1, five)

			return p2.Equal(&p1)
		},
		genS1,
	))

	// extended
	properties.Property("(extended) P+(-P)=O", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var baseExtended, p1, p2, p PointExtended
			baseExtended.FromAffine(&params.Base)
			p1.ScalarMul(&baseExtended, &s1)
			p2.Neg(&p1)

			p.Add(&p1, &p2)

			return p.IsZero()
		},
		genS1,
	))

	properties.Property("(extended) P+P=2*P", prop.ForAll(

		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseExtended, p1, p2, p PointExtended
			baseExtended.FromAffine(&params.Base)
			p.ScalarMul(&baseExtended, &s)

			p1.Add(&p, &p)
			p2.Double(&p)

			return p1.Equal(&p2)
		},
		genS1,
	))

	properties.Property("(extended) [5]P=[2][2]P+P", prop.ForAll(
		func(s1 big.Int) bool {

			params := GetEdwardsCurve()

			var baseExtended, p1, p2 PointExtended
			baseExtended.FromAffine(&params.Base)
			p1.ScalarMul(&baseExtended, &s1)

			five := big.NewInt(5)
			p2.Double(&p1).Double(&p2).Add(&p2, &p1)
			p1.ScalarMul(&p1, five)

			return p2.Equal(&p1)
		},
		genS1,
	))

	// mixed affine+extended
	properties.Property("(mixed affine+extended) P+(-P)=O", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseExtended, pExtended, p PointExtended
			var pAffine PointAffine
			baseExtended.FromAffine(&params.Base)
			pExtended.ScalarMul(&baseExtended, &s)
			pAffine.ScalarMul(&params.Base, &s)
			pAffine.Neg(&pAffine)

			p.MixedAdd(&pExtended, &pAffine)

			return p.IsZero()
		},
		genS1,
	))

	properties.Property("(mixed affine+extended) P+P=2*P", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseExtended, pExtended, p, p2 PointExtended
			var pAffine PointAffine
			baseExtended.FromAffine(&params.Base)
			pExtended.ScalarMul(&baseExtended, &s)
			pAffine.ScalarMul(&params.Base, &s)

			p.MixedAdd(&pExtended, &pAffine)
			p2.MixedDouble(&pExtended)

			return p.Equal(&p2)
		},
		genS1,
	))

	// mixed affine+projective
	properties.Property("(mixed affine+proj) P+(-P)=O", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj, pProj, p PointProj
			var pAffine PointAffine
			baseProj.FromAffine(&params.Base)
			pProj.ScalarMul(&baseProj, &s)
			pAffine.ScalarMul(&params.Base, &s)
			pAffine.Neg(&pAffine)

			p.MixedAdd(&pProj, &pAffine)

			return p.IsZero()
		},
		genS1,
	))

	properties.Property("(mixed affine+proj) P+P=2*P", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj, pProj, p, p2 PointProj
			var pAffine PointAffine
			baseProj.FromAffine(&params.Base)
			pProj.ScalarMul(&baseProj, &s)
			pAffine.ScalarMul(&params.Base, &s)

			p.MixedAdd(&pProj, &pAffine)
			p2.Double(&pProj)

			return p.Equal(&p2)
		},
		genS1,
	))

	properties.Property("scalar multiplication in Proj vs Ext should be consistant", prop.ForAll(
		func(s big.Int) bool {

			params := GetEdwardsCurve()

			var baseProj PointProj
			var baseExt PointExtended
			var p1, p2 PointAffine
			baseProj.FromAffine(&params.Base)
			baseProj.ScalarMul(&baseProj, &s)
			baseExt.FromAffine(&params.Base)
			baseExt.ScalarMul(&baseExt, &s)

			p1.FromProj(&baseProj)
			p2.FromExtended(&baseExt)

			return p1.Equal(&p2)
		},
		genS1,
	))

	properties.TestingRun(t, gopter.ConsoleReporter(false))

}

func TestMarshal(t *testing.T) {
	initOnce.Do(initCurveParams)

	var point, unmarshalPoint PointAffine
	point.Set(&curveParams.Base)
	for i := 0; i < 20; i++ {
		b := point.Marshal()
		unmarshalPoint.Unmarshal(b)
		if !point.Equal(&unmarshalPoint) {
			t.Fatal("error unmarshal(marshal(point))")
		}
		point.Add(&point, &curveParams.Base)
	}
}

// GenBigInt generates a big.Int
// TODO @thomas we use fr size as max bound here
func GenBigInt() gopter.Gen {
	return func(genParams *gopter.GenParameters) *gopter.GenResult {
		var s big.Int
		var b [fr.Bytes]byte
		_, err := rand.Read(b[:])
		if err != nil {
			panic(err)
		}
		s.SetBytes(b[:])
		genResult := gopter.NewGenResult(s, gopter.NoShrinker)
		return genResult
	}
}

// ------------------------------------------------------------
// benches

func BenchmarkScalarMulExtended(b *testing.B) {
	params := GetEdwardsCurve()
	var a PointExtended
	var s big.Int
	a.FromAffine(&params.Base)
	s.SetString("52435875175126190479447705081859658376581184513", 10)
	s.Add(&s, &params.Order)

	var doubleAndAdd PointExtended

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		doubleAndAdd.ScalarMul(&a, &s)
	}
}

func BenchmarkScalarMulProjective(b *testing.B) {
	params := GetEdwardsCurve()
	var a PointProj
	var s big.Int
	a.FromAffine(&params.Base)
	s.SetString("52435875175126190479447705081859658376581184513", 10)
	s.Add(&s, &params.Order)

	var doubleAndAdd PointProj

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		doubleAndAdd.ScalarMul(&a, &s)
	}
}
