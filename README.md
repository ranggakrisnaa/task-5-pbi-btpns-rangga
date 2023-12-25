# User-Profile API with Go-Lang

## Auth Endpoint

> Register User
> POST http://localhost:3001/api/v1/auth/register

* Create New User.
* Request Body :
  ```
  {
     "username": "danuputra",
     "email": "danuputra",
     "password": "danuputra"
  }
  ```
* Response Body :
     * 201 Created
       ```
       {
         "message": "User registered successfully",
         "success": true
       }
       ```
> Login User
> POST http://localhost:3001/api/v1/auth/login

* Authenticates a user.
* Request Body :
  ```
  {
    "email": "danuputra",
    "password": "danuputra"
  }
  ```
* Response Body :
     * 200 OK
       ```
       {
          "message": "User Login successfully",
          "success": true,
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDM1ODgxOTgsImp0aSI6IjIifQ.2D7F-m-U53EYoaNpn_7T4xxHgxtv5lxVk_8PjWqRVpI"
       }
       ```

## User Endpoint
