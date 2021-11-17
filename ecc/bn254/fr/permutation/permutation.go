package permutation

import (
	"crypto/sha256"
	"errors"
	"math/big"
	"math/bits"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/fft"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/polynomial"
	fiatshamir "github.com/consensys/gnark-crypto/fiat-shamir"
)

var (
	ErrIncompatibleSize = errors.New("t1 and t2 should be of the same size")
	ErrSize             = errors.New("t1 and t2 should be of size a power of 2")
	ErrPermutationProof = errors.New("permutation proof verification failed")
)

// Proof proof that the commitments of t1 and t2 come from
// the same vector but permuted.
type Proof struct {

	// size of the polynomials
	size int

	// commitments of t1 & t2, the permuted vectors, and z, the accumulation
	// polynomial
	t1, t2, z kzg.Digest

	// commitment to the quotient polynomial
	q kzg.Digest

	// opening proofs of t1, t2, z, q (in that order)
	batchedProof kzg.BatchOpeningProof

	// shifted opening proof of z
	shiftedProof kzg.OpeningProof
}

// computeZ returns the accumulation polynomial in Lagrange basis.
func computeZ(lt1, lt2 []fr.Element, epsilon fr.Element) []fr.Element {

	s := len(lt1)
	z := make([]fr.Element, s)
	d := make([]fr.Element, s)
	z[0].SetOne()
	d[0].SetOne()
	nn := uint64(64 - bits.TrailingZeros64(uint64(s)))
	var t fr.Element
	for i := 0; i < s-1; i++ {
		_i := int(bits.Reverse64(uint64(i)) >> nn)
		_ii := int(bits.Reverse64(uint64((i+1)%s)) >> nn)
		z[_ii].Mul(&z[_i], t.Sub(&epsilon, &lt1[i]))
		d[i+1].Mul(&d[i], t.Sub(&epsilon, &lt2[i]))
	}
	d = fr.BatchInvert(d)
	for i := 0; i < s-1; i++ {
		_ii := int(bits.Reverse64(uint64((i+1)%s)) >> nn)
		z[_ii].Mul(&z[_ii], &d[i+1])
	}

	return z
}

// computeH computes lt2*z(gx) - lt1*z
func computeH(lt1, lt2, lz []fr.Element, epsilon fr.Element) []fr.Element {

	s := len(lt1)
	res := make([]fr.Element, s)
	var a, b fr.Element
	nn := uint64(64 - bits.TrailingZeros64(uint64(s)))
	for i := 0; i < s; i++ {
		_i := int(bits.Reverse64(uint64(i)) >> nn)
		_ii := int(bits.Reverse64(uint64((i+1)%s)) >> nn)
		a.Sub(&epsilon, &lt2[_i])
		a.Mul(&lz[_ii], &a)
		b.Sub(&epsilon, &lt1[_i])
		b.Mul(&lz[_i], &b)
		res[_i].Sub(&a, &b)
	}
	return res
}

// computeH0 computes L0 * (z-1)
func computeH0(lz []fr.Element, d *fft.Domain) []fr.Element {

	var tn, o, g fr.Element
	s := len(lz)
	tn.SetUint64(2).
		Neg(&tn)
	u := make([]fr.Element, s)
	o.SetOne()
	g.Set(&d.FinerGenerator)
	for i := 0; i < s; i++ {
		u[i].Sub(&g, &o)
		g.Mul(&g, &d.Generator)
	}
	u = fr.BatchInvert(u)
	res := make([]fr.Element, s)
	nn := uint64(64 - bits.TrailingZeros64(uint64(s)))
	for i := 0; i < s; i++ {
		_i := int(bits.Reverse64(uint64(i)) >> nn)
		res[_i].Sub(&lz[_i], &o).
			Mul(&res[_i], &u[i]).
			Mul(&res[_i], &tn)
	}
	return res
}

