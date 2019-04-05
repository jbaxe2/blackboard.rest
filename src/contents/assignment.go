package contents

import "github.com/jbaxe2/blackboard.rest.go/src/course_grades"

/**
 * The [Assignment] type...
 */
type Assignment struct {
  Id, ParentId, Title, Instructions, Description string

  Position int

  FileUploadIds []string

  Availability ContentAvailability

  Grading course_grades.Grading

  Score course_grades.Scoring
}

/**
 * The [CreateAssignmentResult] type...
 */
type CreateAssignmentResult struct {
  ContentId, GradeColumnId, AssessmentId string

  QuestionIds []string
}
