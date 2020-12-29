//args: -Ethelper
package testdata

import "testing"

func thelperWithHelperAfterAssignment(t *testing.T) { // ERROR "test helper function should start from t.Helper()"
	_ = 0
	t.Helper()
}

func thelperWithNotFirst(s string, t *testing.T, i int) { // ERROR "parameter \*testing.T should be the first"
	t.Helper()
}

func thelperWithIncorrectName(o *testing.T) { // ERROR "parameter \*testing.T should have name t"
	o.Helper()
}

func bhelperWithHelperAfterAssignment(b *testing.B) { // ERROR "test helper function should start from b.Helper()"
	_ = 0
	b.Helper()
}

func bhelperWithNotFirst(s string, b *testing.B, i int) { // ERROR "parameter \*testing.B should be the first"
	b.Helper()
}

func bhelperWithIncorrectName(o *testing.B) { // ERROR "parameter \*testing.B should have name b"
	o.Helper()
}
