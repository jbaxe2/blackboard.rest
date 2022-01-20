package contents

import courseGrades "github.com/jbaxe2/blackboard.rest/course_grades"

/**
 * The [Assignment] type...
 */
type Assignment struct {
  Id, ParentId, Title, Instructions, Description string

  Position int

  FileUploadIds []string

  Availability Availability

  Grading courseGrades.Grading

  Score courseGrades.Scoring
}

/**
 * The [CreateAssignmentResult] type...
 */
type CreateAssignmentResult struct {
  ContentId, GradeColumnId, AssessmentId string

  QuestionIds []string
}
