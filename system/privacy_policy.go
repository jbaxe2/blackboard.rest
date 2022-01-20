package system

import "net/url"

/**
 * The [PrivacyPolicy] type...
 */
type PrivacyPolicy struct {
  Blackboard, Institution *url.URL
}
