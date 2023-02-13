package mongo

import (
	"context"
	model "github.com/peacewalker122/project/service/db/model/mongo"
	mongoUtil "github.com/peacewalker122/project/util/db/mongo"
	"github.com/peacewalker122/project/util/file"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestNewStoryRepository(t *testing.T) {
	client, err := mongoUtil.InitMongoClient("mongodb+srv://admin:passwold@project-mongo.hbrb4az.mongodb.net/?retryWrites=true&w=majority")
	require.NoError(t, err)
	testRepo := NewStoryRepository(client.Database("test"))

	testCases := []struct {
		name     string
		testCase func(t *testing.T)
	}{
		{
			name: "test create",
			testCase: func(t *testing.T) {
				err = testRepo.Create(context.Background(), &model.Story{
					ID:      "",
					Content: file.MIME("image/jpeg"),
				})

				require.NoError(t, err)
			},
		},
		{
			name: "test getall",
			testCase: func(t *testing.T) {
				result, err := testRepo.GetAll(context.Background(), []string{"1", "2", "4"})
				log.Println(result)
				require.NoError(t, err)
				require.True(t, len(result) > 2)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, tc.testCase)
	}
}
