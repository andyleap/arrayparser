package arrayparser

import (
	"fmt"
	"testing"
)

var examples = []string{
	"[a,b,c]",
	"[[a,b],[b,c]]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
	"[a,b, c ]]]]]]",
	"a",
	"◌́", // fruit unicode, acute "\u0301"
	"123]]]",
	"\n\n", // Newlines
	"asd",
	"]]][a,b, c ]]]]]]",
	//"",			// Crash the old solution.
}

func TestOddParser(t *testing.T) {

	for _, v := range examples {
		out, err := Parse(v)

		fmt.Println(out, err)
	}
}
