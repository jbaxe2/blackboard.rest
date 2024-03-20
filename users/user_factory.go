package users

import (
  "time"
)

/**
 * The [NewUsers] function creates a slice of User instances from a slice of raw
 * maps that typically were returned from a response from a REST API users call.
 */
func NewUsers (rawUsers []map[string]interface{}) []User {
  newUsers := make ([]User, len (rawUsers))

  for i, rawUser := range rawUsers {
    newUsers[i] = NewUser (rawUser)
  }

  return newUsers
}

/**
 * The [NewUser] function creates a new User instance from a raw map typically
 * returned from a response from a REST API users call.
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
      _parseInstitutionRoleIds (rawUser["institutionRoleIds"].([]interface{})),
    SystemRoleIds: _parseSystemRoles (rawUser["systemRoleIds"].([]interface{})),
    Availability:
      UserAvailability (
        rawUser["availability"].(map[string]interface{})["available"].(string),
      ),
    Name: _parseName (rawUser["name"].(map[string]interface{})),
  }
}

/**
 * The [_parseInstitutionRoles] function builds a slice of institution roles'
 * identifiers.
 */
func _parseInstitutionRoleIds (rawRoleIds []interface{}) []string {
  roleIds := make ([]string, 0)

  for _, rawRoleId := range rawRoleIds {
    roleIds = append (roleIds, rawRoleId.(string))
  }

  return roleIds
}

/**
 * The [_parseSystemRoles] function determines if there are any valid system
 * roles and builds a slice of their corresponding instances.
 */
func _parseSystemRoles (rawSystemRoles []interface{}) []SystemRole {
  var systemRoles = make ([]SystemRole, 0)

  for _, rawRole := range rawSystemRoles {
    systemRoles = append (systemRoles, SystemRole (rawRole.(string)))
  }

  return systemRoles
}

/**
 * The [_parseName] function parses the components of a user's name.
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
