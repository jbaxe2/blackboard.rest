package blackboard_rest

import (
  "net/url"

  "github.com/jbaxe2/blackboard.rest/_scaffolding/services"
  "github.com/jbaxe2/blackboard.rest/oauth2"
)

/**
 * The [_BlackboardRest] type...
 */
type _BlackboardRest struct {
  host url.URL

  accessToken oauth2.AccessToken

  service services.BlackboardRestService
}
