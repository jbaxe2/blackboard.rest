package _scaffolding

/**
 * The [NormalizeRawResponse] function...
 */
func NormalizeRawResponse (rawResponse []interface{}) []map[string]interface{} {
  mappedResponse := make ([]map[string]interface{}, len (rawResponse))

  for i, rawColumn := range rawResponse {
    mappedResponse[i] = rawColumn.(map[string]interface{})
  }

  return mappedResponse
}
