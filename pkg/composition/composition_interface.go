package composition

type Composition interface {
	// Resize we have to scale the composition
	Resize(width, height int)
}
