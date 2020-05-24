package orbit

import (
	"strings"
)

type orbitTable struct {
	childToParent map[string]string
}



func buildOrbitTables(directOrbits []string) orbitTable {
	
	ot := orbitTable{
		childToParent: make(map[string]string),
	}

	for _, orbit := range directOrbits {
		pair := strings.Split(orbit,")")
		
		_, ok := ot.childToParent[pair[1]]
		if !ok {
			ot.childToParent[pair[1]] = pair[0]
		}

	}

	return ot	
}


func (o orbitTable) numOrbits() int {
	orbits := 0
	for k  := range o.childToParent {
		orbits += o.getHops(k,0)
	}
	return orbits
}

func (o orbitTable) getHops(body string, depth int) int {
	parent, ok := o.childToParent[body]
	if ok {
		depth = o.getHops(parent, depth+1)
	}
	return depth
}

func (o orbitTable) getRoute(body string, route []string) []string {
	parent, ok := o.childToParent[body]
	if ok {
		route = append(o.getRoute(parent,route), parent)
	}
	return route
}

func (o orbitTable) getPath(origin string, dest string) int {
	
	r1 := o.getRoute(origin, []string{})
	r2 := o.getRoute(dest, []string{})
	

	i := 0
	for i = 0; r1[i] == r2[i]; i++ {
	}

	// Route from origin is reverse path from origin to intersection
	route := reverseSlice(r1[i-1:]) 
	// and then append the route from the intersection to the destination
	route = append(route, r2[i:]...)
		
	return len(route)-1
}

func reverseSlice(a []string) []string {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}