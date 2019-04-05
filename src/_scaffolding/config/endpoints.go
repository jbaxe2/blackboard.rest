package config

const Base = "/learn/api/public/v1/"

const BaseV2 = "/learn/api/public/v2/"

/**
 * The [OAuth2Endpoints] map...
 */
var OAuth2Endpoints = map[string]string {
  "authorization_code": "oauth2/authorizationcode",
  "request_token": "oauth2/token",
}

/**
 * The [CourseMembershipsEndpoints] map...
 */
var CourseMembershipsEndpoints = map[string]string {
  "course_memberships": "courses/{courseId}/users",
  "user_memberships": "users/{userId}/courses",
  "membership": "courses/{courseId}/users/{userId}",
}

/**
 * The [CoursesEndpoints] map...
 */
var CoursesEndpoints = map[string]string {
  "courses": "courses",
  "course": "courses/{courseId}",
  "children": "courses/{courseId}/children",
  "child": "courses/{courseId}/children/{childCourseId}",
  "crossListSet": "courses/{courseId}/crossListSet",
  "copy": "courses/{courseId}/copy",
}

/**
 * The [SystemEndpoints] map...
 */
var SystemEndpoints = map[string]string {
  "policies":"system/policies/privacy",
  "version": "system/version",
}

/**
 * The [TermsEndpoints] map...
 */
var TermsEndpoints = map[string]string {
  "terms": "terms",
  "term": "terms/{termId}",
}

/**
 * The [UsersEndpoints] map...
 */
var UsersEndpoints = map[string]string {
  "users": "users",
  "user": "users/{userId}",
}
