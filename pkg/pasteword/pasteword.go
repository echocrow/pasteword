package pasteword

// Read reads a password.
func Read(name string) (string, error) {
	return getSecret(name)
}

// Write writes password or deletes it when the new value is empty.
func Write(name string, value string) error {
	if value == "" {
		return delSecret(name)
	}
	return setSecret(name, value)
}
