package serviceexam

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexams"
)


type ServiceExams interface {
	GetAllExams(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error)
	FindExamsbyIdTeacher(Search,Sort string, Page,Perpage,id int) ([]response.Exams, int64,error)
	CreateExams(data *request.Exams, role string,user_id,subject_id int) error
}

type RepoExams struct {
	repo repoexams.RepoExams
}

func ConnectRepo(r repoexams.RepoExams) ServiceExams {
	return &RepoExams{repo: r}
}

func (r *RepoExams) GetAllExams(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error) {
	return r.repo.GetAll(Search,Sort,Page,Perpage)
}

func (r *RepoExams) FindExamsbyIdTeacher(Search,Sort string, Page,Perpage,id int) ([]response.Exams, int64,error) {
	return r.repo.FindByidTeacher(Search,Sort,Page,Perpage,id)
}

func (r *RepoExams) CreateExams(data *request.Exams, role string,user_id,subject_id int) error {
	return r.repo.Create(data,role,user_id,subject_id)
}