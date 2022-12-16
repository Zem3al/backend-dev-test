## Prerequisites

- Install Docker [Follow the link](https://docs.docker.com/desktop/install/windows-install/)
- Install Go [Follow the link](https://go.dev/) - you should install verion 1.18+

## Project structure
```
├───cmd
│   └───app
├───configs
├───deployments
├───docs
├───internal
│   ├───app
│   │   └───handler
│   ├───config
│   ├───model
│   ├───persistence
│   │   ├───gorm
│   │   └───mock
│   ├───service
│   └───test
└───migration
```

## Setting Up

## Step 1 Run docker compose (setting up database )  : 
    - /deployments/docker-compose.yaml

## Step 2 Setting up your configuration :
    - (configs/.env)

## Step 3 Migration Table:
```
go run cmd/app main.go migration up
```

## Step 4: Serve the server

```--
go run cmd/app main.go serve
```

## API

### Creat Hub

```
curl --location --request POST 'http://localhost:8080/hub' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "hub",
    "location": "Ha Noi"
}'

Reponse: {"hub_id":"9e47f3e8-cfb5-4148-9854-42ef134aa11c","name":"hub","location":"Ha Noi"}
```

### Creat Team

```
curl --location --request POST 'http://localhost:8080/team' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "team",
    "type": "tester",
    "hub_id": "9e47f3e8-cfb5-4148-9854-42ef134aa11c"
}'

Reponse: {"team_id":"2eb8226b-58c7-498e-91dc-3b6d45d02db8","name":"team","Type":"tester","hub_id":"9e47f3e8-cfb5-4148-9854-42ef134aa11c"}
```

### Creat User

```
curl --location --request POST 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "minh",
    "age": 18,
    "team_id": "2eb8226b-58c7-498e-91dc-3b6d45d02db8"
}'

Response: {"user_id":"c35010f1-b9e6-40ec-ba40-bd5f97bcd446","name":"minh","age":18,"team_id":"2eb8226b-58c7-498e-91dc-3b6d45d02db8"}
```

### Search API query location AND type
```
curl --location --request GET 'http://localhost:8080/search?location=Ha&&type=normal' \
--data-raw ''
```

