package indicatorbundle

import "github.com/jinzhu/gorm"

type IndicatorMapperPSQL struct {
  db  *gorm.DB
}

func NewIndicatorMapperPSQL(db *gorm.DB) *IndicatorMapperPSQL {
  return &IndicatorMapperPSQL{ db: db }
}

func (imp *IndicatorMapperPSQL) FindAll() ([]Indicator, error) {
  var indicators []Indicator

  err := imp.db.Find(&indicators).Error

  return indicators, err
}

func (imp *IndicatorMapperPSQL) FindIndicatorById(id int) (Indicator, error) {
  var indicator Indicator

  err := imp.db.Where("id = ?", id).First(&indicator).Error

  return indicator, err
}

func (imp *IndicatorMapperPSQL) Insert(indicator *Report) error {
  return imp.db.Create(indicator).Error
}


func (imp *IndicatorMapperPSQL) Delete(int id) error {
  return imp.db.Delete(&Indicator{ Id: id }).Error
}


func (imp *IndicatorMapperPSQL) Update(indicator *Report) error {
  var i Indicator

  imp.db.First(&i, indicator.Id)
  i.Copy(indicator)

  return imp.db.Save(indicator).Error
}
