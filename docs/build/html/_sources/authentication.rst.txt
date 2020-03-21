Authentication
==============
By default, Go-Web provides two ways for authenticating users:

*  JWT-based authentication
*  basic (base) authentication

JWT Authentication
------------------
Commonly used to authenticate users thought mobile applications or a SPA, JWT authentication is implemented by function JWTAuthentication of controller AuthController or, in Go-Web terms,
by endpoint **AuthController@JWTAuthentication**

The JSON structure used to represent credentials of a user must conform to JSON

.. code-block:: none

    {
        "username": <string, mandatory>,
        "password": <string, mandatory>
    }

The result of a successful login attempt with this type of authentication is a HTTP response containing a JWT token.
Resource access can be restricted only to authenticated users by adding middleware Auth to specific routes.

Basic (base) authentication
---------------------------
Basic, or base, authentication is the simplest way to authenticate users for service access; this method is implemented by endpoint
**AuthController@BasicAuth**

The base authentication requires the same data structure as JWT-based method and routes can be protected by using middleware BasicAuth.


