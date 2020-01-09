Test without thycotic credentials.
Invokation: thycotic-0.1.0-linux-x86_64 -t test02.tpl

This should be empty "{{thycotic 39262 "Password"}}"
Path is {{ env "PATH" }}
User is {{ env "USER" }}
Expandenv {{ expandenv "shows when thycotic isn't used $PWD" }}


