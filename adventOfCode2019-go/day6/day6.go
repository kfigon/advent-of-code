package day6

import "strings"

type graphNode struct {
	parent   string
	children []string
}

func (g *graphNode) addChild(children string) {
	if g.children == nil {
		g.children = make([]string, 0)
	}
	g.children = append(g.children, children)
}

type orbitGraph struct {
	m map[string]*graphNode
}

func buildOrbits(input string, lineSep string) *orbitGraph {
	o := &orbitGraph{
		m: map[string]*graphNode{},
	}

	vals := strings.Split(input, lineSep)
	for _, v := range vals {
		splitted := strings.Split(v, ")")
		node := splitted[0]
		orbittingObject := splitted[1]

		o.addNode(node, orbittingObject)
		o.addParent(orbittingObject, node)
	}

	return o
}

func (o *orbitGraph) addNode(node string, children string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.addChild(children)
		o.m[node] = g
	} else {
		val.addChild(children)
		o.m[node] = val
	}
}

func (o *orbitGraph) addParent(node string, parent string) {
	val, ok := o.m[node]
	if !ok {
		g := &graphNode{}
		g.parent = parent
		o.m[node] = g
	} else {
		val.parent = parent
		o.m[node] = val
	}
}

func (o *orbitGraph) calcOrbits(startingNode string) int {
	orbit, ok := o.m[startingNode]
	if !ok {
		return 0
	} else if orbit.parent == "" {
		return 0
	}

	return 1 + o.calcOrbits(orbit.parent)
}

func (o *orbitGraph) calcAllOrbits() int {
	visitedNodes := make(map[string]struct{})

	sum := 0
	for key := range o.m {
		if _, visited := visitedNodes[key]; visited {
			continue
		}
		visitedNodes[key] = struct{}{}

		sum += o.calcOrbits(key)
	}
	return sum
}

func (o *orbitGraph) findPath(start, end string) int {
	startingNode := o.m[start].parent
	targetNode := o.m[end].parent

	buildHistory := func(node string) map[string]int {
		parentsOfStartingNode := map[string]int{}
		steps := 0
		for node != "" {
			parentsOfStartingNode[node] = steps
			steps++
			node = o.m[node].parent
		}
		return parentsOfStartingNode
	}
	fromStart := buildHistory(startingNode)
	fromEnd := buildHistory(targetNode)

	var minNumberOfSteps *int
	for key, stepsFromStart := range fromStart {
		stepsFromEnd, ok := fromEnd[key]
		if !ok {
			continue
		}

		steps := stepsFromStart + stepsFromEnd
		if minNumberOfSteps == nil || steps < *minNumberOfSteps {
			minNumberOfSteps = &steps
		}
	}
	if minNumberOfSteps == nil {
		return -1
	}
	return *minNumberOfSteps
}