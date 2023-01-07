package scripts

import (
	"fmt"
)

// Create specific prompt for command line fixes
// TODO: alos inject current $SHELL in promt for best answer
func GetScriptFixCmdQueryPrompt(text string, cmd string, code string) string {
	return fmt.Sprintf("Fix given command with error. Answer shell command that can be piped. \\n\\nCommand: \"%s\" with error %s \\nFixed command, no leading #: ", cmd, code)
}

// Enforce max tokens for low cost query
func GetScriptFixCmdDefaultTokens() int {
	return 512
}
