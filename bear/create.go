package bear

import "github.com/bjacobel/bear-cli/xcall"

// Create a new note in Bear and return the id
func Create(content string) {
	xcall.Call("open-note?id=0D27A2E0-4FD2-448B-99D1-905D9008203C-76039-00003D388806751F")
}
