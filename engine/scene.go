package engine

type Drawable interface {
    Draw()
}

var _ Drawable = &Composite{}
var _ Drawable = &Scene{}

type Composite struct {
    Children []Drawable
}

func (c *Composite) Draw() {
    for _, child := range(c.Children) {
        child.Draw()
    }
}

type Scene struct {
    root *Composite
}

func (s *Scene) Draw() {
    s.root.Draw()
}
