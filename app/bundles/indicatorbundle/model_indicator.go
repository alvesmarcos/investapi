package indicatorbundle

type Sample struct {
  Id    uint32  `json:"id"`
  Date	string	`json:"date"`
	Value	string	`json:"value"`
}

type Indicator struct {
  Id          uint32		`json:"id"`
	Name				string		`json:"name"`
	Description	string		`json:"description"`
	Metric			string 		`json:"metric"`
	Status			string 		`json:"status"`
	Date				string 		`json:"date"`
	Samples			[]Sample	`json:"samples"`
}

func NewIndicator(id uint32, name, description, metric, status, date string, samples []Samples) *Indicator {
  return &Indicator { Id: id, Name: name, Description: description, Metric: metric, Status: status, Date: date, Samples: samples }
}
