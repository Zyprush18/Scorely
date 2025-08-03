package serviceexamquest

import (
	"github.com/Zyprush18/Scorely/models/response"
	"github.com/Zyprush18/Scorely/repository/repoexamquestions"
)

type ServiceExamQuest interface {
	GetAllExamQuest(Search,Sort,coderole string, Page,Perpage,user_id,idexam int) ([]response.Exam_Questions, int64,error)
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