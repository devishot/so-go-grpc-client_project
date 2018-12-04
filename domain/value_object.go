package domain

// Project:

type ProjectInputValue struct {
	Title       string
	Description string
	ClientID    ID
}

type ProjectConnectionPageValue struct {
	Projects    []ProjectEntity
	Total       int
	EndCursor   ConnCursor
	StartCursor ConnCursor
}

// Connection:

type ConnectionArgumentsValue struct {
	// forward
	First int32
	After ConnCursor
	// backward
	Last   int32
	Before ConnCursor
}

type ConnectionPageInfoValue struct {
	// forward
	HasNextPage bool
	EndCursor   ConnCursor
	// backward
	HasPreviousPage bool
	StartCursor     ConnCursor
}

type ConnectionEdgeValue struct {
	Cursor ConnCursor
	Node   interface{}
}

type ConnectionValue struct {
	Edges    []*ConnectionEdgeValue
	PageInfo ConnectionPageInfoValue
}
