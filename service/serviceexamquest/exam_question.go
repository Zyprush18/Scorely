package serviceexamquest

import (
	"github.com/Zyprush18/Scorely/models/request"
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexamquestions"
)

type ServiceExamQuest interface {
	GetAllExamQuest(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error)
	CreateExamQuest(data *request.Exam_Questions,userid,id_exam int, coderole string) error
	ShowExamQuest(id,user_id,exam_id int,coderole string) (*response.Exam_Questions,error)
}

type RepoExamQuest struct {
	repo repoexamquestions.RepoExamQuest
}

func ConnectRepo(r repoexamquestions.RepoExamQuest) ServiceExamQuest {
	return &RepoExamQuest{repo: r}
}

func (r *RepoExamQuest) GetAllExamQuest(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error) {
	return r.repo.GetAll(Search,Sort,coderole,Page,Perpage,user_id,idexam)
}

func (r *RepoExamQuest) CreateExamQuest(data *request.Exam_Questions,userid,id_exam int,coderole string) error {
	return r.repo.Create(data,userid,id_exam,coderole)
}

func (r *RepoExamQuest) ShowExamQuest(id,user_id,exam_id int,coderole string) (*response.Exam_Questions,error)  {
	return r.repo.Show(id,user_id,exam_id,coderole)
}