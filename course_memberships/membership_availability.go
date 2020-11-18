package course_memberships

/**
 * The [MembershipAvailability] type...
 */
type MembershipAvailability string

const (
  Yes      MembershipAvailability = "Yes"
  No       MembershipAvailability = "No"
  Disabled MembershipAvailability = "Disabled"
)
