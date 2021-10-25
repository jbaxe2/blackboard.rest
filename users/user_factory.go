package users

import (
  "time"
)

/**
 * The [NewUser] function...
 */
func NewUser (rawUser map[string]interface{}) User {
  var created, lastLogin time.Time

  if nil != rawUser["created"] {
    created, _ = time.Parse (time.RFC3339, rawUser["created"].(string))
  }

  if nil != rawUser["lastLogin"] {
    lastLogin, _ = time.Parse (time.RFC3339, rawUser["lastLogin"].(string))
  }

  if nil == rawUser["studentId"] {
    rawUser["studentId"] = ""
  }

  return User {
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
      _parseInstitutionRoles (rawUser["institutionRoleIds"]),
    SystemRoleIds: _parseSystemRoles (rawUser["systemRoleIds"]),
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
func _parseInstitutionRoles (rawInstitutionRoles ...interface{}) []string {
  var institutionRoles = make ([]string, len (rawInstitutionRoles))

  for _, rawRole := range rawInstitutionRoles {
    if institutionRole, haveInstRole := rawRole.(string); haveInstRole {
      institutionRoles = append (institutionRoles, institutionRole)
    }
  }

  return institutionRoles
}

/**
 * The [_parseSystemRoles] function...
 */
func _parseSystemRoles (rawSystemRoles ...interface{}) []SystemRole {
  var systemRoles = make ([]SystemRole, len (rawSystemRoles))

  for _, rawRole := range rawSystemRoles {
    if systemRole, haveSystemRole := rawRole.(string); haveSystemRole {
      systemRoles = append (systemRoles, SystemRole (systemRole))
    }
  }

  return systemRoles
}

/**
 * The [_parseUserAvailability] function...
 */
func _parseUserAvailability (rawUserAvailability string) UserAvailability {
  return UserAvailability (rawUserAvailability)
}

/**
 * The [_parseName] function...
 */
func _parseName (rawName map[string]interface{}) Name {
  name := Name {
    Given: rawName["given"].(string),
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
