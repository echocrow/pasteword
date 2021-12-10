package pasteword

// Read reads and formats a password.
func Read(name, placeholder string, vars ...string) (string, error) {
	tpl, err := getSecret(name)
	if err != nil {
		return "", err
	}
	// todo: format template.
	return tpl, nil
}

// Write writes a password template.
func Write(name string, tpl string) error {
	if tpl == "" {
		return delSecret(name)
	}
	return setSecret(name, tpl)
}
