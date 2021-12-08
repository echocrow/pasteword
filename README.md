# pasteword

> Paste basic or template-based keys and passwords.

Retrieve passwords and temporarily put them into the clipboard for easy copy-pasting.

On macOS, `pasteword` uses Keychain under the hood to store and retrieve keys and passwords.

`pasteword` optionally support reusable password templates with variable placeholders.

_**Note:** Reuse passwords or password templates at your own risk. Reusing passwords (entirely or partially) is not recommended._


## Installation

Via [Homebrew](https://brew.sh/):
```sh
# Install:
brew install echocrow/tap/pasteword
# Update:
brew upgrade echocrow/tap/pasteword
```


## Examples

### Main password
Set password:
```sh
pasteword -set
# …followed by password prompt (pr pipe in password via stdin).
```
Retrieve & copy pass:
```sh
pasteword
# …puts the previously stored password into the clipboard for three seconds.
```

### Named password with placeholders

Set named password template:
```sh
pasteword -k=mypass -set
# …followed by password prompt; enter e.g. "foo-{}.bar-{}.fizz-{}".
```
Retrieve & copy named pass with concrete values:
```sh
pasteword -k=mypass -ttl=10s ton zar ure
# …copies "foo-ton.bar-zar.fizz-ure"" into clipboard for ten seconds.
```
