# **Documentation**
This Application served on HTTP with GraphQL as the top of layer\
and has built in SQL migration app\
Base URL : `localhost:3000`
## **How to run**
### **Prereqquisite**
1. \>= Go.1.18
2. Docker

### **Steps**
1. Clone this repo
> git clone git@github.com:vincen320/user-service-graphql.git
2. Run docker
3. Run docker compose command
> docker compose up -d
4. Run migration
> ./bin/app migration up
5. Run the application
> ./bin/app run

## **Test the Application**
Endpoint:
```http
POST /v1/graphql
```
## **Query**
Get All Users
```json
{
    "query" : "{user{id name email age address salary hobbies{id name}}}"
}
```

Get User By ID
```json
{
    "query" : "{user(id:1){id name email age address salary hobbies{id name}}}"
}
```

Get Hobbies
```json
{
    "query" : "{hobby{id name}}"
}
```

## **Mutation**
Create User
```json
{
    "query" : "mutation createUser($user: CreateUserParam){user(user:$user){id name address}}",
    "operationName": "createUser",
    "variables":{
        "user":  {
                "name": "EmailPassword",
                "email" : "email@mail.com",
                "password" : "password",
                "age": 16,
                "address": "Jl.Tebet No.18",
                "salary": 7700000
            }
    }
}
```

User Login
```json
{
    "query" : "mutation userLogin($login: UserLoginParam){login(login:$login){token}}",
    "operationName": "userLogin",
    "variables":{
        "login":  {
                "email": "budiana@mail.com",
                "password":"password"
            }
    }
}
```

### GraphQL Object types and fields
```graphql
type User{
    Id: Int!
    Name: String!
    Email: String!
    Age: Int!
    Address: String!
    Salary: Float!
    Hobbies: [Hobby]!
}

type Hobby{
    Id: Int!
    Name: String!
}

type CreateUserParam{
    Name: String!
    Email: String!
    Password: String!
    Age: Int!
    Address: String!
    Salary: Float!
}

type UserLoginParam{
    Email: String!
    Password: String!
}
```
