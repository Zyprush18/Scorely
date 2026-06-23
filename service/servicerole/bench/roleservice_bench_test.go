package bench

import (
	"context"
	"testing"
	"time"

	"github.com/Zyprush18/Scorely/helper"
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/service/servicerole"
)

func BenchmarkGetAllData(b *testing.B) {
	now := time.Now()
	mock := new(servicerole.RepoRoleMock)
	service := servicerole.NewRoleService(mock)
	data := []response.Roles{
		{
			IdRole:   1,
			NameRole: "Admin",
			Model: helper.Models{
				CreatedAt: now,
			},
		},
		{
			IdRole:   2,
			NameRole: "User",
			Model: helper.Models{
				CreatedAt: now,
			},
		},
	}

	mock.On("GetAllDataRole", context.Background(),"","",1,10).Return(data, 2,nil)

	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_, _,_ = service.GetAllData(context.Background(),"","",1,10)
	}
	b.StopTimer()
	mock.AssertExpectations(b)
}

func BenchmarkCreateUser(b *testing.B) {
	mock := new(servicerole.RepoRoleMock)
	service := servicerole.NewRoleService(mock)
	rolePass := &request.Roles{
		NameRole: "Admin",
	}
	mock.On("CreateRole", context.Background(), rolePass).Return(nil)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop() {
		_ = service.Create(context.Background(), rolePass)
	}
	b.StopTimer()
	mock.AssertExpectations(b)
}

func BenchmarkShowByIdUser(b *testing.B)  {
	mock := new(servicerole.RepoRoleMock)
	service := servicerole.NewRoleService(mock)
	data := &response.Roles{
		IdRole:   1,
		NameRole: "Admin",
	}

	mock.On("ShowById", context.Background(), 1).Return(data, nil)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop(){
		_,_=service.ShowRoleById(context.Background(), 1)
	}
	b.StopTimer()
	mock.AssertExpectations(b)
}

func BenchmarkUpdate(b *testing.B)  {
	mock := new(servicerole.RepoRoleMock)
	service := servicerole.NewRoleService(mock)
	rolePass := &request.Roles{
		NameRole: "Admin",
	}

	mock.On("UpdateRole", context.Background(), 1, rolePass).Return(nil)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop(){
		_= service.UpdateRole(context.Background(), 1, rolePass)
	}
	b.StopTimer()
	mock.AssertExpectations(b)
}

func BenchmarkDelete(b *testing.B)  {
	mock := servicerole.RepoRoleMock{}
	service := servicerole.NewRoleService(&mock)
	mock.On("DeleteRole", context.Background(), 1).Return(nil)
	b.ResetTimer()
	b.ReportAllocs()
	for b.Loop(){
		_ = service.DeleteRole(context.Background(), 1)
	}
	b.StopTimer()
	mock.AssertExpectations(b)
}