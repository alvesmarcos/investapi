package indicatorbundle

import "github.com/jinzhu/copier"

type Sample struct {
  Id              int     `gorm:"AUTO_INCREMENT" json:"id"`
  ReferIndicator  int     `json:"referIndicator"`
  Date            string  `json:"date"`
  Value           float64 `json:"value"`
}

type Indicator struct {
  Id          int       `gorm:"AUTO_INCREMENT" json:"id"`
  Name        string    `gorm:"not null;unique" json:"name"`
  Description string    `gorm:"not null" json:"description"`
  Metric      string    `gorm:"not null" json:"metric"`
  Status      string    `gorm:"not null" json:"status"`
  Samples     []Sample	`json:"samples"`
}

func NewIndicator(name, description, metric, status string) *Indicator {
  return &Indicator { Name: name, Description: description, Metric: metric, Status: status }
}

func (indicator *Indicator) Copy(i *Indicator) {
  copier.Copy(indicator, i)
}

func (indicator Indicator) Validate() bool {
  if len(indicator.Name) == 0 || len(indicator.Description) == 0 || len(indicator.Metric) == 0 || len(indicator.Status) == 0 {
    return false
  }
  return true
}

func (indicator *Indicator) CompareAndSwap(i Indicator) {
  if indicator.Name != i.Name && len(i.Name) > 0 {
    indicator.Name = i.Name
  }
  if indicator.Description != i.Description && len(i.Description) > 0 {
    indicator.Description = i.Description
  }
  if indicator.Metric != i.Metric && len(i.Metric) > 0 {
    indicator.Metric = i.Metric
  }
  if indicator.Status != i.Status && len(i.Status) > 0 {
    indicator.Status = i.Status
  }
}

func (indicator *Indicator) PushSample(sample Sample) {
  indicator.Samples = append(indicator.Samples, sample)
}

func (indicator *Indicator) UpdateSamples(index int, sample Sample) {
  copier.Copy(indicator.Samples[index], sample)
}