// Prove generates a proof that t1 and t2 are the same but permuted.
// The size of t1 and t2 should be the same and a power of 2.
func Prove(srs *kzg.SRS, t1, t2 []fr.Element) (Proof, error) {

	// res
	var proof Proof
	var err error

	// size checking
	if len(t1) != len(t2) {
		return proof, ErrIncompatibleSize
	}

	// create the domains
	d := fft.NewDomain(uint64(len(t1)), 1, false)
	if d.Cardinality != uint64(len(t1)) {
		return proof, ErrSize
	}
	s := int(d.Cardinality)
	proof.size = s

	// hash function for Fiat Shamir
	hFunc := sha256.New()

	// transcript to derive the challenge
	fs := fiatshamir.NewTranscript(hFunc, "epsilon", "omega", "eta")

	// commit t1, t2
	ct1 := make([]fr.Element, s)
	ct2 := make([]fr.Element, s)
	copy(ct1, t1)
	copy(ct2, t2)
	d.FFTInverse(ct1, fft.DIF, 0)
	d.FFTInverse(ct2, fft.DIF, 0)
	fft.BitReverse(ct1)
	fft.BitReverse(ct2)
	proof.t1, err = kzg.Commit(ct1, srs)
	if err != nil {
		return proof, err
	}
	proof.t2, err = kzg.Commit(ct2, srs)
	if err != nil {
		return proof, err
	}

	// derive challenge for z
	epsilon, err := deriveRandomness(&fs, "epsilon", &proof.t1, &proof.t2)
	if err != nil {
		return proof, err
	}

	// compute Z and commit it
	cz := computeZ(t1, t2, epsilon)
	d.FFTInverse(cz, fft.DIT, 0)
	proof.z, err = kzg.Commit(cz, srs)
	if err != nil {
		return proof, err
	}
	lz := make([]fr.Element, s)
	copy(lz, cz)
	d.FFT(lz, fft.DIF, 1)

	// compute the first part of the numerator
	lt1 := make([]fr.Element, s)
	lt2 := make([]fr.Element, s)
	copy(lt1, ct1)
	copy(lt2, ct2)
	d.FFT(lt1, fft.DIF, 1)
	d.FFT(lt2, fft.DIF, 1)
	h := computeH(lt1, lt2, lz, epsilon)

	// compute second part of the numerator
	h0 := computeH0(lz, d)

	// derive challenge used for the folding
	omega, err := deriveRandomness(&fs, "omega", &proof.z)
	if err != nil {
		return proof, err
	}

	// fold the numerator and divide it by x^n-1
	var t fr.Element
	t.SetUint64(2).Neg(&t).Inverse(&t)
	for i := 0; i < s; i++ {
		h0[i].Mul(&omega, &h0[i]).
			Add(&h0[i], &h[i]).
			Mul(&h0[i], &t)
	}

	// get the quotient and commit it
	d.FFTInverse(h0, fft.DIT, 1)
	proof.q, err = kzg.Commit(h0, srs)
	if err != nil {
		return proof, err
	}

	// derive the evaluation challenge
	eta, err := deriveRandomness(&fs, "eta", &proof.q)
	if err != nil {
		return proof, err
	}

	// compute the opening proofs
	proof.batchedProof, err = kzg.BatchOpenSinglePoint(
		[]polynomial.Polynomial{
			ct1,
			ct2,
			cz,
			h0,
		},
		[]kzg.Digest{
			proof.t1,
			proof.t2,
			proof.z,
			proof.q,
		},
		&eta,
		hFunc,
		d,
		srs,
	)
	if err != nil {
		return proof, err
	}

	eta.Mul(&eta, &d.Generator)
	proof.shiftedProof, err = kzg.Open(
		cz,
		&eta,
		d,
		srs,
	)
	if err != nil {
		return proof, err
	}

	// done
	return proof, nil

}

// TODO put that in fiat-shamir package
func deriveRandomness(fs *fiatshamir.Transcript, challenge string, points ...*bn254.G1Affine) (fr.Element, error) {

	var buf [bn254.SizeOfG1AffineUncompressed]byte
	var r fr.Element

	for _, p := range points {
		buf = p.RawBytes()
		if err := fs.Bind(challenge, buf[:]); err != nil {
			return r, err
		}
	}

	b, err := fs.ComputeChallenge(challenge)
	if err != nil {
		return r, err
	}
	r.SetBytes(b)
	return r, nil
}

// Verify verifies a permutation proof.
func Verify(srs *kzg.SRS, proof Proof) error {

	// hash function that is used for Fiat Shamir
	hFunc := sha256.New()

	// transcript to derive the challenge
	fs := fiatshamir.NewTranscript(hFunc, "epsilon", "omega", "eta")

	// derive the challenges
	epsilon, err := deriveRandomness(&fs, "epsilon", &proof.t1, &proof.t2)
	if err != nil {
		return err
	}

	omega, err := deriveRandomness(&fs, "omega", &proof.z)
	if err != nil {
		return err
	}

	eta, err := deriveRandomness(&fs, "eta", &proof.q)
	if err != nil {
		return err
	}

	// check the relation
	bs := big.NewInt(int64(proof.size))
	var l0, a, b, one, rhs, lhs fr.Element
	one.SetOne()
	rhs.Exp(eta, bs).
		Sub(&rhs, &one)
	a.Sub(&eta, &one)
	l0.Div(&rhs, &a)
	rhs.Mul(&rhs, &proof.batchedProof.ClaimedValues[3])
	a.Sub(&epsilon, &proof.batchedProof.ClaimedValues[1]).
		Mul(&a, &proof.shiftedProof.ClaimedValue)
	b.Sub(&epsilon, &proof.batchedProof.ClaimedValues[0]).
		Mul(&b, &proof.batchedProof.ClaimedValues[2])
	lhs.Sub(&a, &b)
	a.Sub(&proof.batchedProof.ClaimedValues[2], &one).
		Mul(&a, &l0).
		Mul(&a, &omega)
	lhs.Add(&a, &lhs)
	if !lhs.Equal(&rhs) {
		return ErrPermutationProof
	}

	// check the opening proofs
	err = kzg.BatchVerifySinglePoint(
		[]kzg.Digest{
			proof.t1,
			proof.t2,
			proof.z,
			proof.q,
		},
		&proof.batchedProof,
		hFunc,
		srs,
	)
	if err != nil {
		return err
	}

	err = kzg.Verify(&proof.z, &proof.shiftedProof, srs)
	if err != nil {
		return err
	}

	return nil
}
