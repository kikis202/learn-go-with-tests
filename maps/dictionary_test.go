package main

import "testing"

var key = "test"
var value = "this is just a test"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{key: value}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search(key)
		want := value

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrWordNotFound)

	})
}

func TestAdd(t *testing.T) {

	t.Run("add new value", func(t *testing.T) {

		dictionary := Dictionary{}
		dictionary.Add(key, value)

		want := value
		got, err := dictionary.Search(key)
		if err != nil {
			t.Fatal("should find added word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("add existing value", func(t *testing.T) {

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")
		if err == nil {
			t.Fatal("should fail redefining a word:")
		}

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	newValue := "new test"

	t.Run("update existing value", func(t *testing.T) {

		dictionary := Dictionary{key: value}
		err := dictionary.Update(key, newValue)
		if err != nil {
			t.Fatal("should update existing word:")
		}

		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("update non-existing value", func(t *testing.T) {

		dictionary := Dictionary{}
		err := dictionary.Update(key, newValue)
		if err == nil {
			t.Fatal("shouldnt update non-existing words definition")
		}

		assertDefinitionError(t, dictionary, key)
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete non-existing value", func(t *testing.T) {

		dictionary := Dictionary{}
		err := dictionary.Delete(key)
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("deltete existing value", func(t *testing.T) {

		dictionary := Dictionary{key: value}
		err := dictionary.Delete(key)

		assertError(t, err, nil)
		assertDefinitionError(t, dictionary, key)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("got unexpected error:", err)
	}

	if got != value {
		t.Errorf("got %q want %q", got, value)
	}
}

func assertDefinitionError(t testing.TB, dictionary Dictionary, key string) {
	t.Helper()

	_, err := dictionary.Search(key)

	if err == nil {
		t.Fatal("expected to not find the key", err)
	}
}
