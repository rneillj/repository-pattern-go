package repository

type Repository interface {
    List    func(dest interface{}, query string, args ...interface{}) error
    Get     func(ID string, dest interface{}, query string, args ...interface{}) error
    Create  func(data interface{}, query string, args ...interface{}) error
    Update  func(ID string, data interface{}, query string, args ...interface{}) error
    Delete  func(ID string, query string, args ...interface{}) error
}

func List(dest interface{}, query string, args ...interface{}) error {
    err := db.Select(dest, query, args)
    if err != nil {
        return err
    }
    return nil
}

func Get(ID string, dest interface{}, query string, args ...interface{}) error {
    err := db.Select(ID, dest, query, args)
    if err != nil {
        return err
    }
    return nil
}

func Create(data interface{}, query string, args ...interface{}) error {
    err := db.Insert(ID, data, query, args)
    if err != nil {
        return err
    }
    return nil
}

func Update(ID string, data interface{}, query string, args ...interface{}) error {
    err := db.Update(ID, data, query, args)
    if err != nil {
        return err
    }
    return nil
}

func Delete(ID string, query string, args ...interface{}) error {
    err := db.Delete(ID, query, args)
    if err != nil {
        return err
    }
    return nil
}
