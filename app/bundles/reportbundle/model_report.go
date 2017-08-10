package reportbundle

import "github.com/alvesmarcos/investapi/app/core"

type Report struct {
  Id       int              `gorm:"AUTO_INCREMENT" json:"id"`
  Title    string           `gorm:"not null" json:"title"`
  Body     string           `gorm:"not null" json:"body"`
  Images   core.StringSlice `gorm:"type:text[];not null" json:"images"`
}

func NewReport(title, body string, images []string) *Report {
  return &Report { Id: 0, Title: title, Body: body, Images: images }
}

func (r *Report) Copy(report *Report) {
  r.Title = report.Title
  r.Body = report.Body
  r.Images = report.Images
}

func (r Report) Validate() bool {
  if len(r.Title) == 0 || len(r.Body) == 0 || len(r.Images) == 0 {
    return false
  }
  return true
}

func (r *Report) CompareAndSwap(report Report) {
  if r.Title != report.Title && len(report.Title) > 0 {
    r.Title = report.Title
  }
  if r.Body != report.Body && len(report.Body) > 0 {
    r.Body = report.Body
  }
}

func (r *Report) PushImage(path string) {
  r.Images = append(r.Images, path)
}

func (r *Report) UpdateImages(index int, path string) {
  r.Images[index] = path
}
