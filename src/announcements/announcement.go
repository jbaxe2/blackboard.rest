package announcements

import "time"

/**
 * The [Announcement] type...
 */
type Announcement struct {
  Id, Title, Body string

  Availability Availability

  ShowAtLogin, ShowInCourses bool

  Created, Modified time.Time
}

/**
 * The [Availability] type...
 */
type Availability struct {
  Duration AnnouncementDuration
}

/**
 * The [AnnouncementDuration] type...
 */
type AnnouncementDuration struct {
  Type string

  Start, End time.Time
}
