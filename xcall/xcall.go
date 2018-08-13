package xcall

import (
	"encoding/json"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/derekstavis/go-qs"
)

// Response contains the data Bear returns for each x-success response to a x-callback-url call
type Response struct {
	Note             string
	ModificationDate time.Time
	CreationDate     time.Time
	Title            string
	IsTrashed        bool
	Identifier       string
}

// Call the obj-C xcall code to open the passed URI as an x-callback-url
func Call(endpoint string, params map[string]interface{}) *Response {
	qstr, _ := qs.Marshal(params)
	uri := strings.Join([]string{endpoint, qstr}, "?")

	xcallPath, _ := filepath.Abs("xcall/xcall.app/Contents/MacOS/xcall")

	cmd := exec.Command(
		xcallPath,
		"-url",
		strings.Join([]string{"bear://x-callback-url", uri}, "/"),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	var resp Response

	if err := json.Unmarshal(output, &resp); err != nil {
		log.Fatal(err)
	}

	return &resp
}
