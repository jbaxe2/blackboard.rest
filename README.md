**blackboard.rest**

The blackboard.rest library is an implementation of Blackboard Learn's REST
API's, written in Go.  Not all aspects of Blackboard's REST API has been
implemented, and it is likely not all parts will be (due to time constraints).

The interfaces provided by this library hopefully properly reflect what would
be expected by the users of the Blackboard REST API, or close enough thereof.

Documentation for the Blackboard Learn REST API's may be found here:

https://developer.blackboard.com/portal/displayApi

All work in this library stems from the information provided there, including
endpoints, types, protocols, etc.  A fundamental understanding of the REST API
documentation will help with a solid understanding of this library.

*Workflow*

This project aims to be lightweight, yet still flexible and powerful enough to
be meaningfully useful.  Workflow when dealing with this library is as follows:

1. Request a new OAuth2 instance to authenticate/authorize the app/user.

oAuth2 := oauth2.NewOAuth2 ("host", roundTripper)

The host is the URL for the Blackboard Learn host that the app has previously
been approved for use with.  The roundTripper variable is an instance of
http.RoundTripper, and will typically be nil.  It is established mostly to
provide easier testing of the library.

&nbsp;&nbsp;2a. Request an authorization code (for 3-legged OAuth).

response := oAuth2.AuthorizationCode (request, redirectUri, "clientId", "read")

The request is a pointer to an http.Request instance.  This request should be a
GET request to the oauth2/authorizationcode endpoint.  The redirectUri should
point to where the end user in this 3-legged OAuth request should be directed to
upon receiving an authorization code.  The client ID is the application key for
the application via Blackboard's developer portal.  Scope should be an
appropriate scope (such as read, write, or offline); this defaults to read.

Responses should be passed along to the user as-is.  This will include the
formatted redirect for user authorization (that also includes the code).

&nbsp;&nbsp;2b. Set the client ID and secret (for standalone tools).

oAuth2.SetClientIdAndSecret ("clientID", "secret")

Both the client ID and secret will be found in the developer portal.

3. Request an OAuth2 token.

token, err := oAuth2.RequestToken ("authorization_code", code, redirectUri)

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;OR

token, err := oAuth2.RequestToken ("client_credentials", "", redirectUri)

Depending on the use case, either authorization code (if step 2a was used above)
or client credentials (if step 2b was used above) will be provided.  In either
case, a "refresh_token" may be used instead to refresh the OAuth2 token (further
information on this is directed to the documentation above).  If the
authorization code is provided, the associated code must be provided as the
second argument.

If successful, the token variable will have a reference to the OAuth2 token
information.  This token will be used by further interactions with the API.
If the token request resulted in error, an OAuth2Error will be returned which
outlines the reason for the error.

4. Create a new Service instance.

service := api.NewService ("host", token, roundTripper)

Creating a service instance requires the host name of the Blackboard Learn
server.  This should be the same as provided in step 1.  The token is what was
returned from step 3.  Similarly to creating a new OAuth2 instance, a round
tripper instance may be passed in.  Again, this will typically be nil, and the
default http.DefaultTransport round tripper will be used.

5. Create a particular REST service instance.

courseService := blackboardRest.NewCourses (service)

To create the course service instance, pass in the service instance created in
step 4.

6. Use the REST service via the instance.

someCourse, err := courseService.GetCourse ("externalId:some_external_course_id")

In this example, we are retrieving a course from the courses service.  Endpoints
for various REST API calls will have a title descriptor, such as 'Get Course',
'Update Course', 'Create Course', etc.  Method names on the services in this
library attempt to mimic these titles as much as possible.  Parameters for the
methods will also match the endpoint path parameters as much as possible.

Calls may result in error for a variety of reasons.  If the error was provided
by the REST API as part of the response, a RestException will be created which
outlines the details of the error.

The following provides the workflow described above, appropriate error checking
not included.

oAuth2 := oauth2.NewOAuth2 ("host", nil)
oAuth2.SetClientIdAndSecret ("clientID", "secret")

token, _ := oAuth2.RequestToken ("client_credentials", "", nil)
service := api.NewService ("host", token, nil)
courseService := blackboardRest.NewCourses (service)

someCourse, err := courseService.GetCourse ("externalId:some_external_course_id")

One of the best places to find more information on usage of the library,
including setting optional request parameters and more thorough interactions
with services, is the various tests.

Issues may be opened if found.
