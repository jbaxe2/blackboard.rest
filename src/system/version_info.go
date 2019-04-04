package system

/**
 * The [VersionInfo] type...
 */
type VersionInfo struct {
  learn Version
}

/**
 * The [Version] type...
 */
type Version struct {
  major, minor, patch int

  build string
}
