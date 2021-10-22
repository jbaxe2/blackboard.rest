**blackboard.rest**

The blackboard.rest library is an implementation of Blackboard's REST API's,
written in Go.  Not all aspects of Blackboard's REST API has been implemented,
and it is likely not all parts will be.

For example, various filter options may be passed to the REST API to better
emphasize how the results of the API call should be returned.  This includes
offset, limit, and fields specifiers.

Specific status code type errors also are not currently implemented.  While
there are no current plans to implement them, our implementation is flexible
enough to account for them in the future, if need be.

The interfaces provided by this library hopefully properly reflect what would
be expected by the users of the Blackboard REST API, or close enough thereof.
