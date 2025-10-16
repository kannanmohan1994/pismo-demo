## Description
XM - An application to handle companies. It should provide the following operations:
- Create
- Patch
- Delete
- Get (one)

## Installation
- Install Golang => https://golang.org/doc/install
- Install Docker => sudo snap install docker

### Local Setup
- Clone: git clone https://github.com/kannanmohan1994/xm-demo.git
- Setup infra (postgres, pgadmin, kafka): docker-compose up 
- Run migrations: make migrate-up
- Create kafka topics: bash create-topics.sh (OPTIONAL)
- Run application: make tidy, make run 
- Testout API's: Go to http://localhost:9000/swaggerui/ 

### API testing in swagger
- POST /v1/user/register => Generate user with username and password combination. Get access token in response
- POST /v1/user/login => Regenerate access token for username and password combination
- Set access token in swagger-ui in Authorize to try-out following API's
- POST /v1/company => Create company and returns company details with id. AUTHORIZED
- GET /v1/company/{company-id} => Fetch company for company-id. AUTHORIZED
- PATCH /v1/company/{company-id} => Patch company for company-id. AUTHORIZED
- DELETE /v1/company/{company-id} => Delete company for company-id. AUTHORIZED

### Useful Commands
- Download and Install all the dependent packages
```bash
    make tidy
```
- To Run the Server:
```bash
    make run
```
- To Run Tests:
```bash
    make test
```
- To Run Lint:
```bash
    make lint
```
- To Setup Swagger Docs:
```bash
    make swagger
```
- To Run/Remove migrations:
```bash
    make migrate-up
    make migrate-down
```

### Environment Variables Config
Configs are present in .env file

## Tech Stack

*Database:* Postgres
*Server:* Golang
