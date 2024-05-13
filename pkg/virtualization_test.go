package client

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/synology-community/synology-api/pkg/api/virtualization"
)

func Test_virtualizationClient_ImageCreate(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.TaskRef
		wantErr bool
	}{
		{
			name: "Create image",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: ctx,
			},
			want:    &virtualization.TaskRef{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.client.VirtualizationAPI()

			s, err := tt.fields.client.Virtualization.StorageList(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			storageId := s.Storages[0].ID

			got, err := v.ImageCreate(tt.args.ctx, virtualization.Image{
				Name: "testmantic",
				Storages: virtualization.Storages{
					{ID: storageId},
				},
				Type:      "iso",
				FilePath:  "/data/cluster_storage/commoninit.iso",
				AutoClean: false,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.ImageCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.NotNil(t, got) {
				t.Errorf("Task ID: %s", got.TaskInfo.ImageID)
			}
		})
	}
}

func Test_virtualizationClient_ImageDelete(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx     context.Context
		imageID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.TaskRef
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.client.Virtualization
			err := v.ImageDelete(tt.args.ctx, tt.args.imageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.ImageDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_virtualizationClient_ImageList(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.ImageList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: tt.fields.client,
			}
			got, err := v.ImageList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.ImageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.ImageList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_TaskGet(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx    context.Context
		taskID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: tt.fields.client,
			}
			got, err := v.TaskGet(tt.args.ctx, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.TaskGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.TaskGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_GuestGet(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.Guest
		wantErr bool
	}{
		{
			name: "Get guest",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx:  context.Background(),
				name: "testmantic",
			},
			want: &virtualization.Guest{
				ID:          "1",
				Name:        "testmantic",
				Description: "Testmantic",
				Status:      "stopped",
				StorageID:   "1",
				StorageName: "default",
				Autorun:     0,
				VcpuNum:     1,
				VramSize:    512,
				Disks:       virtualization.VDisks{},
				Networks:    virtualization.VNICs{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: tt.fields.client,
			}
			got, err := v.GuestGet(tt.args.ctx, virtualization.Guest{
				Name: tt.args.name,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.GetGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.GetGuest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_ListGuests(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.GuestList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &virtualizationClient{
				client: newClient(t),
			}
			got, err := v.GuestList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.ListGuests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("virtualizationClient.ListGuests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVirtualizationClient(t *testing.T) {
	type args struct {
		client *APIClient
	}
	tests := []struct {
		name string
		args args
		want virtualization.VirtualizationAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVirtualizationClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVirtualizationClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_virtualizationClient_StorageList(t *testing.T) {
	type fields struct {
		client *APIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *virtualization.StorageList
		wantErr bool
	}{
		{
			name: "List storage",
			fields: fields{
				client: newClient(t),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &virtualization.StorageList{
				Storages: []virtualization.Storage{
					{
						ID:         "1",
						Name:       "default",
						Status:     "normal",
						HostName:   "localhost",
						HostID:     "1",
						Size:       1000000000,
						Used:       0,
						VolumePath: "/volume1",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.client.VirtualizationAPI()
			got, err := v.StorageList(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("virtualizationClient.StorageList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, len(got.Storages), len(tt.want.Storages)) {
				t.Errorf("virtualizationClient.StorageList() = %v, want %v", got, tt.want)
			}
		})
	}
}