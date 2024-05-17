package validators

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestData struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Email       string    `validate:"required,min=3" json:"email"`
	FullName    string    `json:"fullname"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Location    string    `json:"location"`
	Enabled     bool      `json:"enabled"`
}

func TestValidator_Validate(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    error
	}{
		{
			name:    "given valid data should return no errors",
			args:    args{data: TestData{Email: "test@fake.com"}},
			wantErr: false,
			want:    nil,
		},
		{
			name:    "given invalid email should return error",
			args:    args{data: TestData{Email: "t@"}},
			wantErr: true,
			want:    errors.New("Key: 'TestData.Email' Error:Field validation for 'Email' failed on the 'min' tag"),
		},
		{
			name:    "given nil email should return error",
			args:    args{data: TestData{}},
			wantErr: true,
			want:    errors.New("Key: 'TestData.Email' Error:Field validation for 'Email' failed on the 'required' tag"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			got := v.Validate(tt.args.data)
			if got != nil && tt.wantErr {
				assert.Equal(t, tt.want.Error(), got.Error())
			}
		})
	}
}
