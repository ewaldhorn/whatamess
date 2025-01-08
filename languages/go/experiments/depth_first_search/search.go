package main

// ----------------------------------------------------------------------------
func (n *Node) DepthFirstSearch(searchFor string) *Node {
	stack := []*Node{n} // set stack to the current node

	for len(stack) > 0 {
		current := stack[len(stack)-1] // grab the top node
		stack = stack[:len(stack)-1]   // pop it off the stack

		if current.Name == searchFor {
			return current // found it, so return
		}

		// didn't find it, need to dive deeper into this branch
		// now make sure we check the children of the current node next
		for _, dependant := range current.Dependants {
			stack = append(stack, dependant)
		}
	}

	// found nothing
	return nil
}
