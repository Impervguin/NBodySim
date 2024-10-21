package object

type VisibleObject struct{}

func (v *VisibleObject) IsVisible() bool {
	return true
}
