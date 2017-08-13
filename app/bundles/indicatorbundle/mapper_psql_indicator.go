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
  var samples []Sample

  err := imp.db.Find(&indicators).Error

  if err != nil {
    return nil, err
  }

  for i := 0 ; i < len(indicators) ; i++ {
    err = imp.db.Where("refer_indicator = ?", indicators[i].Id).Find(&samples).Error

    if err != nil {
      return nil, err
    }

    indicators[i].Samples = samples
  }

  return indicators, nil
}

func (imp *IndicatorMapperPSQL) FindIndicatorById(id int) (Indicator, error) {
  var indicator Indicator
  var samples   []Sample

  err := imp.db.Where("id = ?", id).First(&indicator).Error

  if err != nil {
    return indicator, err
  }

  err = imp.db.Where("refer_indicator = ?", id).Find(&samples).Error

  for _, e := range samples {
    indicator.Samples = append(indicator.Samples, e)
  }

  return indicator, err
}

func (imp *IndicatorMapperPSQL) Insert(indicator *Indicator) error {
  return imp.db.Create(indicator).Error
}


func (imp *IndicatorMapperPSQL) Delete(id int) error {
  return imp.db.Delete(&Indicator{ Id: id }).Error
}


func (imp *IndicatorMapperPSQL) Update(indicator *Indicator) error {
  return imp.db.Save(indicator).Error
}

func (imp *IndicatorMapperPSQL) UpdatePushSample(sample *Sample) error {
  return imp.db.Create(sample).Error
}


func (imp *IndicatorMapperPSQL) UpdateSample(sample *Sample) error {
  return imp.db.Save(sample).Error
}
