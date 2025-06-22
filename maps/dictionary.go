package main

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrWordNotFound = DictionaryErr("could not find the word you are looking for")
	ErrWordExists   = DictionaryErr("word already defined")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
