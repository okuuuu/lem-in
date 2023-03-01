package antfarm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func PopulateFarm() *antfarm {
	farm := &antfarm{}
	farm.Rooms = make(map[string]*room)
	farm.State = 0
	farm.Paths = []*list{}
	return farm
}

func (a *antfarm) ReadLine(line string) error {

	// Check for comment fields
	if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
		return nil
	}
	// Ants fields
	if a.State == 0 {
		if line == "##start" {
			a.State = 1
			return nil
		}
		err := a.AddAnts(line)
		if err != nil {
			return err
		}
		// Rooms fields
	} else if a.State == 1 {
		if line == "##end" {
			a.State = 2
			return nil
		}
		err := a.AddRoom(line, a.State)
		if err != nil {
			return err
		}
		// End room field
	} else if a.State == 2 {
		err := a.AddRoom(line, a.State)
		if err != nil {
			return err
		}
		a.State = 3
		return nil
		// Tunnels fields
	} else if a.State == 3 {
		err := a.AddPath(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *antfarm) AddAnts(line string) error {
	ants, err := strconv.Atoi(line)
	if err != nil || ants < 1 {
		return errors.New("invalid number of ants")
	}
	if a.Ants != 0 {
		return errors.New("ants have already been recorded")
	}
	a.Ants = ants
	return nil
}

func (a *antfarm) AddRoom(line string, state int) error {

	// Split line and check syntax
	strsplit := strings.Split(line, " ")
	if len(strsplit) != 3 || len(strsplit[0]) < 1 {
		return errors.New("invalid format of room")
	} else if strings.HasPrefix(strsplit[0], "L") {
		return errors.New("room cannot start with 'L'")
	} else if _, ok := a.Rooms[strsplit[0]]; ok {
		return fmt.Errorf("room name duplicated: '%v'", strsplit[0])
	}

	// Get room params
	name := strsplit[0]
	x, errX := strconv.Atoi(strsplit[1])
	y, errY := strconv.Atoi(strsplit[2])
	if errX != nil || errY != nil {
		return errors.New("invalid room coordinates")
	}
	room := &room{
		Name: name,
		X:    x,
		Y:    y,
	}
	a.Rooms[name] = room

	// Mark Start and End rooms
	if state == 1 && len(a.Rooms) == 1 {
		a.Start = name
	} else if state == 2 {
		a.End = name
	}

	return nil
}

func (a *antfarm) AddPath(line string) error {
	strsplit := strings.Split(line, "-")
	if len(strsplit) != 2 {
		return errors.New("invalid format of path")
	}
	name1, name2 := strsplit[0], strsplit[1]
	if name1 == name2 {
		return fmt.Errorf("rooms bound to each other. Line: '%v'", line)
	}
	room1 := a.Rooms[name1]
	room2 := a.Rooms[name2]
	room1.Paths = append(room1.Paths, room2)
	room2.Paths = append(room2.Paths, room1)
	return nil
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
