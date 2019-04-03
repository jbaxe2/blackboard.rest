package error

/**
 * The [RestableError] interface...
 */
type RestableError interface {
  error
}

/**
 * The [RestError] type...
 */
type RestError struct {
  status float64

  code, message, developerMessage, extraInfo string

  RestableError
}

func (err *RestError) Status() float64 {
  return err.status
}

func (err *RestError) Code() string {
  return err.code
}

func (err *RestError) Message() string {
  return err.message
}

func (err *RestError) DeveloperMessage() string {
  return err.developerMessage
}

func (err *RestError) ExtraInfo() string {
  return err.extraInfo
}

func (err *RestError) SetStatus (status float64) {
  err.status = status
}

func (err *RestError) SetCode (code string) {
  err.code = code
}

func (err *RestError) SetMessage (message string) {
  err.message = message
}

func (err *RestError) SetDeveloperMessage (developerMessage string) {
  err.developerMessage = developerMessage
}

func (err *RestError) SetExtraInfo (extraInfo string) {
  err.extraInfo = extraInfo
}

/**
 * The [Error] method...
 */
func (err RestError) Error() string {
  return "Error (" + err.Code() + "): " + err.Message()
}
