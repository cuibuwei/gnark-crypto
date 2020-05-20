package fp2

const Base = `
import (
	"github.com/consensys/gurvy/{{.Fpackage}}/fp"
)

// {{.Fp2Name}} is a degree-two finite field extension of fp.Element:
// A0 + A1u where u^2 == {{.Fp2NonResidue}} is a quadratic nonresidue in fp

type {{.Fp2Name}} struct {
	A0, A1 fp.Element
}

// Equal returns true if z equals x, fasle otherwise
// TODO can this be deleted?  Should be able to use == operator instead
func (z *{{.Fp2Name}}) Equal(x *{{.Fp2Name}}) bool {
	return z.A0.Equal(&x.A0) && z.A1.Equal(&x.A1)
}

// SetString sets a {{.Fp2Name}} element from strings
func (z *{{.Fp2Name}}) SetString(s1, s2 string) *{{.Fp2Name}} {
	z.A0.SetString(s1)
	z.A1.SetString(s2)
	return z
}

func (z *{{.Fp2Name}}) SetZero() *{{.Fp2Name}} {
	z.A0.SetZero()
	z.A1.SetZero()
	return z
}

// Clone returns a copy of self
func (z *{{.Fp2Name}}) Clone() *{{.Fp2Name}} {
	return &{{.Fp2Name}}{
		A0: z.A0,
		A1: z.A1,
	}
}

// Set sets an {{.Fp2Name}} from x
func (z *{{.Fp2Name}}) Set(x *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.Set(&x.A0)
	z.A1.Set(&x.A1)
	return z
}

// Set sets z to 1
func (z *{{.Fp2Name}}) SetOne() *{{.Fp2Name}} {
	z.A0.SetOne()
	z.A1.SetZero()
	return z
}

// SetRandom sets a0 and a1 to random values
func (z *{{.Fp2Name}}) SetRandom() *{{.Fp2Name}} {
	z.A0.SetRandom()
	z.A1.SetRandom()
	return z
}

// Equal returns true if the two elements are equal, fasle otherwise
func (z *{{.Fp2Name}}) IsZero() bool {
	return z.A0.IsZero() && z.A1.IsZero()
}

// Neg negates an {{.Fp2Name}} element
func (z *{{.Fp2Name}}) Neg(x *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.Neg(&x.A0)
	z.A1.Neg(&x.A1)
	return z
}

// String implements Stringer interface for fancy printing
func (z *{{.Fp2Name}}) String() string {
	return (z.A0.String() + "+" + z.A1.String() + "*u")
}

// ToMont converts to mont form
func (z *{{.Fp2Name}}) ToMont() *{{.Fp2Name}} {
	z.A0.ToMont()
	z.A1.ToMont()
	return z
}

// FromMont converts from mont form
func (z *{{.Fp2Name}}) FromMont() *{{.Fp2Name}} {
	z.A0.FromMont()
	z.A1.FromMont()
	return z
}

// Add adds two elements of {{.Fp2Name}}
func (z *{{.Fp2Name}}) Add(x, y *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.Add(&x.A0, &y.A0)
	z.A1.Add(&x.A1, &y.A1)
	return z
}

// AddAssign adds x to z
func (z *{{.Fp2Name}}) AddAssign(x *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.AddAssign(&x.A0)
	z.A1.AddAssign(&x.A1)
	return z
}

// Sub two elements of {{.Fp2Name}}
func (z *{{.Fp2Name}}) Sub(x, y *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.Sub(&x.A0, &y.A0)
	z.A1.Sub(&x.A1, &y.A1)
	return z
}

// SubAssign subs x from z
func (z *{{.Fp2Name}}) SubAssign(x *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.SubAssign(&x.A0)
	z.A1.SubAssign(&x.A1)
	return z
}

// Double doubles an {{.Fp2Name}} element
func (z *{{.Fp2Name}}) Double(x *{{.Fp2Name}}) *{{.Fp2Name}} {
	z.A0.Double(&x.A0)
	z.A1.Double(&x.A1)
	return z
}
`