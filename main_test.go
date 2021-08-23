package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	"goUT/client/mock"

	//adding this empty import to enable go.mod to download mockgen
	_ "github.com/golang/mock/mockgen/model"
)

func TestProcessInput(t *testing.T)  {
	myVariable := myStruct{}
	res, err := processInput(myVariable, 1,2)
	if err != nil {
		t.Fatal(err)
	}
	if res != 3{
		t.Fatalf("Expected result %d got %d", 3, res)
	}
}

func TestProcessInputTabularMethod(t *testing.T) {

	cases := []struct {
		name string
		num1 int
		num2 int
		res  int
		expectError bool
	}{
		{
			name: "Case 1",
			num1: 1,
			num2: 2,
			res:  3,
			expectError: false,
		},
		{
			name: "Case 2",
			num1: 100,
			num2: 200,
			res:  300,
			expectError: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			myVariable := myStruct{}
			res, err := processInput(myVariable, tc.num1,tc.num2)
			if err != nil && !tc.expectError{
				t.Fatal(err)
			}
				if res != tc.res{
					t.Fatalf("Expected result %d got %d", tc.res, res)
				}
		})
	}
}

func TestProcessInputCustomMethods(t *testing.T)  {
	myVariable := testStruct{}
	res, err := processInput(myVariable, 1,2)
	if err != nil {
		t.Fatal(err)
	}
	if res != 3{
		t.Fatalf("Expected result %d got %d", 3, res)
	}
}

type testStruct struct {

}

func (input testStruct) GetSum( num1, num2 int) (int, error) {
	return 1, nil
}

func TestProcessInputMocks(t *testing.T){
	mockCtrl := gomock.NewController(t)
	mockClient := mock.NewMockClient(mockCtrl)

	//mockClient.EXPECT().GetSum(11,2).Return(3, nil)
	//mockClient.EXPECT().GetSum(gomock.Any(),gomock.Any()).Return(3, fmt.Errorf("Intention Error "))
	//mockClient.EXPECT().GetSum(1,2).Return(33, nil)
	//mockClient.EXPECT().GetSum(1,2).Return(3, nil).Times(2)

	res, err := processInput(mockClient, 11, 2)
	if err != nil {
		t.Fatal(err)
	}
	if res != 3{
		t.Fatalf("Expected result %d got %d", 3, res)
	}

}
