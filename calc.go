package tmc

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	calcLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Time`, `(\d+:)?\d+`},
		{`Ope`, `[-+]`},
		{`SP`, `\s+`},
	})

	parser = participle.MustBuild[Expr](
		participle.Lexer(calcLexer),
	)
)

type Val time.Duration

func (v *Val) Capture(values []string) error {
	s := values[0]

	if strings.Contains(s, ":") {
		hm := strings.SplitN(s, ":", 2)
		h, err := strconv.ParseInt(hm[0], 10, 64)

		if err != nil {
			return err
		}

		m, err := strconv.ParseInt(hm[1], 10, 64)

		if err != nil {
			return err
		}

		*v = Val(time.Duration(h)*time.Hour + time.Duration(m)*time.Minute)
	} else {
		m, err := strconv.ParseInt(s, 10, 64)

		if err != nil {
			return err
		}

		*v = Val(time.Duration(m) * time.Minute)
	}

	return nil
}

func (v Val) Duration() time.Duration {
	return time.Duration(v)
}

type OpeVal struct {
	Ope string `SP* @Ope`
	Val Val    `SP* @Time`
}

type Expr struct {
	Val     Val      `@Time`
	OpeVals []OpeVal `@@*`
}

func (expr *Expr) Eval() time.Duration {
	sum := expr.Val.Duration()

	for _, opeVal := range expr.OpeVals {
		switch opeVal.Ope {
		case "+":
			sum += opeVal.Val.Duration()
		case "-":
			sum -= opeVal.Val.Duration()
		default:
			panic("must not happen")
		}
	}

	return sum
}

func Eval(str string) (time.Duration, error) {
	str = strings.TrimSpace(str)
	expr, err := parser.ParseString("", str)

	if err != nil {
		return 0, err
	}

	return expr.Eval(), nil
}

func DurToStr(dur time.Duration) string {
	minus := false

	if dur < 0 {
		minus = true
		dur *= -1
	}

	mod := dur % time.Hour
	dur -= mod
	tmStr := fmt.Sprintf("%d:%02d", dur/time.Hour, mod/time.Minute)

	if minus {
		tmStr = "-" + tmStr
	}

	return tmStr
}
