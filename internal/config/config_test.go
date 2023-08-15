package config_test

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"

	"github.com/artefactual-sdps/enduro/internal/config"
)

const testConfig = `# Config
debug = true
debugListen = "127.0.0.1:9001"

[temporal]
address = "host:port"

[upload]
secret = "don't tell"
`

func TestConfig(t *testing.T) {
	t.Setenv("ENDURO_UPLOAD_SECRET", "Sssh-Secret!")
	tmpDir := fs.NewDir(
		t, "",
		fs.WithFile(
			"enduro-config.toml",
			testConfig,
		),
	)
	configFile := tmpDir.Join("enduro-config.toml")

	var c config.Configuration
	found, configFileUsed, err := config.Read(&c, configFile)

	assert.NilError(t, err)
	assert.Equal(t, found, true)
	assert.Equal(t, configFileUsed, configFile)
	assert.Equal(t, c.Temporal.Address, "host:port")
	assert.Equal(t, c.Upload.Secret, "Sssh-Secret!")
}
