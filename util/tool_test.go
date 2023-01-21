package util

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestRandomFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestRandomFileName",
			args: args{
				fileName: "test.jpg",
			},
			want: len("test.jpg"),
		},
		{
			name: "TestRandomFileName2",
			args: args{
				fileName: "test.jpg",
			},
			want: len("test.jpg"),
		},
		{
			name: "TestWithDoubleExtension",
			args: args{
				fileName: "test.jpg.png",
			},
			want: len("test.jpg.png"),
		},
		{
			name: "TestWithNoExtension",
			args: args{
				fileName: "test",
			},
			want: "file is not valid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hashName, err := RandomFileName(tt.args.fileName)
			if err != nil {
				if err.Error() == "file is not valid" {
					require.Equal(t, tt.want, err.Error())
					return
				}
			}
			require.NoError(t, err)

			log.Println(hashName)

			require.Greater(t, len(hashName), tt.want.(int))
		})
	}
}
