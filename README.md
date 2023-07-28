Go Backend Server with MongoDB
This is a basic Go backend server that utilizes MongoDB for data storage. It provides the following operations:

Login: Allows users to log in using their registered credentials.

Signup: Allows users to register with a username, password, and email.

Get User by ID: Retrieves a user's information by their unique ID.

Get All Users: Retrieves a list of all users from the database.

Delete User: Deletes a user by their unique ID.

Update User Info: Updates a user's information by their unique ID.

Requirements
Go (Golang) installed on your machine.
MongoDB server running or a cloud-based MongoDB URL.
Installation and Setup
Clone the repository:

bash
Copy code
git clone https://github.com/your-username/go-backend-mongodb.git
cd go-backend-mongodb
Install dependencies:

go
Copy code
go mod tidy
Create a .env file in the project root and specify the MongoDB URL and other environment variables:

makefile
Copy code
MONGO_URL=mongodb://your-mongodb-url
SECRET_KEY=your-secret-key-for-JWT
PORT=8000
Run the server:

go
Copy code
go run main.go
API Endpoints
Signup
Endpoint: POST /api/signup
Description: Allows users to register with a username, password, and email.
Login
Endpoint: POST /api/login
Description: Allows users to log in using their registered credentials.
Get User by ID
Endpoint: GET /api/user/:id
Description: Retrieves a user's information by their unique ID.
Get All Users
Endpoint: GET /api/users
Description: Retrieves a list of all users from the database.
Delete User
Endpoint: DELETE /api/user/:id
Description: Deletes a user by their unique ID.
Update User Info
Endpoint: PUT /api/user/:id
Description: Updates a user's information by their unique ID.
Error Handling
The server returns appropriate HTTP status codes and error messages for different scenarios. Be sure to handle these responses on the client-side.

Security
Passwords are securely hashed before storing in the database.
JWT-based authentication is implemented for secure user login.
Unit Tests
The code includes unit tests for the API endpoints to ensure their correctness and robustness.

Technologies Used
Go (Golang)
MongoDB
Gin Web Framework
JWT (JSON Web Token) for authentication
Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or bug fixes.
