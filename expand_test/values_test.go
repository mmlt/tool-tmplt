// test templates that use values from yaml file(s).
package expand_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/mmlt/tool-tmplt/expand"
)

var yamlFileTests = map[string]struct {
	tplpath, tpltext string
	tpath, ttext     string
	apath, atext     string
	vpath, vtext     string
	want             string
}{
	// Simple expands a template with a yaml provided value.
	"Simple": {
		tpath: "tpl/example.txt",
		ttext: `
User is {{ .Values.user }}`,
		vpath: "values.yaml",
		vtext: `
user: pipo`,
		want: `
User is pipo`,
	},

	// AllNoDefaults expands a template list with one template, no default values.
	"AllNoDefaults": {
		tplpath: "tpl/example.txt",
		tpltext: `
{{ .Values.team.lead }} says hello {{ .Values.audience }}!`,
		apath: "all.yaml",
		atext: `
templates:
- file: tpl/example.txt
  values:
    audience: all
    team:
      lead: pipo`,
		want: `
pipo says hello all!`,
	},

	// AllWithDefaults expands a template list with one template and default values.
	"AllWithDefaults": {
		tplpath: "tpl/example.txt",
		tpltext: `
{{ .Values.team.lead }} says hello {{ .Values.audience }}!`,
		apath: "all.yaml",
		atext: `
templates:
- file: tpl/example.txt
  values:
    team:
      lead: pipo
values:
  audience: all
  team:
    lead: klukkluk`,
		want: `
pipo says hello all!`,
	},

	// AllWithDefaultsAndValues expands a template list with one template, default values and cli set values.
	"AllWithDefaultsAndValues": {
		tplpath: "tpl/example.txt",
		tpltext: `
{{ .Values.team.lead }} says hello {{ .Values.audience }}!`,
		apath: "all.yaml",
		atext: `
templates:
- file: tpl/example.txt
  values:
    team:
      lead: pipo
values:
  audience: all
  team:
    lead: klukkluk`,
		vpath: "values.yaml",
		vtext: `
team:
  lead: mammaloe`,
		want: `
mammaloe says hello all!`,
	},
}

// TestYamlFiles.
func TestYamlFiles(t *testing.T) {
	for name, tst := range yamlFileTests {
		t.Run(name, func(t *testing.T) {
			tf := testFilesNew()
			defer tf.MustRemoveAll()
			// create file(s)
			tf.MustCreate(tst.tplpath, tst.tpltext)
			tf.MustCreate(tst.tpath, tst.ttext)
			tf.MustCreate(tst.apath, tst.atext)
			tf.MustCreate(tst.vpath, tst.vtext)
			// expand
			var out bytes.Buffer
			err := expand.Run("", "", "", tf.Path(tst.tpath), tf.Path(tst.apath), tf.Path(tst.vpath), nil, &out)
			assert.NoError(t, err)
			// assert
			assert.Equal(t, tst.want, out.String())
		})
	}
}
