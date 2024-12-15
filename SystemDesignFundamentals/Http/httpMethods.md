# HTTP Methods

### GET :
* The GET method is used to retrieve data on a server
  
* We can get all the resources or a specific resource of a given type
  
* It doesn't include Request body because Client is not attempting
 to create or update the data.

      https://userservice.com/user/123


### POST :
  * The POST method is used to create a new resource.
    
  * Unlike get POST include a request body where the Client specifies the attributes
    of the resource to be created

             POST https://userservice.com/user
                Content-Type: application/json

                {
                "firstName": "John",
                "lastName": "Doe",
                "email": "johndoe@example.com",
                "password": "password123"
                } 

### PUT :
 * The Put method used to replace an existing resource with an updated version
 * this will replace the entire resource with the data provided in the Req body
 * this means any fieds not included in the request body are deleted and new fieds
   are added

             PUT https://userservice.com/user
             Content-Type: application/json
       
             {
             "firstName": "John",
             "email": "johndoe@example.com",
             }

             Before request : { "firstName": "John","lastName": "Doe"}
             After Request : { "firstName": "John","email": "johndoe@example.com"}


### PATCH :
 * The PATCH method is used to update an existing resource
 * It is similar to PUT, except that PATCH doesn't overWrite it's just update the
   Existing resource
        
                PUT https://userservice.com/user
                Content-Type: application/json

                {
                 "firstName": "JohnKarter",
                 "email": "JohnKarter@example.com",
                }

                Before request : { "firstName": "John","email": "john@example.com"}
                After Request : { "firstName": "JohnKarter","email": "JohnKarter@example.com"}


### DELETE :
 * The DELETE method is used to remove data from a database.
 * when a client sends a DELETE request it's requesting the specific
   resource at the url to be Removed

         DELETE https://userservice.com/user/123
        
### HEAD :
 * The HEAD method is used to retrieve the headers of the resource 
   without fetching the body Like GET
        
        HEAD https://userservice.com/user

        Response :
                    HTTP/1.1 200 OK
                    Content-Type: image/jpeg
                    Authorization: Bearer <token>
