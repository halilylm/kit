package worker_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/halilylm/kit/worker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("returns error if invalid number of max running jobs defined", func(t *testing.T) {
		cases := []struct {
			nofWorkers int
			err        error
		}{
			{-3, worker.ErrInvalidNumberOfMaxRunningJobs},
			{-2, worker.ErrInvalidNumberOfMaxRunningJobs},
			{-1, worker.ErrInvalidNumberOfMaxRunningJobs},
			{0, worker.ErrInvalidNumberOfMaxRunningJobs},
			{1, nil},
			{2, nil},
		}
		for _, testCase := range cases {
			_, err := worker.New(testCase.nofWorkers)
			assert.Equal(t, testCase.err, err)
		}
	})
}

func TestWorker_Start(t *testing.T) {
	t.Run("start the job", func(t *testing.T) {
		w, err := worker.New(5)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer w.Shutdown(context.Background())
		job := func(ctx context.Context) {}
		workerKey, err := w.Start(context.TODO(), job)
		assert.NoError(t, err)
		_, err = uuid.Parse(workerKey)
		assert.NoError(t, err)
		assert.Equal(t, 1, w.Running())
	})
	t.Run("can run multiple jobs", func(t *testing.T) {
		w, err := worker.New(5)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer w.Shutdown(context.Background())
		job := func(ctx context.Context) {}
		w.Start(context.TODO(), job)
		w.Start(context.TODO(), job)
		w.Start(context.TODO(), job)
		assert.Equal(t, 3, w.Running())
	})
	t.Run("cannot run more than defined capacity", func(t *testing.T) {
		w, err := worker.New(2)
		if err != nil {
			t.Fatal(err.Error())
		}
		defer w.Shutdown(context.Background())
		job := func(ctx context.Context) {}
		w.Start(context.TODO(), job)
		w.Start(context.TODO(), job)
		w.Start(context.TODO(), job)
		assert.Equal(t, 2, w.Running())
	})
}
