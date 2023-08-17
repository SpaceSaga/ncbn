# NCBN (Not Correct But Now)

## NCBN 
is a web API written in Go that connects to a database server specified in the app-settings.json configuration file and executes stored procedures.

Getting Started
These instructions will help you set up and run the NCBN project on your local machine for development and testing purposes.

### Prerequisites

    Go programming language
    PostgreSQL or SQL Server database

### Dependencies

    github.com/lib/pq (for PostgreSQL)
    github.com/denisenkom/go-mssqldb (for SQL Server)
    github.com/spf13/viper
    github.com/swaggo/http-swagger

### Configuration
 Update the app-settings.json file with the appropriate values for your database:

```json
{
"DBDriverName": "postgres",
"DBAddress": "your_host",
"DBUser": "your_user",
"DBPassword": "your_password",
"DBName": "your_dbname"
}
```

Replace your_host, your_user, your_password, and your_dbname with the appropriate values for your PostgreSQL or SQL Server database. Change the DBDriverName to "sqlserver" if you are using SQL Server.
Running the Project
To run the project, navigate to the project directory and execute the following command:

```bash
go run main.go
```

The API server will start and listen on given port