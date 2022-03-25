```bash
curl -Ls -X POST http://localhost:7787/object/item \
  --header 'Content-Type: application/json' \
  --data "{
            \"organizationId\": null,
            \"collectionId\": null,
            \"folderId\": null,
            \"type\": 1,
            \"name\": \"bwca-new-original\",
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
              \"password\": \"....tweet....\",
              \"totp\": null
            },
            \"reprompt\": 0
    }"
```

```bash
curl -Ls -X POST http://localhost:7787/object/item \
  --header 'Content-Type: application/json' \
  --data "{
            \"organizationId\": null,
            \"collectionIds\": null,
            \"folderId\": null,
            \"type\": 1,
            \"name\": \"bwca-item-add\",
            \"notes\": \"Basic Note.\",
            \"favorite\": false,
            \"fields\": [],
            \"login\": {
              \"username\": \"twitter@acmecorp.com\",
              \"password\": \"..........dog..........\"
            },
            \"secureNote\": null,
            \"card\": null,
            \"identity\": null,
            \"reprompt\": 0
    }"
```

```json
{"success":true,"data":{"object":"item","id":"bd6cd799-e10a-4477-b180-ae6300f7cf18","organizationId":null,"folderId":null,"type":1,"reprompt":0,"name":"bwca-item-add","notes":"Basic Note.","favorite":false,"login":{"username":"twitter@acmecorp.com","password":"..........dog..........","totp":null,"passwordRevisionDate":null},"revisionDate":"2022-03-25T15:02:14.681Z","deletedDate":null}}
```

```bash
curl -Ls -X POST http://localhost:7787/object/item \
  --header 'Content-Type: application/json' \
  --data "{
           \"collectionId\": \"\",
           \"favorite\": false,
           \"fields\": null,
           \"folderId\": \"\",
           \"login\": {
             \"password\": \"....iamhere....\",
             \"totp\": \"\",
             \"uris\": null,
             \"username\": \"cmaahs\"
           },
           \"name\": \"bwca-item-added\",
           \"notes\": \"This is a note\",
           \"organizationId\": \"\",
           \"reprompt\": 0,
           \"type\": 1
    }"
```
