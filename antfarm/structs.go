package antfarm

type room struct {
	Name                string // Name
	X, Y                int    // Coordinates
	Paths               []*room
	ParentIn, ParentOut *room // Store parents for new path
	VisitIn, VisitOut   bool  // Flag for checking while traversing
	Next                *room
}

// List of Room nodes. Used to store found paths
type list struct {
	Len   int
	Front *room
	Back  *room
}

type antfarm struct {
	Ants       int
	Start, End string
	Rooms      map[string]*room
	StepsCount int
	State      int
	Result     *Result
}

type Result struct {
	Ants  int
	Paths []*list
}
