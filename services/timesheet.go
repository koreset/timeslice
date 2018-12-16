package services

import "github.com/koreset/timeslice/models"

type TimesheetService struct {
}

func (ts *TimesheetService) CreateEntry(tsEntry *models.TimesheetEntry) error {
	if err = db.Create(tsEntry).Error; err != nil {
		return err
	}
	return nil

}

// UpdateEntry is to geared to make changes to either the description of the entry
// or to update the duration of the entry. Duration validations will have to be made
// to ensure that values do not exceed configured upper and lower bounds
func (ts *TimesheetService) UpdateEntry(tsEntry *models.TimesheetEntry) error {
	return db.Save(tsEntry).Error
}

func (ts *TimesheetService) FindEntryByID(id uint) (models.TimesheetEntry, error) {
	var found models.TimesheetEntry
	err = db.Where("id = ?", id).First(&found).Error
	if err != nil {
		return found, err
	}
	return found, nil
}

func (ts *TimesheetService) FindEntriesByDateRangeForEmployee(startDate, endDate string, employeeID uint) ([]models.TimesheetEntry, error) {
	var entries []models.TimesheetEntry
	err = db.Where("employee_id = ? and entry_date between ? and ?", employeeID, startDate, endDate).Find(&entries).Error
	if err != nil {
		return entries, err
	}

	return entries, nil
}
