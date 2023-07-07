package main

import "github.com/quickfixgo/quickfix"

func main() {
	settings := quickfix.NewSettings()
	session_settings := quickfix.NewSessionSettings()
	settings.AddSession(session_settings)
}
