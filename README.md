Это приложение должно работать так:
Есть сервер на Go (Gin framework). Он сервит статические файлы фронтенда сделанного на React и отвечает на GraphQL запросы. Это CRUD где CRUDят app и его features. Например Zoom это app а его Features: screen sharing,drawing on top of shared screen, remote control.
Попытаюсь сгенерировать CRUD на фронте используя refine npm package а на бэен используя https://github.com/99designs/gqlgen. 


You can replace secrets in dst/.env with ones generated. You can generate secrets running

```sh
pip3 install pipenv
cd ./generate_secrets
pipenv shell
python ./generate_secrets.py
```

## Secrets:

JWT (JSON Web Token) is a standard for creating tokens that securely transmit information between parties. A JWT SIGN KEY is a secret key that is used to sign the JWT, providing authentication and ensuring that the token has not been tampered with.

A HAMC (Hash-based Message Authentication Code) KEY is a secret key that is used to generate a digital signature for a message, allowing the recipient to verify that the message has not been modified during transmission.

PEPPER is a term that is often used in conjunction with password storage, where it refers to a secret value that is added to the password before it is hashed (encrypted) and stored. This makes it more difficult for an attacker to guess or crack the password, even if they have access to the hashed version of the password.

#  How to

## Run

0 run
```
go run github.com/99designs/gqlgen generate
```

1 run `build linux.bat` and you'll get exe in dst

2 copy `dst` and `docker-compose.yml` to your server

3 Run 
```sh
docker-compose up
```

## Use
1 Send POST (e.g. with POSTMAN) to `/api/register` with json
```json
{"email":"vsf","password":"wawf"}
```
copy token from the response and
2 Go to graphql-playground (`/graphql`)
in headers
```json
{
"Authorization":"Bearer <paste token here>"
}
```
3 You can 
```graphql
query {user(id:1){email}}
```

```graphql
mutation {
updateUser(input:{firstName:"Ivan",email:"dwde"}){id}
}
```


##  Make changes to schema by editing the file `schema.graphqls`

After making changes run the below command to auto generate the `models` and `resolvers`.
```shell
go run github.com/99designs/gqlgen generate
```
Programmer needs to implement the resolvers.