package validator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fdefabricio/crawler-novelas/model"
)

func Check(n model.Novela) (es []error) {
	es = make([]error, 0)

	es = append(es, actors(n.Actors)...)
	es = append(es, authors(n.Authors)...)
	es = append(es, chapters(n.Chapters))
	es = append(es, directors(n.Directors)...)
	es = append(es, hour(n.Hour))
	es = append(es, name(n.Name))
	es = append(es, url(n.URL))
	es = append(es, year(n.Year))

	return
}

func actors(as []string) []error {
	return slice(as, "actors")
}

func authors(as []string) []error {
	return slice(as, "authors")
}

func chapters(s string) error {
	if err := empty(s, "chapters"); err != nil {
		return err
	}

	n, err := integer(s, "chapters")
	if err != nil {
		return err
	}

	if n <= 0 {
		return errors.New(fmt.Sprintf("number os chapters is invalid: %d", n))
	}

	return nil
}

func directors(as []string) []error {
	return slice(as, "directors")
}

func hour(s string) error {
	if err := empty(s, "hour"); err != nil {
		return err
	}

	h, err := integer(s, "hour")
	if err != nil {
		return err
	}

	if h < 6 || h > 11 {
		return errors.New(fmt.Sprintf("hour is invalid: %d", h))
	}

	return nil
}

func name(s string) error {
	return empty(s, "name")
}

func url(s string) error {
	if err := empty(s, "url"); err != nil {
		return err
	}

	if !strings.Contains(s, "pt.wikipedia.org") {
		return errors.New("url is from a different domain")
	}

	return nil
}

func year(s string) error {
	if err := empty(s, "year"); err != nil {
		return err
	}

	y, err := integer(s, "year")
	if err != nil {
		return err
	}

	if y < 1900 || y > time.Now().Year() {
		return errors.New(fmt.Sprintf("year is invalid: %d", y))
	}

	return nil
}

func duplicate(slice []string, fieldName string) []error {
	es := make([]error, 0)
	for i, s := range slice {
		if isIn(slice[i+1:], s) {
			es = append(es, errors.New(fmt.Sprintf("%s duplicate entry: %s", fieldName, s)))
		}
	}

	return es
}

func empty(s string, fieldName string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return errors.New(fmt.Sprintf("%s is empty", fieldName))
	}

	return nil
}

func integer(s string, fieldName string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, errors.New(fmt.Sprintf(fmt.Sprintf("%s is not a number: %s", fieldName, s)))
	}

	return i, nil
}

func isIn(slice []string, b string) bool {
	for _, a := range slice {
		if a == b {
			return true
		}
	}
	return false
}

func slice(as []string, fieldName string) []error {
	if err := unfilled(as, fieldName); err != nil {
		return []error{err}
	}

	es := make([]error, 0)

	for _, a := range as {
		if err := empty(a, fieldName); err != nil {
			es = append(es, err)
		}
	}

	es = append(es, duplicate(as, fieldName)...)

	return es
}

func unfilled(a []string, fieldName string) error {
	if len(a) == 0 {
		return errors.New(fmt.Sprintf("%s has zero entries", fieldName))
	}

	return nil
}
