package db

import (
	"context"
	"database/sql"
	"practice/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTasks(t *testing.T) Task {
	arg := CreatetasksParams{
		Taskid:   int32(util.RandomInt(0, 1000)),
		Taskname: sql.NullString{String: util.RandomString(5), Valid: true},
		Tasktime: sql.NullString{String: util.RandomString(5), Valid: true},
		Taskdate: sql.NullString{String: util.RandomString(5), Valid: true},
	}

	tasks, err := testQueries.Createtasks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	require.Equal(t, arg.Taskid, tasks.Taskid)
	require.Equal(t, arg.Taskname, tasks.Taskname)
	require.Equal(t, arg.Tasktime, tasks.Tasktime)
	require.Equal(t, arg.Taskdate, tasks.Taskdate)

	return tasks
}

func TestCreatetasks(t *testing.T) {
	createRandomTasks(t)
}

func TestGettasks(t *testing.T){
	task1 := createRandomTasks(t)
	task2, err := testQueries.Gettasks(context.Background(),task1.Taskid)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.Taskid, task2.Taskid)
	require.Equal(t, task1.Taskname, task2.Taskname)
	require.Equal(t, task1.Tasktime, task2.Tasktime)
	require.Equal(t, task1.Taskdate, task2.Taskdate)

}

func TestUpdatetasks(t *testing.T){
	task1 := createRandomTasks(t)

	arg := UpdatetasksParams{
		Taskid: task1.Taskid,
	}

	task2, err := testQueries.Updatetasks(context.Background(),arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.Taskid, task2.Taskid)
	require.Equal(t, task1.Taskname, task2.Taskname)
	require.Equal(t, task1.Tasktime, task2.Tasktime)
	require.Equal(t, task1.Taskdate, task2.Taskdate)
}

func TestDeletetasks(t *testing.T){
	task1 := createRandomTasks(t)
	err := testQueries.Deletetasks(context.Background(),task1.Taskid)
	require.NoError(t,err)

	task2, err := testQueries.Gettasks(context.Background(),task1.Taskid)
	require.Error(t,err)
	require.EqualError(t,err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}

func TestListtasks(t *testing.T){
	for i := 0; i < 10; i++ {
		createRandomTasks(t)
	}

	tasks, err := testQueries.Listtasks(context.Background())
	require.NoError(t,err)
	require.Len(t,tasks,5)

	for _, account := range tasks {
		require.NotEmpty(t,account)
	}
}

