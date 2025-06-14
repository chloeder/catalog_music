package memberships

import (
	"catalog-music/internal/configs"
	"catalog-music/internal/models/memberships"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {
	controlMock := gomock.NewController(t)
	defer controlMock.Finish()

	mockRepo := NewMockmembershipRepository(controlMock)

	type args struct {
		req *memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				req: &memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(uint(0), args.req.Email, args.req.Username).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "error_get_user",
			args: args{
				req: &memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(uint(0), args.req.Email, args.req.Username).Return(nil, assert.AnError)
			},
		},
		{
			name: "error_create_user",
			args: args{
				req: &memberships.SignUpRequest{
					Email:    "test@test.com",
					Username: "test",
					Password: "test",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(uint(0), args.req.Email, args.req.Username).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &service{
				cfg:      &configs.Config{},
				userRepo: mockRepo,
			}
			if err := r.SignUp(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
