package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
    ErrWordDoesNotExist = DictionaryErr("cannt perform operation on word because it doesn't exist")
)

func (e DictionaryErr) Error() string {
    return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
    definition, ok := d[word]
    if !ok {
        return "", ErrNotFound
    }
    return definition, nil
}

func (d Dictionary) Add(key, value string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        d[key] = value
    case nil:
        return ErrWordExists
    default:
        return err
    }

    return nil
}

func (d Dictionary) Update(key, newValue string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[key] = newValue
    default:
        return err
    }

    return nil
}

func (d Dictionary) Delete(key string) error {
    _, err := d.Search(key)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        delete(d, key)
    default:
        return err
    }
    return nil
}
