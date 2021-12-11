# pasteword

Retrieve passwords by temporarily putting them into the clipboard for easy copy-pasting.

On macOS, `pasteword` uses Keychain under the hood to store and retrieve passwords.


## Installation

Via [Homebrew](https://brew.sh/):
```sh
# Install:
brew install echocrow/tap/pasteword
# Update:
brew upgrade echocrow/tap/pasteword
```


## Examples

Set main password:
```sh
pasteword -set
# …followed by password prompt (also accepts stdin).
```
Retrieve & copy main pass:
```sh
pasteword
# …copies the previously stored main password into the clipboard for three seconds.
```
Set named password:
```sh
pasteword -set mypass
# …followed by password prompt (also accepts stdin).
```
Retrieve named password for ten seconds:
```sh
pasteword -ttl=10s mypass
# …copies "mypass" password into clipboard for ten seconds.
```
