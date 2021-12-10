package pasteword

import (
	"errors"
	"strings"
)

var (
	ErrNoPlaceholder = errors.New("failed to format password: missing placeholder")
	ErrBadVarsLen    = errors.New("failed to format password: variables count mismatch")
)

// Read reads and formats a password.
func Read(name, placeholder string, vars ...string) (string, error) {
	tpl, err := getSecret(name)
	if err != nil {
		return "", err
	}
	if len(vars) == 0 {
		// Assume template-less pw or return template as-is.
		return tpl, nil
	}
	return fmtTpl(tpl, placeholder, vars...)
}

func fmtTpl(s, placeholder string, vars ...string) (string, error) {
	ph := placeholder

	if ph == "" {
		if len(vars) != 0 {
			return "", ErrNoPlaceholder
		} else {
			return s, nil
		}
	}

	n := strings.Count(s, ph)
	if n != len(vars) {
		return "", ErrBadVarsLen
	}

	varsLen := 0
	for _, v := range vars {
		varsLen += len(v)
	}

	// Apply replacements to buffer.
	var b strings.Builder
	b.Grow(len(s) - (n * len(ph)) + varsLen)
	start := 0
	for i := 0; i < n; i++ {
		j := start + strings.Index(s[start:], ph)
		b.WriteString(s[start:j])
		b.WriteString(vars[i])
		start = j + len(ph)
	}
	b.WriteString(s[start:])
	return b.String(), nil
}

// Write writes a password template.
func Write(name string, tpl string) error {
	if tpl == "" {
		return delSecret(name)
	}
	return setSecret(name, tpl)
}
