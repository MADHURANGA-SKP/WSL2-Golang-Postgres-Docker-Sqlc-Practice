package db

import (
	"context"
	"database/sql"
	"practice/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatetasks(t *testing.T)  {
	arg := CreatetasksParams{  
		Taskid: int32(util.RandomInt(0,10)),
		Taskname: sql.NullString{String: util.RandomString(5), Valid: true},
		Tasktime: sql.NullString{String: util.RandomString(5), Valid: true},
		Taskdate: sql.NullString{String: util.RandomString(5), Valid: true},
	}

	tasks, err := testQueries.Createtasks(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,tasks)

	require.Equal(t,arg.Taskid, tasks.Taskid)
	require.Equal(t,arg.Taskname, tasks.Taskname)
	require.Equal(t,arg.Tasktime, tasks.Tasktime)
	require.Equal(t,arg.Taskdate, tasks.Taskdate)

	require.NoError(t,err)
	require.NoError(t,err)

}