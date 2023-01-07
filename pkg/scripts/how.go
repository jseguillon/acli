package scripts

import (
	"fmt"
)

// Create specific prompt for command line fixes
func GetScriptHowCmdQueryPrompt(text string, cmd string) string {
	return fmt.Sprintf("Answer shell command or script that can be piped to solve given problem. \\n\\nProblem: \"%s\" \\nCommand, no leading #: ", cmd)
}

// Enforce max tokens for low cost query
func GetScriptHowCmdDefaultTokens() int {
	return 512
}
