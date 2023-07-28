**Go Backend Server with MongoDB**

This is a basic Go backend server that utilizes MongoDB for data storage. It provides the following operations:

1.Login: Allows users to log in using their registered credentials.

2.Signup: Allows users to register with a username, password, and email.

3.Get User by ID: Retrieves a user's information by their unique ID.

4.Get All Users: Retrieves a list of all users from the database.

5.Delete User: Deletes a user by their unique ID.

6.Update User Info: Updates a user's information by their unique ID.

Requirements
Go (Golang) installed on your machine.
MongoDB server running or a cloud-based MongoDB URL.

1.Clone the repository
```
git clone https://github.com/your-username/go-backend-mongodb.git
cd go-backend-mongodb
```
2.Install dependencies
```
go mod tidy 
```
3.Create a .env file in the project root and specify the MongoDB URL and other environment variables:
```
MONGO_URL=mongodb://your-mongodb-url
SECRET_KEY=your-secret-key-for-JWT
PORT=8000
```
4.Run the server:
```
go run main.go
```


Contributing 
Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or bug fixes.
