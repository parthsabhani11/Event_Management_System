# Event Management System using Gin (Golang)

Welcome to our Event Management System built with Gin, a high-performance HTTP web framework for Golang. This system provides functionalities for managing events and users, including adding, removing, and updating events, as well as registering and logging in users, and registering or unregistering users for events.

## Installation

1. Make sure you have Go installed on your system. You can download and install it from [here](https://golang.org/dl/).

2. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/AdleeAfif/event-management-api
    ```

3. Navigate to the project directory:

    ```bash
    cd event-management-api
    ```

4. Install dependencies:

    ```bash
    go mod tidy
    ```

5. Build the project:

    ```bash
    go build .
    ```

6. Run the application:

    ```bash
    go run .
    ```

## Usage

Once the application is running, you can interact with the Event Management System using the following endpoints:

1. **Add Event**: 
   
   - Endpoint: `/events`
   - Method: POST
   - Description: Adds a new event to the system.
   - Request Body:
     ```json
     {
      "name": "Testing Event",
      "description": "A test event",
      "location": "A test location",
      "dateTime": "2025-01-01T15:30:00.000Z"
      }
     ```

2. **Remove Event**: 

   - Endpoint: `/events/:id`
   - Method: DELETE
   - Description: Removes the event with the specified ID from the system.

3. **Update Event**: 

   - Endpoint: `/events/update/:id`
   - Method: PUT
   - Description: Updates the details of the event with the specified ID.
   - Request Body:
     ```json
     {
      "name": "Not Testing Event",
      "description": "A non test event (updated!)",
      "location": "A non test location",
      "dateTime": "2025-01-01T15:30:00.000Z"
      }
     ```

4. **Signup User**: 

   - Endpoint: `/signup`
   - Method: POST
   - Description: Registers a new user in the system.
   - Request Body:
     ```json
     {
         "email": "email",
         "password": "password"
     }
     ```

5. **Login User**: 

   - Endpoint: `/login`
   - Method: POST
   - Description: Logs in an existing user.
   - Request Body:
     ```json
     {
         "username": "username",
         "password": "password"
     }
     ```

6. **Register User for Event**: 

   - Endpoint: `/events/:id/register/`
   - Method: POST
   - Description: Registers the logged-in user for the event with the specified ID.

7. **Unregister User from Event**: 

   - Endpoint: `/events/:id/register`
   - Method: DELETE
   - Description: Unregisters the logged-in user from the event with the specified ID.

8. **See all users (ADMIN)**: 

   - Endpoint: `/users`
   - Method: GET
   - Description: Display all the users available in the database.

9. **Remove single user by ID (ADMIN)**: 

   - Endpoint: `/users/:id`
   - Method: Delete
   - Description: Remove a user from the system permanently.  

## Contributing

Contributions are welcome! If you find any bugs or want to add new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Special thanks to the Gin contributors for creating such a powerful framework for building web applications in Go.
