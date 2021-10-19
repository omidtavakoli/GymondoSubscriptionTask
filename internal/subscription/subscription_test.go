package subscription

import (
	"reflect"
	"testing"
)

func TestCreateService(t *testing.T) {
	type args struct {
		repository PgSQLRepository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: &service{
				psql:       nil,
				redis:      nil,
				logger:     nil,
				prometheus: nil,
				config:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateService(nil, nil, tt.args.repository, nil, nil, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestService_GetProductsList(t *testing.T) {
//	type args struct {
//		email       string
//		password    string
//		userProfile map[string]interface{}
//	}
//	tests := []struct {
//		description   string
//		args          args
//		initMock      func() Repository
//		expectedError error
//		expected      *User
//	}{
//		{
//			description: "return a validator error when email is not correctly formated",
//			args: args{
//				email:       "test",
//				password:    "some_non_hashed_password",
//				userProfile: nil,
//			},
//			expectedError: validatorError{errors.New("some validation error")},
//			initMock: func() Repository {
//				mockRep := new(MockRepository)
//				mockRep.On("Create", mock.Anything).Return(&User{}, nil)
//				return mockRep
//			},
//		}
//	}
//	for _, tt := range tests {
//		t.Run(tt.description, func(t *testing.T) {
//			repo := tt.initMock()
//			s := CreateService(repo)
//			got, err := s.Signup(tt.args.email, tt.args.password, tt.args.userProfile)
//			if !errorsHelper.AssertErrors(err, tt.expectedError) {
//				t.Fatalf("we got %v as error but expected %v", err, tt.expectedError)
//			}
//
//			if !reflect.DeepEqual(got, tt.expected) {
//				t.Fatalf("we got %v but expected %v", got, tt.expected)
//			}
//		})
//	}
//}
