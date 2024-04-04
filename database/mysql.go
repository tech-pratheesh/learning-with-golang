//go:build mysql

package database

type Mysql struct{}

func (m *Mysql) GetUsers() ([]string, error) {
	return []string{
		"mysql-user-01",
		"mysql-user-02",
	}, nil
}

func init() {
	operations = &Mysql{}
}
