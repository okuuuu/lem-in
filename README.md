<p style='text-align: justify;'>
  
## Lem-in

### Description

This project is meant to write a simplified digital version of an [ant farm](https://en.wikipedia.org/wiki/Ant_colony_optimization_algorithms). The program reads data describing the ants and the colony from a file passed in the arguments. The available paths are being found in the order of increasing length using BFS algorithm. Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

**The current project is 80% complete.**

File <code>antfarm/antfarm.go</code> contains all functions used for initializing and populating an antfarm using file contents.

File <code>antfarm/bfs.go</code> contains the push-pull functions for linked list which have been used to make queues for FIFO method. The file contains the BFS algorithm itself.

Files <code>antfarm/solve.go</code> and <code>antfarm/structs.go</code> contain auxilary functions to get and process the final result as well as the structures used throughout the program.

### Allowed Packages

- Only the [standard go](https://golang.org/pkg/) packages are allowed.

### Audit Details

- Here you can see [audit details](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit).

### Usage

- To run the code type in your terminal:
```
go run main.go
```

### Credits

- fpetuhov (okuu)

</p>