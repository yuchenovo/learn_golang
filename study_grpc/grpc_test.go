package main

import (
	"context"
	"fmt"
	student_service "g/idl/my_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestService(t *testing.T) {
	dial, err := grpc.Dial("127.0.0.1:2345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer dial.Close()
	client := student_service.NewStudentServiceClient(dial)
	info, err := client.GetStudentInfo(context.TODO(), &student_service.Request{StudentId: "1"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("%+v\n", info)
}
