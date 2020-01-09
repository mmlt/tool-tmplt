// test templates that use values from environment.
package expand_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/mmlt/tool-tmplt/expand"
)

// TestFunctionEnv yamlFileTests function 'env' (reading from environment)
func TestFunctionEnv(t *testing.T) {
	tf := testFilesNew()
	defer tf.MustRemoveAll()
	// create file(s) and environment
	tf.MustCreate("tpl/example.yaml", `
User is {{ env "TESTUSER" }}`)
	env := map[string]string{"TESTUSER": "Pipo"}
	// expand
	var out bytes.Buffer
	err := expand.Run("", "", "", "", tf.Path("tpl/example.yaml"), "", "", env, &out)
	assert.NoError(t, err)
	// assert
	assert.Equal(t, []byte(`
User is Pipo`), out.Bytes())
}
