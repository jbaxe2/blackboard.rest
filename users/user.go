package users

import "time"

/**
 * The [User] type...
 */
type User struct {
  id, uuid, externalId, dataSourceId, userName, email, studentId string

  created, lastLogin time.Time

  institutionRoleIds []string

  systemRoleIds []SystemRole

  availability UserAvailability

  name Name
}

func (user *User) Id() string {
  return user.id
}

func (user *User) Uuid() string {
  return user.uuid
}

func (user *User) ExternalId() string {
  return user.externalId
}

func (user *User) DataSourceId() string {
  return user.dataSourceId
}

func (user *User) UserName() string {
  return user.userName
}

func (user *User) Email() string {
  return user.email
}

func (user *User) StudentId() string {
  return user.studentId
}

func (user *User) Created() time.Time {
  return user.created
}

func (user *User) LastLogin() time.Time {
  return user.lastLogin
}

func (user *User) InstitutionRoleIds() []string {
  return user.institutionRoleIds
}

func (user *User) SystemRoleIds() []SystemRole {
  return user.systemRoleIds
}

func (user *User) Availability() UserAvailability {
  return user.availability
}

func (user *User) Name() Name {
  return user.name
}
