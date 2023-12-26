# User-Profile API with Go-Lang
> [!NOTE]
> Postman Documentation: [here](https://documenter.getpostman.com/view/29492816/2s9Ykt4eHW)
## Auth Endpoint

> Register User - 
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
> Login User -
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

> Update User -
> PUT http://localhost:3001/api/v1/user/:id

* Updates User information.
* Requires authentication via JWT Token.
* Request Body :
  ```
  {
    "username": "gg",
    "email": "gg",
    "password": "rangga123"
  }
  ```
* Response Body :
    * 200 OK
      ```
      {
        "message": "User Updated Successfully",
        "success": true
      }
      ```

> Delete User -
> DELETE http://localhost:3001/api/v1/user/:id

* Delete a User
* Requires authentication via JWT Token.
* Request Body :
  ```
  { null }
  ```
* Response Body :
   * 200 OK
     ```
      {
        "message": "User Deleted Successfully",
        "success": true
      }
      ```

## Photo Endpoint

> Upload Photo - 
> POST http://localhost:3001/api/v1/photos

* Upload User Photo Profile.
* Requires authentication via JWT Token.
* Request Body :
  ```
  {
     "photo_file": "/downloads/image/bantengmerah.png",
     "title": "ini merupakan photo saya",
     "caption": "photo"
  }
  ```
* Response Body :
     * 201 Created
       ```
       {
         "message": "Photo uploaded successfully",
         "success": true
       }
       ```
> Get Photo - 
> GET http://localhost:3001/api/v1/photos

* Get All User Photo Profile.
* Requires authentication via JWT Token.
* Request Body :
  ```
  {
     null
  }
  ```
* Response Body :
     * 200 OK
       ```
       {
         "data": [
        {
            "ID": 5,
            "CreatedAt": "2023-12-25T23:30:18.127115+07:00",
            "UpdatedAt": "2023-12-25T23:30:18.127115+07:00",
            "DeletedAt": null,
            "title": "ini adalah profile foto saya",
            "caption": "foto",
            "photo_url": "http://localhost:8080/api/v1/photos/1703521818078112535-screenshot.png",
            "user_id": "1",
            "User": {
                "ID": 1,
                "CreatedAt": "2023-12-25T23:30:02.815548+07:00",
                "UpdatedAt": "2023-12-25T23:30:02.815548+07:00",
                "DeletedAt": null,
                "username": "danuputra",
                "email": "danuputra",
                "password": "$2a$10$m39D.L3S3eQQQlnTnbHMze35LW7X4CyXKOnh1pkZOR5.WUvZd8SqK"
            }
        },
           ],
        "message": "Successfully retrieve data",
        "success": true
       }
       ```
> Update Photo - 
> PUT http://localhost:3001/api/v1/photos/:id

* Update User Photo Profile.
* Requires authentication via JWT Token.
* Request Body :
  ```
  {
     "photo_file": "/downloads/image/bantengmerah.png",
     "title": "ini merupakan photo saya",
     "caption": "photo"
  }
  ```
* Response Body :
     * 200 OK
       ```
       "data": {
          "ID": 8,
          "CreatedAt": "2023-12-26T11:36:49.70466+07:00",
          "UpdatedAt": "2023-12-26T11:46:49.219493791+07:00",
          "DeletedAt": null,
          "title": "ini merupakan photo saya",
          "caption": "photo",
          "photo_url": "http://localhost:3001/api/v1/photos/1703565409660618324-bantengmerah.png.png",
          "user_id": "1",
          "User": {
            "ID": 1,
            "CreatedAt": "2023-12-25T23:30:02.815548+07:00",
            "UpdatedAt": "2023-12-25T23:30:02.815548+07:00",
            "DeletedAt": null,
            "username": "danuputra",
            "email": "danuputra",
            "password": "$2a$10$StjEXtgXzdHSLfnUCdT5GOCwbqz1yM9r2Y/WQe7vOBqvZshdqAyMG"
        }
      },
      "message": "Photo updated successfully",
      "success": true
       ```
> Delete Photo - 
> DELETE http://localhost:3001/api/v1/photos/:id

* Delete User Photo Profile.
* Requires authentication via JWT Token.
* Request Body :
  ```
  {
     null
  }
  ```
* Response Body :
     * 200 OK
       ```
       {
         "message": "Photo Deleted successfully",
         "success": true
       }
       ```
