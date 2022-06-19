package ucweight

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/backend/domain/weight"
	"github.com/golang/mock/gomock"
)

func TestUsecase_CreateWeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	weightDomainMock := weight.NewMockDomainItf(ctrl)
	type fields struct {
		weightDomain weight.DomainItf
	}
	type args struct {
		weightData weight.WeightData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "default",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			args: args{
				weightData: weight.WeightData{
					ID:   1,
					Date: "2022-11-22",
					Min:  11,
					Max:  22,
				},
			},
			wantErr: false,
			mock: func() {
				weightDomainMock.EXPECT().CreateWeightData(gomock.Any()).Return(nil)
			},
		},
		{
			name: "error domain",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			args: args{
				weightData: weight.WeightData{
					ID:   1,
					Date: "2022-11-22",
					Min:  11,
					Max:  22,
				},
			},
			wantErr: true,
			mock: func() {
				weightDomainMock.EXPECT().CreateWeightData(gomock.Any()).Return(fmt.Errorf("error db"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				weightDomain: tt.fields.weightDomain,
			}
			if err := uc.CreateWeight(tt.args.weightData); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateWeight() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUsecase_GetAllWeightData(t *testing.T) {
	ctrl := gomock.NewController(t)
	weightDomainMock := weight.NewMockDomainItf(ctrl)
	type fields struct {
		weightDomain weight.DomainItf
	}
	tests := []struct {
		name     string
		fields   fields
		wantResp WeightResp
		wantErr  bool
		mock     func()
	}{
		{
			name: "default",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			wantResp: WeightResp{
				Data: []weight.WeightData{
					{
						Date: "2022-11-11",
						Min:  11,
						Max:  11,
						Diff: 0,
					},
				},
				Mean: Mean{
					Max:  11,
					Min:  11,
					Diff: 0,
				},
			},
			wantErr: false,
			mock: func() {
				weightDomainMock.EXPECT().GetAllWeightData().Return([]weight.WeightData{
					{
						Date: "2022-11-11",
						Min:  11,
						Max:  11,
					},
				}, nil)
			},
		},
		{
			name: "error db",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			wantResp: WeightResp{},
			wantErr:  true,
			mock: func() {
				weightDomainMock.EXPECT().GetAllWeightData().Return([]weight.WeightData{}, fmt.Errorf("error db"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				weightDomain: tt.fields.weightDomain,
			}
			gotResp, err := uc.GetAllWeightData()
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.GetAllWeightData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Usecase.GetAllWeightData() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestUsecase_EditWeightData(t *testing.T) {
	ctrl := gomock.NewController(t)
	weightDomainMock := weight.NewMockDomainItf(ctrl)
	type fields struct {
		weightDomain weight.DomainItf
	}
	type args struct {
		weightData weight.WeightData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "default",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			args: args{
				weightData: weight.WeightData{
					ID:   1,
					Date: "2022-11-22",
					Min:  11,
					Max:  22,
				},
			},
			wantErr: false,
			mock: func() {
				weightDomainMock.EXPECT().EditWeightData(gomock.Any()).Return(nil)
			},
		},
		{
			name: "error domain",
			fields: fields{
				weightDomain: weightDomainMock,
			},
			args: args{
				weightData: weight.WeightData{
					ID:   1,
					Date: "2022-11-22",
					Min:  11,
					Max:  22,
				},
			},
			wantErr: true,
			mock: func() {
				weightDomainMock.EXPECT().EditWeightData(gomock.Any()).Return(fmt.Errorf("error db"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			uc := &Usecase{
				weightDomain: tt.fields.weightDomain,
			}
			if err := uc.EditWeightData(tt.args.weightData); (err != nil) != tt.wantErr {
				t.Errorf("Usecase.EditWeightData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
