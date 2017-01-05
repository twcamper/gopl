package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expression string
		env        Env
		expected   string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"sin(x) + pow(y, 3)", Env{"x": 9, "y": 10}, "1000.41"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expression != prevExpr {
			fmt.Printf("\n%s\n", test.expression)
			prevExpr = test.expression
		}
		expression, err := Parse(test.expression)
		if err != nil {
			t.Error(err)
			continue
		}
		actual := fmt.Sprintf("%.6g", expression.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, actual)
		if actual != test.expected {
			t.Errorf("%s.Eval() in %v = %q, expected %q\n",
				test.expression, test.env, actual, test.expected)
		}
	}
}
