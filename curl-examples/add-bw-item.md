```bash
curl -Ls -X POST http://localhost:7787/object/item \
  --header 'Content-Type: application/json' \
  --data "{
            \"folderId\": null,
            \"type\": 1,
            \"name\": \"test-bwca-add\",
            \"notes\": null,
            \"favorite\": false,
            \"fields\": [
              {
                \"Security Question\": \"Bitwarden Rules\"
              }
            ],
            \"login\": {
              \"uris\": [
                {
                  \"match\": 0,
                  \"uri\": \"https://twitter.com/login\"
                }
              ],
              \"username\": \"twitter@acmecorp.com\",
              \"password\": \"..........dog..........\",
              \"totp\": null
            },
            \"reprompt\": true
    }"
```

```bash
curl -Ls -X POST http://localhost:7787/object/item \
  --header 'Content-Type: application/json' \
  --data "{
            \"type\": 1,
            \"name\": \"test-bwca-add\",
            \"favorite\": false,
            \"fields\": [
              {
                \"Security Question\": \"Bitwarden Rules\"
              }
            ],
            \"login\": {
              \"uris\": [
                {
                  \"match\": 0,
                  \"uri\": \"https://twitter.com/login\"
                }
              ],
              \"username\": \"twitter@acmecorp.com\",
              \"password\": \"..........dog..........\"
            },
            \"reprompt\": true
    }"
```
```plantext
            \"organizationId\": null,
            \"collectionId\": null,
```
