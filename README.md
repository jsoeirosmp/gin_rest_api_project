# Student Management System

This is a simple application that uses Gin, Gorm, and Docker to perform CRUD operations on a database of fictional students. The data in the database can be viewed using PgAdmin, which is also run on Docker.

## Dependencies

The following dependencies were used in this application:

- Gin: A web framework for Go that provides a set of features for building web applications.
- Gorm: A Golang ORM library that helps with database management.
- Docker: A platform that allows developers to easily create, deploy, and run applications in containers.
- Postman: A powerful tool for testing APIs.

## Getting Started

To run the application, follow these steps:

1. Clone the repository to your local machine.
2. Open a terminal window and navigate to the project directory.
3. Run the command `docker-compose up` to start the database and PgAdmin.
4. Open a web browser and go to `http://localhost:54321` to view PgAdmin. Use the following credentials to log in:
   - Email: `admin@admin.com`
   - Password: `admin`
5. Create a new server in PgAdmin with the following settings:
   - General:
     - Name: `YOUR_SERVER_NAME`
   - Connection:
   - For the host, type in `docker-compose exec postgres sh`, then `hostname -i`, and copy the adress which will be pasted on PgAdmin`s setup as shown below:
     - Host: `YOUR_HOST_IP`
     - Port: `5432`
     - Maintenance database: `root`
     - Username: `root`
     - Password: `root`
6. Save the server settings and connect to the database.
7. Open another terminal window and navigate to the project directory.
8. Run the command `go run Main.go` to start the application.
9. Open a web browser or Postman and go to `http://localhost:8080` to use the application.

## Usage

The application provides the following endpoints:

- GET `/alunos`: Retrieves a list of all students in the database.
- GET `/alunos/:id`: Retrieves a specific student from the database by ID.
- POST `/alunos`: Creates a new student in the database, the json structure is as follows:
  `{
    "nome":"Student's Name",
    "cpf":"document_number",
    "rg":"rg_number"
  }`
- PUT `/alunos/:id`: Updates an existing student in the database by ID, with the same JSON structure shown above.
- DELETE `/alunos/:id`: Deletes a specific student from the database by ID.

## Conclusion

This application demonstrates how to use Gin, Gorm, Docker, and Postman to create a simple CRUD application for managing student data and testing API endpoints. By following the steps outlined above, you should be able to get the application up and running on your local machine in no time.
