package rm

import (
	"database/sql"
)

type Connector struct {
	db 			*sql.DB
	driver 		string
	url 		string
}


func NewConnector(driver string, url string) *Connector {
	connector := new(Connector)
	connector.db, _ = sql.Open(driver, url)

	return connector
}

func (c *Connector)Insert(sql string) (bool, error){
	_, e := c.db.Exec(sql)
	if e != nil {
		return false, e
	}
	return true, nil
}
