package utils

/**
 * The [StringInStrings] function determines whether a particular string is
 * contained within a slice of strings.
 */
func StringInStrings (search string, stringsArr []string) bool {
  for _, stringItem := range stringsArr {
    if stringItem == search {
      return true
    }
  }

  return false
}

/**
 * The [NormalizeRawResponse] function takes an interface slice and normalizes it
 * to a slice of maps with string-based keys.
 */
func NormalizeRawResponse (rawResponse []interface{}) []map[string]interface{} {
  mappedResponse := make ([]map[string]interface{}, 0)

  for _, rawColumn := range rawResponse {
    normalizedColumn, isNormalized := rawColumn.(map[string]interface{})

    if isNormalized {
      mappedResponse = append (mappedResponse, normalizedColumn)
    }
  }

  return mappedResponse
}
