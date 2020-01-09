package expand

import (
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"github.com/mmlt/tool-tmplt/azkv"
	"github.com/mmlt/tool-tmplt/files"
	"github.com/mmlt/tool-tmplt/thycotic"
)

// Run expands one or more templates.
func Run(provider, url, username, passw, tmplt, all, setFile string, env map[string]string, out io.Writer) error {
	// get override values
	cliValues, err := readValuesFromYamlFile(setFile)
	if err != nil {
		return fmt.Errorf("reading %v: %v", setFile, err)
	}

	// remove sensitive data from OS environment
	env = sanitizeValue(env, passw)
	env = sanitizeKey(env, "AZURE_.*")

	// get template functions
	functions := getDefaultFunctions()
	// override sprig function to make sure a sanitized environment is used.
	functions["env"] = func(s string) string { return env[s] }
	functions["expandenv"] = func(s string) string { return "<expandenv is not supported>" }
	functions = addSecretFunction(functions, provider, url, username, passw)

	if tmplt != "" {
		// expand template
		err := expand(tmplt, functions, &Template{Values: cliValues, Files: files.Dir(filepath.Dir(tmplt))}, out)
		if err != nil {
			return fmt.Errorf("expanding: %v", err)
		}
	} else {
		bag, err := readConfigFromYamlFile(all)
		if err != nil {
			return fmt.Errorf("reading %v: %v", all, err)
		}

		err = expandAll(bag, functions, cliValues, filepath.Dir(all), out)
		if err != nil {
			return fmt.Errorf("expanding %s: %v", all, err)
		}
	}
	return nil
}

func expandAll(bag *Config, functions template.FuncMap, cliValues Values, basePath string, out io.Writer) error {
	for _, t := range bag.Templates {
		// get generic values
		v := deepCopy(bag.Values)
		// and merge template specific values
		merge(t.Values, v)
		// and merge cli provided values
		merge(cliValues, v)
		f := filepath.Join(basePath, t.File)
		err := expand(f, functions, &Template{Values: v, Files: files.Dir(filepath.Dir(f))}, out)
		if err != nil {
			return err
			//glog.Exitf("expanding %s: %v", *all, err)
		}
	}
	return nil
}

// Expand 'filename' template and write result to 'out'.
func expand(filename string, functions template.FuncMap, data interface{}, out io.Writer) error {
	// read template file
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	in := string(b)

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("input").Funcs(functions).Parse(in)
	if err != nil {
		glog.Exitf("parsing of '%s' failed: %s", filename, err)
	}
	// expand...
	return tmpl.Execute(out, data)
}

// AddSecretFunction
func addSecretFunction(functions template.FuncMap, provider, url, username, passw string) template.FuncMap {
	switch provider {
	case "thycotic":
		client, token, err := thycotic.Login(url, username, passw)
		if err != nil {
			glog.Exitf("thycotic: login failed: %v", err)
		}
		// add function to handle {{thycotic 1234 "Fieldnane"}} calls.
		functions["thycotic"] = func(id int, item string) string {
			s, err := thycotic.Get(int32(id), item, client, token)
			if err != nil {
				glog.Errorf("thycotic: %v", err)
			}
			return s
		}
	case "azkv":
		client, err := azkv.Login(url)
		if err != nil {
			glog.Exitf("azure key vault: login failed: %v", err)
		}
		// add function to handle {{secret "name-of-secret"}} calls.
		functions["secret"] = func(id string) string {
			s, err := azkv.Get(id, client, url)
			if err != nil {
				glog.Errorf("azure key vault: %v", err)
			}
			return s
		}
	}
	return functions
}

/*TODO // AddThycotic adds a function to retrieve secrets from thycotic.
func addThycotic(functions template.FuncMap, url, username, passw string) {
	client, token, err := thycoticLogin(url, username, passw)
	if err != nil {
		glog.Exitf("thycotic login failed: %v", err)
	}

	// add {{thycotic 1234 "Fieldnane"}} support.
	functions["thycotic"] = func(id int, item string) string { return getSecret(int32(id), item, client, token) }
}*/

