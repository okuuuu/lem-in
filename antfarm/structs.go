package antfarm

type room struct {
	Name    string // Name
	X, Y    int    // Coordinates
	Paths   []*room
	Parent  *room // Store parents for new path
	Visit   bool  // Flag for checking while traversing
	Blocked bool
}

// List of Room nodes. Used to store found paths
type list struct {
	Len   int
	Front *node
	Back  *node
}

type node struct {
	Room *room
	Next *node
}

type antfarm struct {
	Ants       int
	Start, End string
	Rooms      map[string]*room
	StepsCount int
	State      int
	Paths      []*list
}
