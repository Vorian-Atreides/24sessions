# 24sessions's assessment

Develop geolocation service based on user IP address.

Frontend contains one page which shows the IP address of the current user and button "Get geolocation". When the user clicks the button, we load location info (city and country) via Ajax request to the backend.

Backend retrieves information from the JSON service of ipinfo.io with caching into the database. Only city and country must be cached and returned to Frontend. Let's assume that we have MySQL database "test" on localhost with user "test_user" and with password "secret"

The backend should be built on Silex (latest version) with ActiveRecord ORM (https://packagist.org/packages/php-activerecord/php-activerecord) using Composer.

For Frontend, jQuery can be used.

Please, when you are done put your code into one of the online Git repositories (e.g. Bitbucket) and send me a link to it.

Notes:
- Add documentation as 'readme'.
- Tests should cover main functionality.

May the force be with you.

## Usage

### With Docker

The whole solution has been packaged with Docker and can be run with one command:
> $> docker-compose up dev

It will run:
- MySQL container, using the port 3306
- Backend container, using port 8080
- Frontend container, using port 3000

### Without Docker

#### Backend

If you want to build the application by yourself, you will need to have a working Golang workspace and place the application at:
> github.com/Vorian-Atreides/24sessions

Then you should be able to build the backend application:
```
$> cd backend
$> go build ./cmd/api
```

The default parameters should be compliant with the assessment requirements, but you can customise them to match your configuration if needed:

```
$> ./api -h
$>
NAME:
   api - 24sessions assessment

USAGE:
   api [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --api_token value    APIToken to allow the authenticate the application with IpInfo [$API_TOKEN]
   --db_username value  Username for the dabatase connection (default: "test_user") [$DB_USERNAME]
   --db_password value  Password for the database connection (default: "secret") [$DB_PASSWORD]
   --db_host value      Host where is located the database (default: "localhost") [$DB_HOST]
   --db_name value      Database name used to store the data (default: "test") [$DB_NAME]
   --help, -h           show help
   --version, -v        print the version
```

#### Frontend

The application has been created with create-react-app, which mean that you can run the dev environment with npm:
```
$> cd frontend
$> npm run start
```

The frontend solution is expecting to contact the backend at
> localhost:8080

## Solution

The solution is rather simple, with a REST API in the backend and a ReactJS app in frontend.

It could be optmised with an NGINX cache for the backend, avoiding to hit the DB for each redundant query.

The solution could be extended with another microservice to abstract the usage of ipinfo, keeping the REST API logic focused on the persistent layer.