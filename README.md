### Retrieve access and refresh tokens by user credentials

```shell
curl --location 'localhost:3000/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "user",
    "password": "user"
}'
```