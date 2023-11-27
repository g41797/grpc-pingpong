package server

import (
	"embed"
	"os"
	"testing"
)

//go:embed _configuration
var embconf embed.FS

func TestMain(m *testing.M) {
	ecode := testMain(m)
	os.Exit(ecode)
}

// second level function, because os.Exit does not honor defer
func testMain(m *testing.M) int {

	clean, err := PrepareTestEnvironment(&embconf)
	defer clean()

	if err != nil {
		return -1
	}

	ecode := m.Run()

	return ecode
}
