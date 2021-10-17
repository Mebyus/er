package er

import (
	"fmt"
	"strings"
)

type Er = *er

type er struct {
	code  Code
	codes []Code
	texts []string
}

func New(code Code, text string) Er {
	return &er{
		code:  code,
		codes: []Code{code},
		texts: []string{text},
	}
}

func Newf(code Code, format string, v ...interface{}) Er {
	return New(code, fmt.Sprintf(format, v...))
}

func Newv(code Code, v ...interface{}) Er {
	return New(code, fmt.Sprint(v...))
}

func From(code Code, err error) Er {
	return New(code, err.Error())
}

func (e Er) Error() string {
	if e == nil {
		// prevent nil pointer dereference in situations where nil Er
		// was stored in interface
		return "<nil>"
	}

	// this is basically a copy of strings.Join but
	// it joins strings in reverse order and applies
	// predefined formatting
	switch len(e.codes) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("[#%08X](%s)", e.codes[0], e.texts[0])
	}
	n := 15*len(e.codes) - 2
	for i := 0; i < len(e.texts); i++ {
		n += len(e.texts[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprintf("[#%08X]", e.codes[len(e.codes)-1]))
	b.WriteByte('(')
	b.WriteString(e.texts[len(e.texts)-1])
	b.WriteByte(')')
	for i := len(e.texts) - 2; i >= 0; i-- {
		b.WriteString(fmt.Sprintf("<-[#%08X]", e.codes[i]))
		b.WriteByte('(')
		b.WriteString(e.texts[i])
		b.WriteByte(')')
	}
	return b.String()
}

func (e Er) Code() Code {
	return e.code
}

func (e Er) Up(code Code, text string) Er {
	e.codes = append(e.codes, code)
	e.texts = append(e.texts, text)
	return e
}

func (e Er) Upf(code Code, format string, v ...interface{}) Er {
	return e.Up(code, fmt.Sprintf(format, v...))
}

func (e Er) Upv(code Code, v ...interface{}) Er {
	return e.Up(code, fmt.Sprint(v...))
}
