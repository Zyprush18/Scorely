package serviceexam

import (
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexams"
)


type ServiceExams interface {
	GetAllExams(Search,Sort string, Page,Perpage int) ([]response.Exams, int64,error)
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