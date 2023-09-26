package lemin

import (
	lemin "lemin/functions"
	"testing"
)

func TestAllocateAnts(t *testing.T) {
    // Create a sample Allpaths struct with paths and ants
    room1 := &lemin.Room{}
    room2 := &lemin.Room{}
    path1 := &lemin.Path{P: []*lemin.Room{room1, room2}}
    path2 := &lemin.Path{P: []*lemin.Room{room1, room2}}
    optimalPath := []*lemin.Path{path1, path2}
    pathsInfo := &lemin.Allpaths{
        OptPath: optimalPath,
        MinTravelTime: 5,
    }
    

    // pathsInfo := &lemin.Allpaths{
    //     Paths: []*lemin.Path{
    //         &lemin.Path{P: []*lemin.Room{}},
    //         &lemin.Path{P: []*lemin.Room{}},
    //         &lemin.Path{P: []*lemin.Room{}},
    //     },
    // }
    // totalAnts := 10
    ants := 10

    // Test the AllocateAnts function
    antAllocations := lemin.AllocateAnts(pathsInfo, ants)
    expectedPath1Allocation := 6
    expectedPath2Allocation := 4

    if antAllocations[0] != expectedPath1Allocation{
        t.Errorf("expected allocation for path1 %d, but got %d", expectedPath1Allocation, antAllocations[0])
    }
    if antAllocations[1] != expectedPath2Allocation{
        t.Errorf("expected allocation for path2 %d, but got %d", expectedPath2Allocation, antAllocations[1])
    }
    // Check if the allocations are correct
    // expectedAllocations := []int{3, 3, 4} // Expected allocations for the given paths and ants
    // for i, allocation := range antAllocations {
    //     if allocation != expectedAllocations[i] {
    //         t.Errorf("Expected allocation of %d ants for path %d, but got %d", expectedAllocations[i], i, allocation)
    //     }
    // }

    // Test case with no ants to allocate
    // emptyPathsInfo := &lemin.Allpaths{}
    // emptyAnts := 0
    // emptyAllocations := lemin.AllocateAnts(emptyPathsInfo, emptyAnts)
    // if len(emptyAllocations) != 0 {
    //     t.Errorf("Expected empty allocations, but got %v", emptyAllocations)
    // }
}
