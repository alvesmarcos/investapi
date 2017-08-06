package reportbundle

import "github.com/jinzhu/gorm"

type ReportMapperPSQL struct {
  db *gorm.DB
}

func NewReportMapperPSQL(db *gorm.DB) *ReportMapperPSQL {
  return &ReportMapperPSQL{ db: db }
}

func (rmp *ReportMapperPSQL) FindAll() ([]Report, error) {
  var reports []Report

  err := rmp.db.Find(&reports).Error

  return reports, err
}

func (rmp *ReportMapperPSQL) FindReportById(id int) (Report, error) {
  var report Report

  err := rmp.db.Where("id = ?", id).First(&report).Error

  return report, err
}

func (rmp *ReportMapperPSQL) Insert(report *Report) error {
  return rmp.db.Create(report).Error
}

func (rmp *ReportMapperPSQL) Delete(id int) error {
  return rmp.db.Delete(&Report{Id: id}).Error
}

func (rmp *ReportMapperPSQL) Update(report *Report) error {
  var r Report

  rmp.db.First(&r, report.Id)
  r.Copy(report)

  return rmp.db.Save(&r).Error
}
