Test with thycotic credentials.
Invokation: thycotic-0.1.0-linux-x86_64 -url https://secret.example.com -u xyz -p $PW -t test01.tpl

This is a password {{ thycotic 39262 "Password" }}
This is a password {{ thycotic 39262 "Password" | b64enc }}
User is {{ env "USER" }}
Password isn't visible {{ env "PW" }} (assuming -p $PW is used)
Expandenv {{ expandenv "is disabled when Thycotic is used" }}
