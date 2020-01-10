package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"github.com/mmlt/tool-tmplt/expand"
)

var (
	// Version as set during build.
	Version string

	provider = flag.String("provider", "",
		`Provider is thycotic | azkv when secrets need to be fetched.`)
	url = flag.String("url", "",
		`Url:port of the secret store. For provider=azkv use https://name-of-keyvault.vault.azure.net`)
	username = flag.String("u", "",
		`For provider=thycotic; username of account to retrieve secret with.`)
	passw = flag.String("p", "",
		`For provider=thycotic; password of account to retrieve secret with.`)
	domain = flag.String("d", "",
		`For provider=thycotic; domain of account to retrieve secret with.`)
	tmplt = flag.String("t", "",
		`Filename of the template to expand.`)
	all = flag.String("a", "",
		`Filename of a yaml file that lists templates and the values to expand.`)
	setFile = flag.String("set-file", "",
		`Filename of a yaml file with values.`)
	usage = `Sxtmplt %s 
sxtmplt reads template files and expands {{ }} occurrences. Output goes to stdout.

Examples:
sxtmplt -t <templatefile> to expand a single template file.

sxtmplt -a <listfile> takes a yaml file with templates and values fields like this:
    templates:
    - file: hello.tpl
      values:
        team:
          lead: klukkluk
    values:
      who: world
      team:
        lead: pipo
Assuming 'hello.tpl' contains: {{ .Values.team.lead }} says hello {{ .Values.audience }}!
running sxtmplt -a test.yaml produces: klukkluk says hello world!


Functions:
Templating functions of 'https://golang.org/pkg/text/template/' and 'http://masterminds.github.io/sprig/' are included.
Other functions:
    thycotic - When provider=thycotic is selected occurrences like {{thycotic 1234 "Password"}} are replaced with 
    the corresponding Thycotic secret value (in this example 1234 represents the secret ID (an int32) and "Password"
    represents the field name of the secret value).
    Thycotic authentication uses -u and -p values.

    secret - When provider=azkv is selected occurrences like {{secret name-of-secret"}} are replaced with the
    corresponding value from https://name-of-keyvault.vault.azure.net/secrets/name-of-secret.
	Name-of-secret should match [0-9a-zA-Z\-]
    AZ KeyVault authentication uses -url value and AZURE_TENANT_ID, AZURE_CLIENT_ID, AZURE_CLIENT_SECRET environment variables.
    See https://docs.microsoft.com/en-us/azure/go/azure-sdk-go-authorization#use-environment-based-authentication

    filebase, filedir, fileclean, fileext
    Versions of base, dir, clean, ext that also work on Windows.

    {{ .Files.Get "filename" }}
	
    {{ range $name, $content := .Files.Glob "examples/*.yaml" }}
      {{ filebase $name }}: |
    {{ $content | indent 4 }}{{ end }}
	
    {{ (.Files.Glob "examples/*.yaml").AsConfig | indent 4 }}
	
    {{ (.Files.Glob "secrets/*").AsSecrets }}

Beware: file access is not sanitized.


Usage: sxtmplt [options...]
`
)

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, usage, Version)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	defer glog.Flush()
	if msg, ok := validate(); !ok {
		_, _ = fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	glog.V(2).Infof("provider=%s url=%s tmplt=%s all=%s set-file=%s", *provider, *url, *tmplt, *all, *setFile)
	err := expand.Run(*provider, *url, *username, *passw, *domain, *tmplt, *all, *setFile, expand.OSEnvironment(), os.Stdout)
	if err != nil {
		glog.Exit(err)
	}
}

// Validate checks prerequisite flags and environment variables.
// On validation failures it returns ok == false with a msg explaining the failures.
func validate() (msg string, ok bool) {
	if *tmplt == "" && *all == "" {
		return "-t or -a should be defined.", false
	}

	switch *provider {
	case "thycotic":
		if (*username == "") || (*passw == "") || (*domain == "") {
			return "provider=thycotic requires -u -p and -d to be set.", false
		}
		if *url == "" {
			return "provider=thycotic requires -url to be set.", false
		}
	case "azkv":
		env := expand.OSEnvironment()
		for _,v := range []string{"AZURE_TENANT_ID", "AZURE_CLIENT_ID", "AZURE_CLIENT_SECRET"} {
			if _, ok := env[v]; !ok{
				return fmt.Sprintf("provider=azkv requires environment variable %s to be set.", v), false
			}
		}
		if *url == "" {
			return "provider=thycotic requires -url to be set.", false
		}
	case "":
		if (*username != "") || (*passw != "") {
			return "Since v0.6.0 you need to set -provider=thycotic in combination with -u and -p.", false
		}
	default:
		return "-provider should be set to 'thycotic' or 'azkv' or not be set.", false
	}
	return "", true
}
