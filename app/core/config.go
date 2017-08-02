package core

type Config struct {
  DBConnection  string
  DBTYPE        string
}

func (c *Config) Fetch() {
  c.DBConnection = "host=localhost user=postgres dbname=sda sslmode=disable password=admin123"
  c.DBTYPE = "postgres"
}
