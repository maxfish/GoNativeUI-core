package utils

type Insets struct {
	Top, Right, Bottom, Left int
}

func HomogeneousInsets(inset int) Insets {
	return Insets{Top: inset, Right: inset, Bottom: inset, Left: inset}
}
