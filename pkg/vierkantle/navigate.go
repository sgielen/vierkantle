package vierkantle

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Path []Coord

func (b *Board) candidatesForInitialNavigation() []Coord {
	var res []Coord
	for x := 0; x < b.Width; x++ {
		for y := 0; y < b.Height; y++ {
			res = append(res, Coord{X: x, Y: y})
		}
	}
	return res
}

func (b *Board) candidatesToNavigateFrom(c Coord) []Coord {
	var res []Coord
	if c.X > 0 {
		if c.Y > 0 {
			res = append(res, Coord{X: c.X - 1, Y: c.Y - 1})
		}
		res = append(res, Coord{X: c.X - 1, Y: c.Y})
		if c.Y+1 < b.Height {
			res = append(res, Coord{X: c.X - 1, Y: c.Y + 1})
		}
	}
	if c.Y > 0 {
		res = append(res, Coord{X: c.X, Y: c.Y - 1})
	}
	res = append(res, Coord{X: c.X, Y: c.Y})
	if c.Y+1 < b.Height {
		res = append(res, Coord{X: c.X, Y: c.Y + 1})
	}
	if c.X+1 < b.Width {
		if c.Y > 0 {
			res = append(res, Coord{X: c.X + 1, Y: c.Y - 1})
		}
		res = append(res, Coord{X: c.X + 1, Y: c.Y})
		if c.Y+1 < b.Height {
			res = append(res, Coord{X: c.X + 1, Y: c.Y + 1})
		}
	}
	return res
}

func (b *Board) NextNavigationsFrom(path Path) []Path {
	var res []Path
	if len(path) == 0 {
		for _, candidate := range b.candidatesForInitialNavigation() {
			res = append(res, Path{candidate})
		}
		return res
	}

	lastCoord := path[len(path)-1]
	candidates := b.candidatesToNavigateFrom(lastCoord)
	for _, candidate := range candidates {
		candidateInPath := false
		for _, coord := range path {
			if candidate.X == coord.X && candidate.Y == coord.Y {
				candidateInPath = true
				break
			}
		}
		if !candidateInPath {
			newPath := make(Path, len(path)+1)
			copy(newPath, path)
			newPath[len(path)] = candidate
			res = append(res, newPath)
		}
	}

	return res
}
