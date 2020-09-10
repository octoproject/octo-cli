# Get users service

Following are step-by-step instructions for how to deploy service using Octo-CLI


## Step1: Create the database

1. Create on the cloud or local database 
```
$ docker run -d \
    --name users \
    -e POSTGRES_PASSWORD=123456 \
    -p 5432:5432  \
    postgres
```

2. Import the sample data 
```
$ git clone https://github.com/octoproject/octo-cli
$ cd    
$ cat  examples/get-users/users.sql | docker exec -i users psql -U postgres 
 ```

## Step2: Generate Octo configuration
 Generate Octo configuration by running
   ```
$ go run . init 
 ```
or skip this step by using the file under examples/get-users/get-users.yml and you only need to change the database credential and specify the platform in the file.

## Step3:  Create a new service

Create a new service
 ``` 
 $ go run .  create -f examples/get-users/get-users.yml 
 ```

 ## Step4:  Build function Docker container

```
go run .  build -f examples/get-users/get-users.yml  --prefix dev.local
 ```

 ## Step5:  Deploy the service
 ```
 # openfaas

$ go run .  deploy -f examples/get-users/get-users.yml -i dev.local/get-users:latest   -u admin -p 41d21dfa77da9 -g http://127.0.0.1:8080

# knative
$ go run .  deploy -f examples/get-users/get-users.yml  --pullPolicy IfNotPresent -i dev.local/get-users:latest 
 ```

## Step6: Test  the service
```
# request 

 # openfaas
curl --location --request GET 'http://localhost:8080/function/get-users?name=alice'

# knative
curl --location --request GET 'http://get-users.default.example.com/get-users?name=alice'

# response 
{
    "data": [
        {
            "id": 12345,
            "active": true,
            "name": "alice"
        }
    ]
}
```
