package graphql_connection

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/devishot/so-go-grpc-client_project/domain"
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

	pageArgsExp := domain.ProjectRepositoryPageArgs{
		First: 10,
		After: afterProject.Timestamp,
	}
	connArgs := domain.ConnectionArgumentsValue{
		First: pageArgsExp.First,
		After: domain.EncodeTimestampCursor(pageArgsExp.After),
	}
	connValsExp := &domain.ConnectionValue{
		Edges: []*domain.ConnectionEdgeValue{
			{
				Cursor: domain.EncodeTimestampCursor(projects[1].Timestamp),
				Node:   projects[1],
			},
			{
				Cursor: domain.EncodeTimestampCursor(projects[2].Timestamp),
				Node:   projects[2],
			},
		},
		PageInfo: domain.ConnectionPageInfoValue{
			HasNextPage: false,
			EndCursor:   domain.EncodeTimestampCursor(lastProject.Timestamp),
		},
	}

	connRepoMock := &ProjectConnectionRepositoryMock{
		GetLastByClientFunc: func(clientID domain.ID) (entity domain.ProjectEntity, e error) {
			assert.Equal(t, cID, clientID)
			return lastProject, nil
		},
		PaginateByTimestampFunc: func(clientID domain.ID, args domain.ProjectRepositoryPageArgs) (
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
