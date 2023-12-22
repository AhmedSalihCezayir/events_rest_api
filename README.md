# Events Rest Api #
This project is an event management system with a RESTful API using Go, designed to facilitate various event-related operations. It includes endpoints for retrieving event details, managing user registrations, and handling user authentication. The API ensures data security through authentication mechanisms (JWT) and provides specific admin-only functionalities.

In the project there are 3 types of endpoint access which are public, auth and admin. Public endpoints are accessible by everyone whereas to access an authenticated endpoint, an authentication token needs to be passed in request headers. Admin endpoints are only accessible by admin users and they are used to see and manage all users. Details of these endpoints can be seen below:

| Endpoint | Access Level | Description |
| --- | --- | --- |
| `GET api/v1/events` | `Public` | Retrieve a list of events. |
| `GET api/v1/events/:id` | `Public` | Retrieve details of a specific event by ID. |
| `GET api/v1/events/:id/attendees` | `Public` | Retrieve the list of attendees for a specific event. |
| `POST api/v1/signup` | `Public` | Sign up for a new user account. |
| `POST api/v1/login` | `Public` | Log in to an existing user account. |
| `POST api/v1/events` | `Auth` | Create a new event. |
| `PUT api/v1/events/:id` | `Auth` | Update details of a specific event by ID. |
| `DELETE api/v1/events/:id` | `Auth` | Delete a specific event by ID. |
| `POST api/v1/events/:id/register` | `Auth` | Register for a specific event. |
| `DELETE api/v1/events/:id/register` | `Auth` | Unregister from a specific event. |
| `GET api/v1/registered-events` | `Auth` | Retrieve a list of events a user is registered for. |
| `PUT api/v1/update-user` | `Auth` | Update user details (e.g., profile information). |
| `GET api/v1/users` | `Admin` | Retrieve a list of users (admin access required). |
| `DELETE api/v1/users/:id` | `Admin` | Delete a user by ID (admin access required). |
