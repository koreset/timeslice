package services

import (
	"github.com/koreset/timeslice/models"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService struct {
}

func init() {
	if db == nil {
		db = GetDB()
	}
}

func (e *EmployeeService) Create(emp *models.Employee) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(emp.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	emp.PasswordHash = string(hashed)
	emp.Password = ""
	if err := db.Create(&emp).Error; err != nil {
		return err
	}
	return nil
}

func (e *EmployeeService) Update(emp *models.Employee) error {
	return db.Save(emp).Error
}

func (e *EmployeeService) EmployeeByID(id uint) (models.Employee, error) {
	var emp models.Employee

	if err = db.Where("id = ?", id).First(&emp).Error; err != nil {
		return emp, err
	}

	return emp, nil
}

func (e *EmployeeService) EmployeeByEmail(email string) (models.Employee, error) {
	var emp models.Employee

	if err = db.Where("email = ?", email).First(&emp).Error; err != nil {
		return emp, err
	}

	return emp, nil
}
