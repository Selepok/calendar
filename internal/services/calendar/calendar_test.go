package calendar

import (
	"github.com/Selepok/calendar/internal/model"
	"github.com/golang/mock/gomock"
	"testing"
)

const (
	correctTitle   = "correct title"
	incorrectTitle = "incorrect title"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	repositoryMock.EXPECT().Create(model.Event{Title: correctTitle}).Return(nil)
	//repositoryMock.EXPECT().Create(model.Event{Title: incorrectTitle}).Return(nil)

	//service := Service{repositoryMock}

	//tests := []struct {
	//	name  string
	//	title string
	//	error error
	//}{
	//	{
	//		name:  "CreateUser create success",
	//		title: correctTitle,
	//		error: nil,
	//	},
	//}
	//assertion := assert.New(t)
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		err := service.CreateEvent(model.Event{Title: tt.title})
	//		assertion.Equalf(tt.error, err, "Test case: %s", tt.name)
	//	})
	//}
}
