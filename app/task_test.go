package app

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/dattranman/todo/model"
	"github.com/dattranman/todo/model/request"
	"github.com/dattranman/todo/model/response"
	"github.com/dattranman/todo/model/schema"
	"github.com/dattranman/todo/store"
	"github.com/dattranman/todo/store/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	onErr bool
}

func (r *MockTaskRepository) Create(schema.Task) error {
	if r.onErr {
		return errors.New("error while inserting")
	}
	return nil
}

func TestApp_CreateTask(t *testing.T) {
	type fields struct {
		StoreErr error
		Config   model.Configuration
		Store    store.Store
	}
	type args struct {
		task request.Task
	}
	timeNow := time.Now()
	taskStore := &mocks.TaskStore{}
	store := &mocks.Store{}
	store.On("Task").Return(taskStore)
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.Task
		wantErr  bool
	}{
		// Test cases.
		{
			name: "Create success",
			fields: fields{
				Store: store,
			},
			args: args{
				task: request.Task{
					Name:     "test",
					Time:     timeNow,
					Point:    1,
					Priority: 1,
				},
			},
			wantResp: response.Task{
				Name:     "test",
				Time:     timeNow,
				Point:    1,
				Priority: 1,
			},
			wantErr: false,
		},
		{
			name: "Create fail",
			fields: fields{
				StoreErr: errors.New("errors"),
				Store:    store,
			},
			args: args{
				task: request.Task{
					Name:     "test",
					Point:    1,
					Priority: 1,
				},
			},
			wantResp: response.Task{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskStore.On("Create", mock.AnythingOfType("*schema.Task")).
				Return(tt.fields.StoreErr).
				Once()
			a := &App{
				Store: tt.fields.Store,
			}

			gotResp, err := a.CreateTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotResp.Name, tt.wantResp.Name)
			assert.Equal(t, gotResp.Point, tt.wantResp.Point)
			assert.Equal(t, gotResp.Priority, tt.wantResp.Priority)
			assert.WithinDuration(t, gotResp.Time, tt.wantResp.Time, 0)
		})
	}
}

