package users

/**
 * The [Name] type...
 */
type Name struct {
  given, family, middle, other, suffix, title string
}

func (name *Name) Given() string {
  return name.given
}

func (name *Name) Family() string {
  return name.family
}

func (name *Name) Middle() string {
  return name.middle
}

func (name *Name) Other() string {
  return name.other
}

func (name *Name) Suffix() string {
  return name.suffix
}

func (name *Name) Title() string {
  return name.title
}
