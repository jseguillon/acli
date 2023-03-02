package scripts

import (
	"fmt"
)

// Create specific prompt for command line fixes
// TODO: alos inject current $SHELL in promt for best answer
func GetScriptFixCmdQueryPrompt(text string, cmd string, code string) string {
	return fmt.Sprintf("Fix linux commands I give you. Only output the fixed command ready to be copy-paste, with no comment. \\nCommand with error: '%s' which gives error code %s. \\nFixed command, direclty pipeable, no quotes: ", cmd, code)
}
