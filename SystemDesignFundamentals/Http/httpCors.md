# Cross-Origin Resource Sharing ( CORS ) :
* CORS is a Security feauture in web browser that allows a web page to request Resources from other Server than it's own.It is 
  important because web browsers enforce same-origin-policy that restrict cross origin requests by default to protect users 
  from malicious websites

# CORS headers :

* Access-Control-Allow-Origin - Specifies which origins can access the resource. eg: https://example.com
  

* Access-Control-Allow-Methods - HTTP methods (e.g : GET, POST, PUT, DELETE) that are allowed when
  accessing the resource.

* Access-Control-Allow-Credentials: Indicates whether the request can include user credentials 
      like cookies and HTTP authentication.
                            

# Preflight Requests : 
* For certain types of requests ( PUT,PATCH ), the browser sends a preliminary request 
  (OPTIONS method) to check if the actual request is allowed

# Security :
* Be cautious with Access-Control-Allow-Origin: * as it allows any website to access your resources,
  which can be a security risk.