package repository

import(
    "github.com/jmoiron/sqlx"
)

struct database type {
    conn *sqlx.DB
}

func connect(driverName string, dataSourceName string) (*sqlx.DB, error) {
    db, err := sqlx.Connect(driverName, dataSourceName)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return db, nil
}

func (db database) dbList(dest interface{}, query string) error {
    err := db.conn.Select(dest, query)
    if err != nil {
        return err
    }
    return nil
}

func (db database) dbGet(ID string, dest interface{}, query string) error {
    err := db.conn.Select(dest, query, ID)
    if err != nil {
        return err
    }
    return nil
}

func (db database) dbCreate(data interface{}, query string) error {
    err := db.conn.Query(query, data)
    if err != nil {
        return err
    }
    return nil
}

func (db database) dbUpdate(ID string, data interface{}, query string) error {
    err := db.conn.Query(query, ID, data)
    if err != nil {
        return err
    }
    return nil
}

func (db database) dbDelete(query string) error {
    err := db.conn.Query(query, ID)
    if err != nil {
        return err
    }
    return nil
}
