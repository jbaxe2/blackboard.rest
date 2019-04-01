package users

/**
 * The [UserAvailability] type...
 */
type UserAvailability string

const (
  Yes       UserAvailability = "Yes"
  No        UserAvailability = "No"
  Disabled  UserAvailability = "Disabled"
)
