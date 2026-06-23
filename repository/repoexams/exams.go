package repoexams

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Zyprush18/Scorely/models/entity"
	"gorm.io/gorm"
)

type RepoExams interface {
	GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Exams, int64, error)
	FindByidTeacher(ctx context.Context, Search, Sort string, Page, Perpage, id int) ([]entity.Exams, int64, error)
	Create(ctx context.Context, data *entity.Exams) error
	Show(ctx context.Context, id, userid int, coderole string) (*entity.Exams, error)
	Update(ctx context.Context, id int, data *entity.Exams) error
	Delete(ctx context.Context, id, userid int, coderole string) error

	TxExams(ctx context.Context, fn func(repo RepoExams) error ) error
	CheckRoleforCreateOrUpdate(ctx context.Context, role string, user_id, subject_id int, teacher_id *uint) (uint, error)
}

type MysqlStruct struct {
	db *gorm.DB
}

func ConnectDb(d *gorm.DB) RepoExams {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) WithTx(d *gorm.DB) RepoExams {
	return &MysqlStruct{db: d}
}

func (m *MysqlStruct) TxExams(ctx context.Context, fn func(repo RepoExams) error ) error  {
	tx := m.db.WithContext(ctx).Begin()
    if tx.Error != nil {
        return tx.Error
    }
    repo := m.WithTx(tx)
    err := fn(repo)
    if err != nil {
        tx.Rollback()
        return err
    }
    return tx.Commit().Error
}



func (m *MysqlStruct) GetAll(ctx context.Context, Search, Sort string, Page, Perpage int) ([]entity.Exams, int64, error) {
	var modelexam []entity.Exams
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.WithContext(ctx).Model(&entity.Exams{}).Debug().Preload("TeacherSubject.Subject").Preload("ExamQuestion").Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&modelexam).Error; err != nil {
		return nil, 0, err
	}

	return modelexam, count, nil
}

func (m *MysqlStruct) FindByidTeacher(ctx context.Context, Search, Sort string, Page, Perpage, id int) ([]entity.Exams, int64, error) {
	var finddata []entity.Exams
	var count int64
	order := fmt.Sprintf("created_at %s", Sort)
	offset := (Page - 1) * Perpage

	if err := m.db.WithContext(ctx).Model(&entity.Exams{}).Debug().Preload("TeacherSubject.Subject").Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = exams.teacher_subject_id").Joins("JOIN subjects AS s ON s.id_subject = ts.id_subjects").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ?", id).Where("name_exams LIKE ?", "%"+Search+"%").Count(&count).Order(order).Limit(Perpage).Offset(offset).Find(&finddata).Error; err != nil {
		return nil, 0, err
	}

	return finddata, count, nil
}

func (m *MysqlStruct) Create(ctx context.Context, data *entity.Exams) error {
	return  m.db.WithContext(ctx).Model(&entity.Exams{}).Create(data).Error
}



func (m *MysqlStruct) Show(ctx context.Context, id, userid int, coderole string) (*entity.Exams, error) {
	var modelexam entity.Exams
	query := m.db.WithContext(ctx).Model(&entity.Exams{}).Preload("TeacherSubject.Subject")
	switch coderole {
	case "admin":
		if err := query.Where("id_exam = ?", id).First(&modelexam).Error; err != nil {
			return nil, err
		}
	case "teacher":
		if err := query.Joins("JOIN teacher_subjects AS ts ON ts.id_teacher_subject = exams.teacher_subject_id").Joins("JOIN subjects AS s ON s.id_subject = ts.id_subjects").Joins("JOIN teachers AS t ON t.id_teacher = ts.id_teachers").Where("t.user_id = ? AND id_exam = ?", userid, id).First(&modelexam).Error; err != nil {
			return nil, err
		}
	default:
		log.Println("Invalid Code Role")

		return nil, errors.New("invalid code Role")
	}

	return &modelexam, nil
}

func (m *MysqlStruct) Update(ctx context.Context, id int, data *entity.Exams) error {
	result := m.db.WithContext(ctx).Model(&entity.Exams{}).Where("id_exam = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *MysqlStruct) Delete(ctx context.Context, id, userid int, coderole string) error {
	switch coderole {
	case "teacher":
		resp := m.db.WithContext(ctx).Debug().Exec("DELETE exams FROM exams JOIN teacher_subjects AS ts ON ts.id_teacher_subject = exams.teacher_subject_id JOIN teachers AS t ON t.id_teacher = ts.id_teachers WHERE t.user_id = ? AND id_exam = ?", userid, id)
		if resp.Error != nil  {
			return resp.Error
		}

		if resp.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

	default:
		result := m.db.WithContext(ctx).Model(&entity.Exams{}).Delete(id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

func (r *MysqlStruct) CheckRoleforCreateOrUpdate(ctx context.Context, role string, user_id, subject_id int, teacher_id *uint) (uint, error) {
	var model_teachersubject entity.TeacherSubjects
	query := r.db.WithContext(ctx).Model(&entity.TeacherSubjects{})
	switch role {
	case "teacher":
		if err := query.Debug().Joins("JOIN teachers AS t ON t.id_teacher = teacher_subjects.id_teachers").Where("t.user_id = ? AND id_subjects = ?", user_id, subject_id).First(&model_teachersubject).Error; err != nil {
			return 0, err
		}
		return model_teachersubject.IdTeacherSubject, nil
	default:
		if err := query.Debug().Where("id_teachers = ? AND id_subjects = ?", teacher_id, subject_id).First(&model_teachersubject).Error; err != nil {
			return 0, err
		}
		return model_teachersubject.IdTeacherSubject, nil
	}
}