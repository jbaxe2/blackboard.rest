package blackboard_rest

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [_BlackboardRest] type provides some internal components needed by all of
 * the services implemented by this library.
 */
type _BlackboardRest struct {
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService
}
