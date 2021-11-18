package api

const Base = "/learn/api/public/v{v}/"

/**
 * The [OAuth2Endpoint] type provides defining typed OAuth2 endpoint constants.
 */
type OAuth2Endpoint string

const (
  AuthorizationCode OAuth2Endpoint = "oauth2/authorizationcode"
  RequestToken      OAuth2Endpoint = "oauth2/token"
)

/**
 * The [ContentEndpoint] type provides defining typed content endpoint constants.
 */
type ContentEndpoint string

const (
  Contents          ContentEndpoint = "courses/{courseId}/contents"
  CreateAssignment  ContentEndpoint = "courses/{courseId}/contents/createAssignment"
  Content           ContentEndpoint = "courses/{courseId}/contents/{contentId}"
  ContentChildren   ContentEndpoint = "courses/{courseId}/contents/{contentId}/children"
)

/**
 * The [CourseGradeAttemptsEndpoint] type provides defining typed course grade
 * attempts endpoint constants.
 */
type CourseGradeAttemptsEndpoint string

const (
  FileMetadataList  CourseGradeAttemptsEndpoint =
    "courses/{courseId}/gradebook/attempts/{attemptId}/files"
  AttachFile        CourseGradeAttemptsEndpoint =
    "courses/{courseId}/gradebook/attempts/{attemptId}/files"
  FileMetadata      CourseGradeAttemptsEndpoint =
    "courses/{courseId}/gradebook/attempts/{attemptId}/files/{attemptFileId}"
  Download          CourseGradeAttemptsEndpoint =
    "courses/{courseId}/gradebook/attempts/{attemptId}/files/{attemptFileId}/download"
)

/**
 * The [CourseGradesEndpoint] type provides defining typed course grades endpoint
 * constants.
 */
type CourseGradesEndpoint string

const (
  GradeColumns    CourseGradesEndpoint = "courses/{courseId}/gradebook/columns"
  GradeColumn     CourseGradesEndpoint =
    "courses/{courseId}/gradebook/columns/{columnId}"
  ColumnAttempts  CourseGradesEndpoint =
    "courses/{courseId}/gradebook/columns/{columnId}/attempts"
  ColumnAttempt   CourseGradesEndpoint =
    "courses/{courseId}/gradebook/columns/{columnId}/attempts/{attemptId}"
)

/**
 * The [CourseGroupUsersEndpoint] type provides defining typed course group users
 * endpoint constants.
 */
type CourseGroupUsersEndpoint string

const (
  GroupMemberships  CourseGroupUsersEndpoint =
    "courses/{courseId}/groups/{groupId}/users"
  GroupMembership   CourseGroupUsersEndpoint =
    "courses/{courseId}/groups/{groupId}/users/{userId}"
)

/**
 * The [CourseGroupsEndpoint] type provides defining typed course group endpoint
 * constants.
 */
type CourseGroupsEndpoint string

const (
  Groups          CourseGroupsEndpoint = "courses/{courseId}/groups"
  GroupSets       CourseGroupsEndpoint = "courses/{courseId}/groups/sets"
  GroupSet        CourseGroupsEndpoint = "courses/{courseId}/groups/sets/{groupId}"
  GroupSetGroups  CourseGroupsEndpoint =
    "courses/{courseId}/groups/sets/{groupId}/groups"
  Group           CourseGroupsEndpoint = "courses/{courseId}/groups/{groupId}"
)

/**
 * The [CourseMembershipsEndpoint] type provides defining typed course
 * memberships endpoint constants.
 */
type CourseMembershipsEndpoint string

const (
  CourseMemberships CourseMembershipsEndpoint = "courses/{courseId}/users?expand=user"
  UserMemberships   CourseMembershipsEndpoint = "users/{userId}/courses"
  Membership        CourseMembershipsEndpoint =
    "courses/{courseId}/users/{userId}?expand=user"
)

/**
 * The [CoursesEndpoint] type provides defining typed courses endpoint constants.
 */
type CoursesEndpoint string

const (
  Courses       CoursesEndpoint = "courses"
  Course        CoursesEndpoint = "courses/{courseId}"
  Children      CoursesEndpoint = "courses/{courseId}/children"
  Child         CoursesEndpoint = "courses/{courseId}/children/{childCourseId}"
  CrossListSet  CoursesEndpoint = "courses/{courseId}/crossListSet"
  Copy          CoursesEndpoint = "courses/{courseId}/copy"
)

/**
 * The [SystemEndpoint] type provides defining typed system endpoint constants.
 */
type SystemEndpoint string

const (
  Policies  SystemEndpoint = "system/policies/privacy"
  Version   SystemEndpoint = "system/version"
)

/**
 * The [TermsEndpoint] type provides defining typed terms endpoint constants.
 */
type TermsEndpoint string

const (
  Terms TermsEndpoint = "terms"
  Term  TermsEndpoint = "terms/{termId}"
)

/**
 * The [UsersEndpoint] type provides defining typed users endpoint constants.
 */
type UsersEndpoint string

const (
  Users UsersEndpoint = "users"
  User  UsersEndpoint = "users/{userId}"
)
