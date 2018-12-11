package connection

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/devishot/so-go-grpc-client_project/domain"
	conn "github.com/devishot/so-go-grpc-client_project/infrastructure/graphql_connection"
	"github.com/devishot/so-go-grpc-client_project/interface/graphql"
)

func TestProjectConnectionService_Connection(t *testing.T) {
	var cID domain.ID = "123-123"

	projects := []domain.ProjectEntity{
		domain.NewProject(cID, "A", "aaa"),
		domain.NewProject(cID, "B", "bbb"),
		domain.NewProject(cID, "C", "ccc"),
	}
	for i, p := range projects {
		p.Timestamp.Add(time.Hour * 24 * time.Duration(i+1))
	}

	afterProject := projects[0]
	lastProject := projects[len(projects)-1]

	pageArgsExp := graphql.ProjectRepositoryPageArgs{
		First: 10,
		After: afterProject.Timestamp,
	}
	connArgs := conn.ConnectionArgumentsValue{
		First: pageArgsExp.First,
		After: encodeTimestampCursor(pageArgsExp.After),
	}
	connValsExp := &conn.ConnectionValue{
		Edges: []*conn.ConnectionEdgeValue{
			{
				Cursor: encodeTimestampCursor(projects[1].Timestamp),
				Node:   projects[1],
			},
			{
				Cursor: encodeTimestampCursor(projects[2].Timestamp),
				Node:   projects[2],
			},
		},
		PageInfo: conn.ConnectionPageInfoValue{
			HasNextPage: false,
			EndCursor:   encodeTimestampCursor(lastProject.Timestamp),
		},
	}

	connRepoMock := &ProjectConnectionRepositoryMock{
		GetLastByClientFunc: func(clientID domain.ID) (entity domain.ProjectEntity, e error) {
			assert.Equal(t, cID, clientID)
			return lastProject, nil
		},
		PaginateByTimestampFunc: func(clientID domain.ID, args graphql.ProjectRepositoryPageArgs) (
			entities []domain.ProjectEntity, e error) {
			assert.Equal(t, cID, clientID)
			return projects[1:], nil
		},
	}
	s := ProjectConnectionService{
		ConnRepo: connRepoMock,
	}

	val, err := s.Connection(cID, connArgs)
	assert.Nil(t, err)
	assert.Equal(t, connValsExp, val)
}
