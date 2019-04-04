package config

const Base = "/learn/api/public/v1/"

const BaseV2 = "/learn/api/public/v2/"

/**
 * The [OAuth2Endpoints] function...
 */
func OAuth2Endpoints() map[string]string {
  endpoints := make (map[string]string)

  endpoints["authorization_code"] = "oauth2/authorizationcode"
  endpoints["request_token"] = "oauth2/token"

  return endpoints
}

/**
 * The [CourseMembershipsEndpoints] function...
 */
func CourseMembershipsEndpoints() map[string]string {
  endpoints := make (map[string]string)

  endpoints["course_memberships"] = "courses/{courseId}/users"
  endpoints["user_memberships"] = "users/{userId}/courses"
  endpoints["membership"] = "courses/{courseId}/users/{userId}"

  return endpoints
}

/**
 * The [CoursesEndpoints] function...
 */
func CoursesEndpoints() map[string]string {
  endpoints := make (map[string]string)

  endpoints["courses"] = "courses"
  endpoints["course"] = "courses/{courseId}"
  endpoints["children"] = "courses/{courseId}/children"
  endpoints["child"] = "courses/{courseId}/children/{childCourseId}"
  endpoints["crossListSet"] = "courses/{courseId}/crossListSet"
  endpoints["copy"] = "courses/{courseId}/copy"

  return endpoints
}

/**
 * The [UsersEndpoints] function...
 */
func UsersEndpoints() map[string]string {
  endpoints := make (map[string]string)

  endpoints["users"] = "users"
  endpoints["user"] = "users/{userId}"

  return endpoints
}