func TestApp_DeleteTaskByID(t *testing.T) {
	type fields struct {
		StoreErr error
		Config   model.Configuration
		Store    store.Store
	}
	type args struct {
		id string
	}
	taskStore := &mocks.TaskStore{}
	store := &mocks.Store{}
	store.On("Task").Return(taskStore)
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.UtilResponse
		wantErr  bool
	}{
		{
			name: "Delete success",
			fields: fields{
				Store: store,
			},
			args: args{
				id: "test",
			},
			wantResp: response.UtilResponse{
				Status: response.StatusSuccess,
			},
			wantErr: false,
		}, {
			name: "Delete fail",
			fields: fields{
				StoreErr: errors.New("errors"),
				Store:    store,
			},
			args: args{
				id: "test",
			},
			wantResp: response.UtilResponse{
				Status: response.StatusFail,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskStore.On("Delete", mock.AnythingOfType("string")).
				Return(tt.fields.StoreErr).
				Once()
			a := &App{
				Store: tt.fields.Store,
			}
			gotResp, err := a.DeleteTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.DeleteTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("App.DeleteTaskByID() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestApp_UpdateTask(t *testing.T) {
	type fields struct {
		Config   model.Configuration
		Store    store.Store
		StoreErr error
	}
	type args struct {
		req request.Task
	}

	taskStore := &mocks.TaskStore{}
	store := &mocks.Store{}
	store.On("Task").Return(taskStore)
	timeNow := time.Now()
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.Task
		wantErr  bool
	}{
		{
			name: "Update success",
			fields: fields{
				Store: store,
			},
			args: args{
				req: request.Task{
					Name:     "test",
					Time:     timeNow,
					Point:    1,
					Priority: 1,
				},
			},
			wantResp: response.Task{
				Name:     "test",
				Time:     timeNow,
				Point:    1,
				Priority: 1,
			},
			wantErr: false,
		},
		{
			name: "Update fail",
			fields: fields{
				StoreErr: errors.New("errors"),
				Store:    store,
			},
			args: args{
				req: request.Task{
					Name:     "test",
					Point:    1,
					Priority: 1,
				},
			},
			wantResp: response.Task{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskStore.On("Update", mock.AnythingOfType("*schema.Task")).
				Return(tt.fields.StoreErr).
				Once()
			a := &App{
				Store: tt.fields.Store,
			}
			gotResp, err := a.UpdateTask(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotResp.Name, tt.wantResp.Name)
			assert.Equal(t, gotResp.Point, tt.wantResp.Point)
			assert.Equal(t, gotResp.Priority, tt.wantResp.Priority)
			assert.WithinDuration(t, gotResp.Time, tt.wantResp.Time, 0)
		})
	}
}

func TestApp_GetTaskByID(t *testing.T) {
	type fields struct {
		Config        model.Configuration
		Store         store.Store
		StoreErr      error
		StoreRespTask schema.Task
	}
	type args struct {
		id string
	}
	taskStore := &mocks.TaskStore{}
	store := &mocks.Store{}
	store.On("Task").Return(taskStore)
	timeNow := time.Now()
	testID := "test_id"
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp response.Task
		wantErr  bool
	}{
		{
			name: "GetTaskByID success",
			fields: fields{
				Store: store,
				StoreRespTask: schema.Task{
					ID:       testID,
					Name:     "test",
					Time:     timeNow,
					Point:    1,
					Priority: 1,
				},
			},
			args: args{
				id: testID,
			},
			wantResp: response.Task{
				ID:       testID,
				Name:     "test",
				Time:     timeNow,
				Point:    1,
				Priority: 1,
			},
			wantErr: false,
		},
		{
			name: "GetTaskByID fail",
			fields: fields{
				StoreErr: errors.New("errors"),
				Store:    store,
			},
			args: args{
				id: testID,
			},
			wantResp: response.Task{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskStore.On("GetByID", mock.AnythingOfType("string")).
				Return(tt.fields.StoreRespTask, tt.fields.StoreErr).
				Once()
			a := &App{
				Store: tt.fields.Store,
			}
			gotResp, err := a.GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotResp.Name, tt.wantResp.Name)
			assert.Equal(t, gotResp.Point, tt.wantResp.Point)
			assert.Equal(t, gotResp.Priority, tt.wantResp.Priority)
			assert.WithinDuration(t, gotResp.Time, tt.wantResp.Time, 0)
		})
	}
}

func TestApp_GetListTasks(t *testing.T) {
	type fields struct {
		Config        model.Configuration
		Store         store.Store
		StoreRespTask []schema.Task
		StoreErr      error
	}
	taskStore := &mocks.TaskStore{}
	store := &mocks.Store{}
	store.On("Task").Return(taskStore)
	timeNow := time.Now()
	testID := "test_id"
	tests := []struct {
		name     string
		fields   fields
		wantResp response.ListTasks
		wantErr  bool
	}{
		{
			name: "GetTaskByID success",
			fields: fields{
				Store: store,
				StoreRespTask: []schema.Task{{
					ID:       testID,
					Name:     "test",
					Time:     timeNow,
					Point:    1,
					Priority: 1,
				},
				},
			},
			wantResp: response.ListTasks{
				List: []response.Task{
					{
						ID:       testID,
						Name:     "test",
						Time:     timeNow,
						Point:    1,
						Priority: 1,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "GetTaskByID fail",
			fields: fields{
				StoreErr:      errors.New("errors"),
				Store:         store,
				StoreRespTask: nil,
			},
			wantResp: response.ListTasks{
				List: []response.Task{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskStore.On("GetList").
				Return(tt.fields.StoreRespTask, int64(0), tt.fields.StoreErr).
				Once()
			a := &App{
				Store: tt.fields.Store,
			}
			gotResp, err := a.GetListTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("App.GetListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := range gotResp.List {
				assert.Equal(t, gotResp.List[i].Name, tt.wantResp.List[i].Name)
				assert.Equal(t, gotResp.List[i].Point, tt.wantResp.List[i].Point)
				assert.Equal(t, gotResp.List[i].Priority, tt.wantResp.List[i].Priority)
				assert.WithinDuration(t, gotResp.List[i].Time, tt.wantResp.List[i].Time, 0)
			}
		})
	}
}
