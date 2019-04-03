package factory

import (
  "github.com/jbaxe2/blackboard.rest.go/src/users"
  "reflect"
  "time"
)

/**
 * The [UserFactory] type...
 */
type UserFactory struct {}

/**
 * The [Create] method...
 */
func (userFactory *UserFactory) NewUser (rawUser map[string]interface{}) users.User {
  for k, v := range rawUser {
    print ("" + k + ": " + reflect.TypeOf (v).String())
  }

  var created, lastLogin time.Time

  if nil != rawUser["created"] {
    created, _ = time.Parse (time.RFC3339, rawUser["created"].(string))
  }

  if nil != rawUser["lastLogin"] {
    lastLogin, _ = time.Parse (time.RFC3339, rawUser["lastLogin"].(string))
  }

  return users.User {
    Id: rawUser["id"].(string),
    Uuid: rawUser["uuid"].(string),
    ExternalId: rawUser["externalId"].(string),
    DataSourceId: rawUser["dataSourceId"].(string),
    UserName: rawUser["userName"].(string),
    Email: rawUser["email"].(string),
    StudentId: rawUser["studentId"].(string),
    Created: created,
    LastLogin: lastLogin,
    InstitutionRoleIds: rawUser["institutionRoleIds"].([]string),
    SystemRoleIds: _parseSystemRoles (rawUser["systemRoleIds"].([]string)),
    Availability: _parseUserAvailability (rawUser["availability"].(string)),
    Name: _parseName (rawUser["name"].(map[string]string)),
  }
}

/**
 * The [_parseSystemRoles] function...
 */
func _parseSystemRoles (rawSystemRoles []string) []users.SystemRole {
  var systemRoles = make ([]users.SystemRole, len(rawSystemRoles))

  for key, role := range rawSystemRoles {
    systemRoles[key] = users.SystemRole (role)
  }

  return systemRoles
}

/**
 * The [_parseUserAvailability] function...
 */
func _parseUserAvailability (rawUserAvailability string) users.UserAvailability {
  return users.UserAvailability (rawUserAvailability)
}

/**
 * The [_parseName] function...
 */
func _parseName (rawName map[string]string) users.Name {
  return users.Name {
    Given: rawName["given"],
    Family: rawName["family"],
    Middle: rawName["middle"],
    Other: rawName["other"],
    Suffix: rawName["suffix"],
    Title: rawName["title"],
  }
}
