package graphql_connection

type ConnectionArgumentsValue struct {
	// forward
	First int
	After Cursor
	// backward
	Last   int
	Before Cursor
}

func (args ConnectionArgumentsValue) IsForward() (forward bool, err error) {
	if args.First != 0 && args.After != "" {
		forward = true
	} else if args.Last != 0 && args.Before != "" {
		forward = false
	} else {
		err = IncorrectConnectionArgsError
	}
	return forward, err
}

type ConnectionPageInfoValue struct {
	// forward
	HasNextPage bool
	EndCursor   Cursor
	// backward
	HasPreviousPage bool
	StartCursor     Cursor
}

type ConnectionEdgeValue struct {
	Cursor Cursor
	Node   interface{}
}

type ConnectionValue struct {
	Edges    []*ConnectionEdgeValue
	PageInfo ConnectionPageInfoValue
}
