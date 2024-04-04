//go:build psql

package database

type Psql struct{}

func (m *Psql) GetUsers() ([]string, error) {
	return []string{
		"psql-user-01",
		"psql-user-02",
	}, nil
}

func init() {
	operations = &Psql{}
}
