**blackboard.rest**

## 0.2.3+3
- Tweaked some code for OAuth2 service and token to better work with server
responses.

## 0.2.3+2
- More work on developing the general Service type and the REST exception.

## 0.2.3+1
- More work on developing OAuth2, the general Service type, and errors.
- Start of considering the Courses service.

## 0.2.3
- Developed and refined the OAuth2 service and token types, error types, base
service type, and various tests.

## 0.2.2+2
- Redefined endpoints from mutable maps to immutable constants.
- More work on OAuth2 service.

## 0.2.2+1
- Work on OAuth2 service.

## 0.2.2
- Removed most, if not all, of old code that will not be used with redesign.

## 0.2.1
- More work on redesign.

## 0.2.0 - ** BREAKING CHANGES **
- Some redesign of the library.

## 0.1.8
- Updating code base to newer REST API definitions.

## 0.1.7
- Added some tests.
- Included group attempt ID for column attempts.

## 0.1.6+6
- Resolved a column attempt factory parsing issue.

## 0.1.6+5
- Worked to resolve issues with course user parsing, including for loading
course and user memberships.

## 0.1.6+4
- Added inclusion of user information when loading enrollments.

## 0.1.6+3
- Added ability to retrieve group sets for a course ID.

## 0.1.6+2
- More work on course groups.

## 0.1.6+1
- Work on course groups factory.

## 0.1.6
- Renaming and restructuring of project.
- Tweaks (re-establishing imports, formatting, etc.) to code to reflect the
renaming/restructuring.

## 0.1.5
- Fleshed out the README a bit (more is obviously needed).
- Start of creating the course group users service.

## 0.1.4+1
- Minor refactoring across code base to clean it up.
- Work on services.

## 0.1.4
- Start of working on course groups.
- Refactoring of service structures.

## 0.1.3+1
- More work on course grade attempts.

## 0.1.3
- Start of adding course grade attempts.

## 0.1.2+3
- Minor refactoring.

## 0.1.2+2
- Minor refactoring.

## 0.1.2+1
- Added check for empty response for retrieving an access token.

## 0.1.2
- Removing custom errors.
  - While seems like a good idea, they have issues for users of the library.

## 0.1.1+2
- Resolved a host issue.

## 0.1.1+1
- Resolved new authorization scheme issues.

## 0.1.1
- Reworking authorization scheme.

## 0.1.0+2
- Minor refactoring.

## 0.1.0+1
- All tests pass.

## 0.1.0
- Reached feature parity with the Dart version of blackboard.rest.

## 0.0.11+1
- More work on CourseGrades and System code.

## 0.0.11
- Refactored endpoints to simplify their use, refactoring as necessary.
- More work on Terms code.

## 0.0.10
- Start of CourseGrades code.
- Start of Announcements code.
- Start of Contents code.
- Start of CourseGroups code.
- Start of Roles code.
- Start of Terms code.
- Refactored to greatly simplify error types.

## 0.0.9
- Start of System code.

## 0.0.8
- Start of CourseMemberships code.

## 0.0.7+1
- More work on Courses.

## 0.0.7
- Start of Courses code.

## 0.0.6
- Start of creating custom errors.
- More work on OAuth2 and Users.
- More testing.

## 0.0.5
- Adding the connector to allow arbitrary REST services to connect to the API.

## 0.0.4
- Finished client-based authorization for OAuth2.
- Start of creating Users service.

## 0.0.3
- Moved source code to 'src' folder, and created a 'test' folder at same level.
- Start of adding tests.
- Worked on OAuth2 code.

## 0.0.2
- More work of OAuth2 code.
- Start of Users code.

## 0.0.1
- Initial commit.
- Start of porting the blackboard.rest library from Dart to Go.
  - Start of adding OAuth2 and endpoints code.
