package factory

import (
  "github.com/jbaxe2/blackboard.rest.go/src/users"
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
  created, _ := time.Parse (time.RFC3339, rawUser["created"].(string))
  lastLogin, _ := time.Parse (time.RFC3339, rawUser["lastLogin"].(string))

  return users.User {
    rawUser["id"].(string), rawUser["uuid"].(string),
    rawUser["externalId"].(string), rawUser["dataSourceId"].(string),
    rawUser["userName"].(string), rawUser["email"].(string),
    rawUser["studentId"].(string), created,
    lastLogin, rawUser["institutionRoleIds"].([]string),
    _parseSystemRoles (rawUser["systemRoleIds"].([]string)),
    _parseUserAvailability (rawUser["availability"].(string)),
    _parseName (rawUser["name"].(map[string]string)),
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
    rawName["given"], rawName["family"], rawName["middle"],
    rawName["other"], rawName["suffix"], rawName["title"],
  }
}
