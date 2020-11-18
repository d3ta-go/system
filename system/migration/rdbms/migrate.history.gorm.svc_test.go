package rdbms

import (
	"reflect"
	"testing"
)

func newGormMigrationHistorySvc(t *testing.T) (*GormMigrationHistoryService, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, err
	}

	dbGorm, err := h.GetGormDB("db-identity")
	if err != nil {
		return nil, err
	}

	svc, err := NewGormMigrationHistoryService(h, dbGorm)
	if err != nil {
		return nil, err
	}

	return svc, err
}

func TestGormMigrationHistoryService_FindByScriptName(t *testing.T) {
	svc, err := newGormMigrationHistorySvc(t)
	if err != nil {
		t.Errorf("Error: newGormMigrationHistorySvc(), [%s]", err.Error())
		return
	}
	if svc != nil {
		type args struct {
			scriptName string
		}
		tests := []struct {
			name      string
			args      args
			wantExist bool
			wantEtt   *GormMigrationHistoryEntity
			wantErr   bool
		}{
			// TODO: Add test cases.
			{
				name:      "not found",
				args:      args{scriptName: "not-found-script"},
				wantExist: false,
				wantEtt:   &GormMigrationHistoryEntity{},
				wantErr:   false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotExist, gotEtt, err := svc.FindByScriptName(tt.args.scriptName)
				if (err != nil) != tt.wantErr {
					t.Errorf("GormMigrationHistoryService.FindByScriptName() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if gotExist != tt.wantExist {
					t.Errorf("GormMigrationHistoryService.FindByScriptName() gotExist = %v, want %v", gotExist, tt.wantExist)
				}
				if !reflect.DeepEqual(gotEtt, tt.wantEtt) {
					t.Errorf("GormMigrationHistoryService.FindByScriptName() gotEtt = %#v, want %#v", gotEtt, tt.wantEtt)
				}
			})
		}
	}
}

func TestGormMigrationHistoryService_HasBeenExecuted(t *testing.T) {
	svc, err := newGormMigrationHistorySvc(t)
	if err != nil {
		t.Errorf("Error: newGormMigrationHistorySvc(), [%s]", err.Error())
		return
	}
	if svc != nil {
		type args struct {
			scriptName string
		}
		tests := []struct {
			name    string
			args    args
			want    bool
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "not found",
				args:    args{scriptName: "not-found-script"},
				want:    false,
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := svc.HasBeenExecuted(tt.args.scriptName)
				if (err != nil) != tt.wantErr {
					t.Errorf("GormMigrationHistoryService.HasBeenExecuted() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GormMigrationHistoryService.HasBeenExecuted() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestGormMigrationHistoryService_SaveRunExecution(t *testing.T) {
	svc, err := newGormMigrationHistorySvc(t)
	if err != nil {
		t.Errorf("Error: newGormMigrationHistorySvc(), [%s]", err.Error())
		return
	}
	if svc != nil {
		type args struct {
			scriptName string
			scriptType string
			note       string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "save run execution OK - migrate",
				args:    args{scriptName: "dummy-migrate-test-script", scriptType: "MIGRATE", note: ""},
				wantErr: false,
			},
			{
				name:    "save run execution OK - seed",
				args:    args{scriptName: "dummy-seed-test-script", scriptType: "SEED", note: ""},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := svc.SaveRunExecution(tt.args.scriptName, tt.args.scriptType, tt.args.note); (err != nil) != tt.wantErr {
					t.Errorf("GormMigrationHistoryService.SaveRunExecution() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func TestGormMigrationHistoryService_SaveRollBackExecution(t *testing.T) {
	svc, err := newGormMigrationHistorySvc(t)
	if err != nil {
		t.Errorf("Error: newGormMigrationHistorySvc(), [%s]", err.Error())
		return
	}
	if svc != nil {
		type args struct {
			scriptName string
			note       string
		}
		tests := []struct {
			name string

			args    args
			wantErr bool
		}{
			// TODO: Add test cases.
			{
				name:    "save rollback execution OK - migrate",
				args:    args{scriptName: "dummy-migrate-test-script", note: ""},
				wantErr: false,
			},
			{
				name:    "save rollback execution OK - seed",
				args:    args{scriptName: "dummy-seed-test-script", note: ""},
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := svc.SaveRollBackExecution(tt.args.scriptName, tt.args.note); (err != nil) != tt.wantErr {
					t.Errorf("GormMigrationHistoryService.SaveRollBackExecution() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}
