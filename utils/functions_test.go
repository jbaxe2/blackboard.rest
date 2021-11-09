package utils_test

import (
  "testing"

  "github.com/jbaxe2/blackboard.rest/utils"
)

/**
 * The [TestStringInStrings] function...
 */
func TestStringInStrings (t *testing.T) {
  println ("A string can be found in an array of strings.")

  search := "some string"
  stringsArr := []string {"a string", "second string", search, "other string"}

  if !utils.StringInStrings (search, stringsArr) {
    t.Error ("A string should be found to be in a slice of strings.")
  }
}

/**
 * The [TestStringNotInStrings] function...
 */
func TestStringNotInStrings (t *testing.T) {
  println ("A string not in an array of strings will be confirmed as such.")

  stringsArr := []string {"a string", "second string", "other string"}

  if utils.StringInStrings ("string not in strings", stringsArr) {
    t.Error ("A string not in strings should not be found to be in as such.")
  }
}

/**
 * The [TestStringCanSearchAgainstEmptyStrings] function...
 */
func TestStringCanSearchAgainstEmptyStrings (t *testing.T) {
  println ("A string can be searched against an empty slice of strings.")

  if utils.StringInStrings ("a string", make ([]string, 0)) {
    t.Error ("Empty slice of strings should not result in finding search string.")
  }
}

/**
 * The [TestStringCanSearchAgainstNilSlice] function...
 */
func TestStringCanSearchAgainstNilSlice (t *testing.T) {
  println ("Searching for string against nil slice results in string not found.")

  if utils.StringInStrings ("some string", nil) {
    t.Error ("Nil slice of strings should not result in finding search string.")
  }
}

/**
 * The [TestNormalizeRawResponse] function...
 */
func TestNormalizeRawResponse (t *testing.T) {
  println ("A slice of interfaces normalizes to a slice of maps with string-based keys.")

  checkedValue := "some value"

  rawResponse := []interface{} {
    map[string]interface{} {
      "key1a": checkedValue,
      "key2a": 300,
    },
    map[string]interface{} {
      "key1b": "some other value",
      "key2b": true,
    },
  }

  normalizedResponse := utils.NormalizeRawResponse (rawResponse)
  normalizedValue, _ := normalizedResponse[0]["key1a"]

  if !(len (normalizedResponse) == len (rawResponse) &&
       normalizedValue == checkedValue) {
    t.Error ("Raw response should normalize into proper type.")
  }
}
