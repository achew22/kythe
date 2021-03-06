// Package anon tests properties of anonymous types.
//
// Parameters and variables can be declared with anonymous types that are
// specified as part of their declaration. In the body of functions and of
// initializers, those fields may have references that we'd like to capture.
//
package anon

//- @planx defines/binding Planx
//- Planx.node/kind variable
func f(planx struct {
	//- T.node/kind variable
	//- T.subkind field
	//- @+3"T" defines/binding T

	// Count of thingies.
	T int

	// If they wrote comments, grab them, because heaven knows nobody is going
	// to understand this without 'em.
	//
	//- TDoc documents T
	//- TDoc.node/kind doc
	//- TDoc.text "Count of thingies."
}) int {
	//- @T ref T
	//- @planx ref Planx
	return planx.T
}

var v = struct {
	//- @V defines/binding V
	//- V.node/kind variable
	//- V.subkind field
	V int
}{
	//- @V ref V
	V: 25,
}

//- @V ref V
var _ = v.V

var w struct {
	//- @X defines/binding X
	//- X.node/kind variable
	//- X.subkind field
	X uint32
}

//- @X ref X
var _ = w.X

//- @elt defines/binding Elt
//- Elt.node/kind variable
var g = func(elt struct {
	//- @P defines/binding P
	//- P.node/kind variable
	//- P.subkind field
	P string
}) int {
	//- @P ref P
	//- @elt ref Elt
	return len(elt.P)
}

//- @em defines/binding Em
//- Em.node/kind record
type em struct {
	v struct {
		//- @X defines/binding EmX
		//- EmX.node/kind variable
		//- EmX.subkind field
		X string
	}
}
