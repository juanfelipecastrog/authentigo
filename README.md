# Coding Exercise: Simple OAuth2 Server

## Running program & Command-line Arguments

```sh
go run main.go -privateKeyPath yourCustomPathWithPrivateKey/jwtRS256.key
```

The program accepts the following command-line argument:

```sh
-privateKeyPath: Specifies the file path for the private key. If not provided, a default private key path will be used
```

### Port: 8080

## Endpoints:
### IssueAccessToken:
This endpoint issues a JWT access token in response to a request to grant client credentials with basic authentication.

#### Endpoint: /api/token

#### HTTP Method: Get

#### Request
The request to this endpoint should include basic authentication headers with the client ID and client secret.
this program only uses local("hardcoded") users you can use any of the following to authenticate using basic auth

-  "user":  "password",
-  "user2": "password2",
-  "user3": "password3",

Example using Curl
```sh
curl -X GET http://localhost:8080/api/token \
     -u "user:password"
```


#### Response

- **200 OK:** If the client credentials are valid, an access token is generated successfully. The response body contains the signed access token as plain text.

- **400 Bad Request:** If the request does not include valid basic authentication headers.

- **401 Unauthorized:** If the provided client credentials are invalid.

- **500 Internal Server Error:** If there is an internal error while creating the access token.

************************

### IntrospectToken:
This endpoint verifies and displays information about a JWT access token.

#### Endpoint: /api/token/validator

#### HTTP Method: Get

#### Request
The request to this endpoint should include the access token in the token form parameter.
- The key should be **token**
  
Example using Curl

```sh
curl -X GET http://localhost:8080/api/token/validator \
     -F "token=<access_token>"
```

#### Response

- **200 OK:** If the access token is valid, the response body contains the token's claims as a JSON object.

- **400 Bad Request:** If the request does not include a valid access token.

- **401 Unauthorized:** If the server's private key is missing or invalid, or if the access token's signature is invalid.

- **500 Internal Server Error:** If there is an internal error while processing the token.

************************

### ListSigningKeys:
This endpoint returns the list of signing keys(Public keys that have been used to sign tokens) in JSON Web Key (JWK) format.

#### Endpoint: /api/keys

#### HTTP Method: Get

#### Request
No request parameters are required for this endpoint.

Example using Curl

```sh
curl -X GET http://localhost:8080/api/keys
```

#### Response

- **200 OK:** The response body contains a JSON object representing the list of signing keys in JWK format.

- **500 Internal Server Error:** If there is an internal error while generating the JWK response.
