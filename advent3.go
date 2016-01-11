package advent

// Advent03Houses takes a list of ^>v< and returns the number of positions visited
func Advent03Houses(s string) (sum1, sum2 int) {
	return NewRoute().Nav(s).HowManyUnique(), NewRoute().DualNav(s).HowManyUnique()
}

// RoutePos is used as keys in map Route.visited
type routePos struct {
	x, y int
}

// Route records Santa's X/Y movement
type Route struct {
	// List of positions visited so far
	visited map[routePos]struct{}
	// Current positions of Santa and RoboSanta
	posX1, posY1, posX2, posY2 int
}

// NewRoute creates a blank Route, with a visit to 0,0, ready to record movements
func NewRoute() (r *Route) {
	r = &Route{}
	r.visited = make(map[routePos]struct{})
	// don't move any, just get the side effect of recording the origin visit
	r.MoveSanta(0, 0)
	return r
}

// Nav moves Santa around according to ^>v< instructions and records visits
func (r *Route) Nav(s string) *Route {
	for _, m := range s {
		switch m {
		case '^':
			r.MoveSanta(0, 1)
		case '>':
			r.MoveSanta(1, 0)
		case 'v':
			r.MoveSanta(0, -1)
		case '<':
			r.MoveSanta(-1, 0)
		}
	}
	return r
}

// DualNav moves Santa and Robo-Santa around and records their visits
// according to alternating ^>v< instructions
func (r *Route) DualNav(s string) *Route {
	for i, m := range s {
		f := r.MoveSanta
		if i%2 != 0 {
			f = r.MoveRoboSanta
		}
		switch m {
		case '^':
			f(0, 1)
		case '>':
			f(1, 0)
		case 'v':
			f(0, -1)
		case '<':
			f(-1, 0)
		}
	}
	return r
}

// MoveSanta changes position in a delta of x/y direction and records the visit
func (r *Route) MoveSanta(dX, dY int) {
	r.posX1 += dX
	r.posY1 += dY
	r.visited[routePos{r.posX1, r.posY1}] = struct{}{}
}

// MoveRoboSanta changes position in a delta of x/y direction and records the visit
func (r *Route) MoveRoboSanta(dX, dY int) {
	r.posX2 += dX
	r.posY2 += dY
	r.visited[routePos{r.posX2, r.posY2}] = struct{}{}
}

// HowManyUnique returns the number of locations visited
func (r *Route) HowManyUnique() int {
	return len(r.visited)
}
