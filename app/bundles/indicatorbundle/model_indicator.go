package indicatorbundle

import "github.com/jinzhu/copier"

type Sample struct {
  Id          int     `json:"id"`
  IndicatorId int     `gorm:"index" json:"indicatorId"`
  Date        string  `json:"date"`
  Value	      string  `json:"value"`
}

type Indicator struct {
  Id          int       `gorm:"AUTO_INCREMENT" json:"id"`
  Name        string    `gorm:"not null" json:"name"`
  Description string    `gorm:"not null" json:"description"`
  Metric      string    `gorm:"not null" json:"metric"`
  Status      string    `gorm:"not null" json:"status"`
  Samples     []Sample	`json:"samples"`
}

func NewIndicator(id uint32, name, description, metric, status string, samples []Samples) *Indicator {
  return &Indicator { Id: id, Name: name, Description: description, Metric: metric, Samples: samples, Status: status }
}

func (indicator *Indicator) Copy(i *Indicator) {
  copier.Copy(indicator, i)
}

func (indicator Indicator) Validate() bool {
  if len(name) == 0 || len(description) == 0 || len(metric) == 0 || len(status) == 0 {
    return false
  }
  return true
}
