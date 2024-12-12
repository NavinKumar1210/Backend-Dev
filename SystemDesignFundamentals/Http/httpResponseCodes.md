# HTTP Response Codes

## Response Types :
* Successful response = 200-299

* Redirection response = 300-399

* Client error response = 400-499

* Server error response = 500-599
                    

## Successful responses : 
* 200 - OK - The request succeeded.

* 201 - created - The request succeeded and the resource created

* 202 - accepted - The request has been receiver but not yet processed

* 204 - No Content - The server successfully processed the request, but there is no content to return.

## Redirection responses :
* 301 - Moved Permanently - The resource has been permanently moved to a new URL.

* 302 - Found - The resource is temporarily located at a different URL.

* 304 - Not Modified - The resource has not been modified since the last request.

## Client Errors responses :
* 400 - Bad Request - The server could not understand the request due to invalid syntax.

* 401 - Unauthorized - The request requires user authentication.

* 403 - Forbidden - The server understands the request but refuses to authorize it.

* 404 - Not Found - The server cannot find the requested resource.

* 405 - Method Not Allowed - The request method is not supported for the requested resource.

* 408 - Request Timeout - The server timed out waiting for the request.

* 429 - Too Many Requests - The user has sent too many requests in a given amount of time.

## Server Errors responses :
* 500 - Internal Server Error - The server encountered an unexpected condition that prevented it from fulfilling the request.

* 502 - Bad Gateway - The server received an invalid response from the upstream server.

* 503 - Service Unavailable - The server is not ready to handle the request, often due to maintenance or overload.

* 504 - Gateway Timeout - The server did not receive a timely response from the upstream server.