func getDefaultFunctions() template.FuncMap {
	answer := sprig.TxtFuncMap()
	// add extra functionality
	answer["toToml"] = files.ToToml
	answer["toYaml"] = files.ToYaml
	answer["fromYaml"] = files.FromYaml
	answer["toJson"] = files.ToJson
	answer["fromJson"] = files.FromJson
	// add answer that sprig doesn't implement cross-platform (that don't work on windows)
	answer["filebase"] = filepath.Base
	answer["filedir"] = filepath.Dir
	answer["fileclean"] = filepath.Clean
	answer["fileext"] = filepath.Ext

	return answer
}

// ReadValuesFromYamlFile returns the contents of 'file' in 'Values' structure.
func readValuesFromYamlFile(file string) (Values, error) {
	answer := make(Values)
	if file != "" {
		// read file
		setYaml, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
			//glog.Exitf("reading: %v", err)
		}
		// unmarshal
		err = yaml.Unmarshal(setYaml, answer)
		if err != nil {
			return nil, fmt.Errorf("parsing: %v", err)
			//glog.Exitf("parsing %v: %v", *setFile, err)
		}
	}
	return answer, nil
}

// ReadConfigFromYamlFile returns the contents of 'file' in 'Config' structure.
func readConfigFromYamlFile(file string) (*Config, error) {
	answer := &Config{}
	// read file
	configYaml, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
		//glog.Exitf("reading: %v", err)
	}
	// unmarshal
	err = yaml.Unmarshal(configYaml, answer)
	if err != nil {
		return nil, fmt.Errorf("parsing: %v", err)
		//glog.Exitf("parsing %v: %v", *all, err)
	}
	return answer, nil
}

// Values is the map of template parameters.
type Values map[interface{}]interface{}

// Config is the file format of the yaml file used in combination with the -a flag.
// Values are 'global' values that are overridden by Template.Values.
type Config struct {
	Templates []TemplateConfig `yaml:"templates"`
	Values    Values           `yaml:"values"`
}

// TemplateConfig is the per-file entry as read from the Config.
type TemplateConfig struct {
	// File is the path to the template file.
	File string `yaml:"file"`
	// Values are the values that override global values.
	Values Values `yaml:"values"`
}

// Template contains the values and methods to use in template as in {{ .Values }} and {{ .Files }}
type Template struct {
	Values Values
	Files  files.Dir
}

// Merge merges src values into dst values.
func merge(src, dst Values) {
	for key, sv := range src {
		dv, found := dst[key]

		sm, sIsMap := sv.(Values)
		dm, dIsMap := dv.(Values)
		if found && sIsMap && dIsMap {
			merge(sm, dm)
		} else {
			dst[key] = sv
		}
	}
}

// DeepCopy Values.
func deepCopy(mp Values) Values {
	c := make(Values)
	for k, v := range mp {
		vm, ok := v.(Values)
		if ok {
			c[k] = deepCopy(vm)
		} else {
			c[k] = v
		}
	}

	return c
}

// OSEnvironment returns a map with OS environment variables.
func OSEnvironment() map[string]string {
	result := make(map[string]string)
	for _, s := range os.Environ() {
		sl := strings.SplitN(s, "=", 2)
		if len(sl) != 2 {
			continue
		}
		result[sl[0]] = sl[1]
	}
	return result
}

// SanitizeValue makes secret value in 'in' map unreadable.
func sanitizeValue(in map[string]string, secret string) map[string]string {
	for k, v := range in {
		if v == secret {
			in[k] = "********"
		}
	}
	return in
}

// Sanitize makes secret key in 'in' map unreadable.
func sanitizeKey(in map[string]string, secret string) map[string]string {
	r := regexp.MustCompile(secret)
	for k, _ := range in {
		if r.MatchString(k) {
			in[k] = "********"
		}
	}
	return in
}
