package scripts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetScriptFixCmdQueryPrompt(text string, cmd string, code string) string {
	return fmt.Sprintf("Fix given command with error. Answer shell command that can be piped. \\n\\nCommand: \"%s\" with error %s \\nFixed command, no leading #: ", cmd, code)
}

func GetScriptFixCmdDefaultTokens() int {
	return 512
}

func RunFix(text string) {
	fmt.Fprint(os.Stderr, strings.TrimLeft(text, "\n"))
	fmt.Fprintln(os.Stderr, " [^C to escape or Enter to run ]")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Println(strings.TrimLeft(text, "\n"))
}
