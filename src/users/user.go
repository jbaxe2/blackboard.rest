package users

import (
  "time"
)

/**
 * The [User] type...
 */
type User struct {
  Id, Uuid, ExternalId, DataSourceId, UserName, Email, StudentId string

  Created, LastLogin time.Time

  InstitutionRoleIds []string

  SystemRoleIds []SystemRole

  Availability UserAvailability

  Name Name
}
