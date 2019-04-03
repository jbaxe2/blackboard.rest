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
    Email: rawUser["contact"].(map[string]interface{})["email"].(string),
    StudentId: rawUser["studentId"].(string),
    Created: created,
    LastLogin: lastLogin,
    InstitutionRoleIds:
      _parseInstitutionRoles (rawUser["institutionRoleIds"].([]interface{})),
    SystemRoleIds: _parseSystemRoles (rawUser["systemRoleIds"].([]interface{})),
    Availability:
      _parseUserAvailability (
        rawUser["availability"].(map[string]interface{})["available"].(string),
      ),
    Name: _parseName (rawUser["name"].(map[string]interface{})),
  }
}

/**
 * The [_parseInstitutionRoles] function...
 */
func _parseInstitutionRoles (rawInstitutionRoles []interface{}) []string {
  var institutionRoles = make ([]string, len (rawInstitutionRoles))

  for i, role := range rawInstitutionRoles {
    institutionRoles[i] = role.(string)
  }

  return institutionRoles
}

/**
 * The [_parseSystemRoles] function...
 */
func _parseSystemRoles (rawSystemRoles []interface{}) []users.SystemRole {
  var systemRoles = make ([]users.SystemRole, len(rawSystemRoles))

  for i, role := range rawSystemRoles {
    systemRoles[i] = users.SystemRole (role.(string))
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
func _parseName (rawName map[string]interface{}) users.Name {
  name := users.Name {
    Given:  rawName["given"].(string),
    Family: rawName["family"].(string),
  }

  if nil != rawName["middle"] {
    name.Middle = rawName["middle"].(string)
  }

  if nil != rawName["other"] {
    name.Other = rawName["other"].(string)
  }

  if nil != rawName["suffix"] {
    name.Suffix = rawName["suffix"].(string)
  }

  if nil != rawName["title"] {
    name.Title = rawName["title"].(string)
  }

  return name
}
