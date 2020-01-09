Test .Files functions.
Invokation: thycotic -a test05.yaml

{ { .Files.Get "test.yaml" } }

{{ (.Files.Glob "*.yaml").AsConfig | indent 4 }}

{{ range $name, $content := .Files.Glob "*.yaml" }}
Config: {{ filebase $name }}
Content: {{ $content | indent 4 }}
{{ end }}
