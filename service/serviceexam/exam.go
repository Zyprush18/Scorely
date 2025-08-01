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
	ShowExams(id,userid int,coderole string) (*response.Exams,error)
	UpdateExam(data *request.Exams,role string,id,userid int) error
	DeleteExam(id,userid int,coderole string) error
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

func (r *RepoExams) ShowExams(id,userid int,coderole string) (*response.Exams,error) {
	return r.repo.Show(id,userid,coderole)
}

func (r *RepoExams) UpdateExam(data *request.Exams,role string,id,userid int) error  {
	return r.repo.Update(data,role,id,userid)
}

func (r *RepoExams) DeleteExam(id,userid int,coderole string) error {
	return r.repo.Delete(id,userid,coderole)
}