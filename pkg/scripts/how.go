package scripts

import (
	"fmt"
)

// Create specific prompt for command line fixes
func GetScriptHowCmdQueryPrompt(text string, cmd string) string {
	return fmt.Sprintf("Give linux commands for problem I give you. Only output the command ready to be copy-paste, with no comment. \\nProblem: \\n%s. \\n\\nCommand, direclty pipeable, no quotes: \\n", cmd)
}
