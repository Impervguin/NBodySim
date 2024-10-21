package object

type InvisibleObject struct{}

func (i *InvisibleObject) IsVisible() bool {
	return false
}
