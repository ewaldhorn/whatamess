package components

type HomeComponent struct{}

var Home HomeComponent

func (h HomeComponent) Render() string {
	return "<H1>Home Component</H1>"
}
