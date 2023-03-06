Это приложение должно работать так:
Есть сервер на Go (Gin framework). Он сервит статические файлы фронтенда сделанного на React и отвечает на GraphQL запросы. Это CRUD где CRUDят app и его features. Например Zoom это app а его Features: screen sharing,drawing on top of shared screen, remote control.
Попытаюсь сгенерировать CRUD на фронте используя refine npm package а на бэке используя https://github.com/99designs/gqlgen. 

Based on https://github.com/yhagio/go_api_boilerplate
and https://github.com/bahdcoder/jwt-best-practices

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

**Used libraries:**
- [gin](https://github.com/gin-gonic)
- [gorm](https://gorm.io/docs/)
- [jwt-go](https://pkg.go.dev/gopkg.in/dgrijalva/jwt-go.v3?tab=doc)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [gqlgen](https://github.com/99designs/gqlgen)

- Components Diagram

#  How to

##  Modify

Make changes to schema by editing the file `schema.graphqls`

Backend
After making changes run the below command to auto generate the `models` and `resolvers`.
```shell
cd graphql
go run github.com/99designs/gqlgen generate
```
Programmer needs to implement the resolvers.

## Rebuild
on Windows for Linux   
```shell
build linux.bat
``` 
and you'll get exe in dst

## Deploy Windows->Linux
### Step 0: define variables if u haven't done
run `./env.bat`
make sure deployment_path exists
### Step 1:
```shell
deploy.bat
```
If u don't want to enter password:
### Step 0: Generate a public and private key pair
```shell
ssh-keygen -t rsa
```
Beware if u already have the keys the system should warn you. Do not override those.
### Step 1: Create dir & copy your public key to your remote server
ssh %server_user%@%server_addr% "mkdir ~/.ssh"
scp id_rsa.pub %server_user%@%server_addr%:~/
scp id_rsa.pub %server_user%@%server_addr%:~/.ssh/authorized_keys

Frontend
```shell
npm run graphql-codegen-watch
```
and in a separate window
```shell
npm start
```
## Run
from Windows on Linux
```sh
ssh -t %server_user%@%server_addr% "cd %deployment_path% && su root -c 'docker kill $(docker ps -q); docker-compose up'"
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
## Why
# Updated to go 1.19 (from 1.13)
because of https://github.com/golang/go/issues/44129
but it didn't fix the issue :(

https://www.youtube.com/watch?v=d2gfJ8UVPDo&t=402s

swag init -g app/app.go

go test \
./common/...\
./controllers/...\
./gql/...\
./infra/...\
./middlewares/...\
./repositories/...\
./services/...\
-coverprofile=coverage.txt -covermode=atomic