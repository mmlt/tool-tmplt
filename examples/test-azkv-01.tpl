# Invokation
#   AZURE_TENANT_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
#   AZURE_CLIENT_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
#   AZURE_CLIENT_SECRET=xyz
# ./bin/tmplt -provider azkv -url https://xyz.vault.azure.net -t ./examples/test-azkv-01.tpl
The secret test-secret-001 has value {{secret "test-secret-001" }}