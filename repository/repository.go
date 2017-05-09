package repository

type Repository interface {
    List    func(dest interface{}, query string) error
    Get     func(ID string, dest interface{}, query string) error
    Create  func(data interface{}, query string) error
    Update  func(ID string, data interface{}, query string) error
    Delete  func(ID string, query string) error
}

func List(dest interface{}, query string) error {
    err := dbList(dest, query)
    if err != nil {
        return err
    }
    return nil
}

func Get(ID string, dest interface{}, query string) error {
    err := dbGet(ID, dest, query)
    if err != nil {
        return err
    }
    return nil
}

func Create(data interface{}, query string) error {
    err := dbCreate(data, query)
    if err != nil {
        return err
    }
    return nil
}

func Update(ID string, data interface{}, query string) error {
    err := dbUpdate(ID, data, query)
    if err != nil {
        return err
    }
    return nil
}

func Delete(ID string, query string) error {
    err := dbDelete(ID, query)
    if err != nil {
        return err
    }
    return nil
}
