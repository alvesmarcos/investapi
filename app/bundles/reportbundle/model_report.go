package reportbundle

// Report struct
type Report struct {
	Id     uint32		`json:"id"`
	Title  string		`json:"title"`
	Body   string		`json:"body"`
	Images []string	`json:"images"`
}

// NewReport create a new Report
func NewReport(id uint32, title, body string, images []string) *Report {
  return &Report { Id: id, Title: title, Body: body, Images: images }
}
