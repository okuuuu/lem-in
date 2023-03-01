package antfarm

import "errors"

func findPath(farm *antfarm) bool {
	queue := &list{}
	start := farm.Rooms[farm.Start]
	end := farm.Rooms[farm.End]
	start.Visit = true
	queue.Enqueue(start)
	for queue.Front != nil {
		room := queue.Dequeue()
		if room == end {
			path := &list{
				Len: -1,
			}
			for room != start {
				if room != end {
					room.Blocked = true
				}
				path.Enqueue(room)
				path.Len++
				room = room.Parent
			}
			path.Enqueue(room)
			farm.Paths = append(farm.Paths, path)
			Reset(farm)
			return true
		}
		for _, next := range room.Paths {
			if !next.Visit && !next.Blocked {
				next.Visit = true
				next.Parent = room
				queue.Enqueue(next)
			}
		}
	}
	return false
}

func checkPaths(farm *antfarm) bool {
	return true
}

func (q *list) Enqueue(r *room) {
	node := &node{
		Room: r,
	}
	if q.Front == nil {
		q.Front = node
		q.Back = node
		return
	}
	q.Back.Next = node
	q.Back = node
}

func (q *list) Dequeue() *room {
	if q.Front == nil {
		return nil
	}
	res := q.Front
	if q.Front == q.Back {
		q.Front = nil
		q.Back = nil
	} else {
		q.Front = q.Front.Next
	}
	return res.Room
}

func Reset(farm *antfarm) {
	for _, room := range farm.Rooms {
		room.Parent = nil
		room.Visit = false
	}
}

func (a *antfarm) FindResult() error {
	for {
		if !findPath(a) {
			// path not found, then check for prev path count
			if a.StepsCount > 0 {
				return nil
			}
			return errors.New("path not found")
		}
		if !checkPaths(a) {
			return nil
		}
	}
}
