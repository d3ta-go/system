package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func newDummyContext() echo.Context {
	req := httptest.NewRequest(http.MethodPost, "/dummy/test", nil)
	res := httptest.NewRecorder()
	return echo.New().NewContext(req, res)
}

func TestResult(t *testing.T) {
	type args struct {
		httpStatus int
		status     string
		data       interface{}
		msg        string
		c          echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Result: OK",
			args:    args{httpStatus: http.StatusOK, status: SUCCESS, data: nil, msg: "OK", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Result(tt.args.httpStatus, tt.args.status, tt.args.data, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Result() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOk(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test OK(): OK",
			args:    args{c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OK(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("OK() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOKWithMesssage(t *testing.T) {
	type args struct {
		msg string
		c   echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test OKWithMesssage(): OK",
			args:    args{msg: "OK", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OKWithMesssage(tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("OKWithMesssage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOKWithData(t *testing.T) {
	type args struct {
		data interface{}
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test OKWithData(): OK",
			args:    args{data: nil, c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OKWithData(tt.args.data, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("OKWithData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreatedWithData(t *testing.T) {
	type args struct {
		data interface{}
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test CreatedWithData(): OK",
			args:    args{data: nil, c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatedWithData(tt.args.data, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CreatedWithData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOKDetailed(t *testing.T) {
	type args struct {
		data interface{}
		msg  string
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test OKDetailed(): OK",
			args:    args{data: nil, msg: "OK", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OKDetailed(tt.args.data, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("OKDetailed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFail(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test Fail(): OK",
			args:    args{c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Fail(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Fail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFailWithMessage(t *testing.T) {
	type args struct {
		msg string
		c   echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test FailWithMessage(): OK",
			args:    args{msg: "Fail Message", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FailWithMessage(tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FailWithMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFailWithDetailed(t *testing.T) {
	type args struct {
		status string
		data   interface{}
		msg    string
		c      echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test FailWithDetailed: OK",
			args:    args{status: ERROR, data: nil, msg: "Fail Message", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FailWithDetailed(tt.args.status, tt.args.data, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FailWithDetailed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFailWithDataMessage(t *testing.T) {
	type args struct {
		data interface{}
		msg  string
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test FailWithDataMessage(): OK",
			args:    args{data: nil, msg: "Fail Message", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FailWithDataMessage(tt.args.data, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FailWithDataMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFailDetailedwithCode(t *testing.T) {
	type args struct {
		code int
		data interface{}
		msg  string
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test FailDetailedwithCode(): OK",
			args:    args{code: http.StatusInternalServerError, data: nil, msg: "Fail Message", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FailDetailedwithCode(tt.args.code, tt.args.data, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FailDetailedwithCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFailWithMessageWithCode(t *testing.T) {
	type args struct {
		code int
		msg  string
		c    echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test FailWithMessageWithCode(): OK",
			args:    args{code: http.StatusInternalServerError, msg: "Fail Message", c: newDummyContext()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FailWithMessageWithCode(tt.args.code, tt.args.msg, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("FailWithMessageWithCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
