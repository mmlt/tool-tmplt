# tmplt
tmplt is a tool to expand (interpolate) [GO templates](https://golang.org/pkg/text/template/).

For example given a template like `The secret {{thycotic 37027 "Password"}} is fetched from Thycotic.` 
running `tmplt -provider thycotic -url https://secret.example.com -u <user> -p <password> -t template.txt`
produces `The secret xxxx is fetched from Thycotic.` on stdout.

For more examples see `tmplt -h`

Compared to Helm this tool provides more functions;
- access environment
- access secrets stores's (currently Thycotic and AZ Key Vault)
- etc
(Note Since Helm v3 it's even more difficult to replace this tool with Helm as they made Render private.
 There is discussion around plugins for rendering however since more people have this requirement)

Compared to Kustomize this tool allows to expand the same template multiple times.



## Versions
0.1.0 the [sprig functions](http://masterminds.github.io/sprig) are included.
 
0.2.2 `-a` flag support that allows the same template to be deployed multiple times

0.3.0 `{{ .Files }}` support

0.4.0 `{{ toToml toYaml fromYaml toJson fromJson }}` support

0.5.0 `--set-file` support, clean-up, unit-tests.

0.6.0 `--provider` support for Azure Key Vault, tool rename to 'tmplt'.
    Note: this version is not backwards compatible.
    For interaction with Thycotic --provider=thycotic needs to be set explicitly (this used to be implicit when providing -u -p).


## Limitations/known issue's

### Formatting (1)
When using `{{ (.Files.Glob "*.yaml").AsConfig }}` the content is formatted nicely as long as it doesn't
contain `\r`, `\t` or `#` characters. 

Including such a characters results in content being formatted as a quoted string like this:
 `"package kubernetes.admission\n\n# deny if Pod spec (init)container(s)"`

### Formatting (2)
With `-a` templates are concatenated. But when a template doesn't end with a `\n` the next template is continued on the
same line. This is almost never what's expected. 
Current work-a-round is to be carefull to end templates with `\n`. 
It would be nice to have the tool automatically insert those `\n` however. 


## To do's
Add `--set key=value` 

Allow for more then one `--set` or `--set-file` on the cli

Add more unit-test; thycotic. 

Read `-t` (template) or `--set` from stdin. 
This would allows us to create pipelines 
`curl -s https://api.github.com/repos/golang/go/events | jp "[?type=='IssuesEvent'].payload. {Title: issue.title, URL: issue.url, User: issue.user.login, Event: action}" \
 | tmplt -t event-html.tpl -v -`


## Run
Run `tmplt -h` for help.

Run `tmplt -a test04.yaml` with
```
cat test04.tpl
{{ .Values.team.lead }} says hello {{ .Values.audience }}!
```
and 
```
cat test04.yaml
templates:
- test04.tpl
values:
  audience: world
  team:
    lead: pipo
```
produces: 'pipo says hello world!'


Run `thycotic -url https://secret.example.com -u <thycoticuser> -p <password> -t test.txt -v 5 -alsologtostderr true` with
```
cat test.txt
The password {{thycotic 123456 "Password"}} is fetched from Thycotic
```
produces: 'The password SuPeRsEcReT is fetched from Thycotic'


## Development
### Build
Install the [GO toolchain](https://golang.org/) and run `make build` to build Windows, Mac and Linux binaries.

#### Thycotic
Thycotic 9.1 and later support REST API's, from 10.1 docs are available. 
We currently (Nov 2017) run 9.0 and therefore have to use the SOAP interface.

To generate the Thycotic SOAP interface:
   1. Get gowsdl `go get github.com/hooklift/gowsdl/...`
   2. Get Thycotic wdsl `https://secret.example.com/SecretServer/webservices/sswebservice.asmx?wsdl >thycotic.wsdl`
   3. Generate code `gowsdl -o thycotic.go -p thycotic thycotic.wsdl`
   4. Fix code according to [README-generate-from-wsdl](./thycotic/README-generate-from-wsdl.md)

More info on Thycotic API:
- API information https://thycotic.force.com/support/s/developer-resources
- download for the .net API files: http://updates.thycotic.net/secretserver/configapi/appsettingsinterception.zip
- SOAP: https://updates.thycotic.net/secretserver/documents/SS_AppServerAPIGuide.pdf and https://secret.example.com/SecretServer/webservices/sswebservice.asmx


### Release

1. git commit changes
2. Run `VERSION=v0.x.y ARTIFACTORY_APIKEY=xxxx make release`
3. git push origin master --tags



