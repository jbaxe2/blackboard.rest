package utils

/**
 * The [StringInStrings] function determines whether a particular string is
 * contained within an array of strings.
 */
func StringInStrings (search string, stringsArr []string) bool {
  for _, stringItem := range stringsArr {
    if stringItem == search {
      return true
    }
  }

  return false
}
