package reportbundle

type Report struct {
	Id       int       `gorm:"AUTO_INCREMENT" json:"id"`
	Title    string    `gorm:"not null" json:"title"`
	Body     string    `gorm:"not null" json:"body"`
	Images   []string  `gorm:"not null" json:"images"`
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
  if len(r.Title) == 0 || len(r.Body) == 0 || len(r.Images) {
    return false
  }
  return true
}
