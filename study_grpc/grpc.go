package main

import (
	"context"
	"fmt"
	student_service "g/idl/my_proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

type StudentServer struct {
}

func GetInfo(StudentId string) student_service.Student {
	stu := student_service.Student{}
	stu.Name = "ww"
	stu.Birthday = timestamppb.Now()
	return stu
}
func (s *StudentServer) GetStudentInfo(ctx context.Context, request *student_service.Request) (*student_service.Student, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error:%+v\n", err)
		}
	}()
	studentId := request.StudentId
	student := GetInfo(studentId)
	return &student, nil
}
func main() {
	listen, err := net.Listen("tcp", ":2345")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, new(StudentServer))
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
