package components

const MAX_Z = 100

/** Apply childFunc to all Component children with z index as priority **/
func ChildFunc(c Component, childFunc func(*Component)) {

	children := c.GetChildren()

	countRendered := 0
	var z int32 = 0
	for z = range MAX_Z {

		if len(children) == countRendered {
			return
		}

		for _, child := range children {
			if (*child).GetZ() == z {
				childFunc(child)
			}
		}
	}
}

/** Remove child from list of children on parent **/
func RemoveChild(child *Component) {

	parent := (*(*child).GetParent())
	temp := parent.GetChildren()

	for i, c := range parent.GetChildren() {

		if c != child {
			continue
		}

		temp = append(temp[:i], temp[i+1:]...)
		break
	}

	parent.SetChildren(temp)
}
