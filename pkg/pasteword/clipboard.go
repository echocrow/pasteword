package pasteword

import (
	"time"

	"golang.design/x/clipboard"
)

// TempCopy copies string s into the clipboard and clears it after ttl.
func TempCopy(s string, ttl time.Duration) error {
	changed := clipboard.Write(clipboard.FmtText, []byte(s))
	select {
	case <-changed:
		// pass
	case <-time.After(ttl):
		clearClipboard()
	}
	return nil
}

func clearClipboard() {
	clipboard.Write(clipboard.FmtText, nil)
}
