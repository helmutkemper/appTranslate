package translate

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type Entry struct {
	Page    string
	Tag     string
	Key     string
	Message interface{}
}

type Entries []Entry

func New(translationList Entries) error {
	var err error

	for _, entry := range translationList {
		tag := language.MustParse(entry.Tag)
		switch msg := entry.Message.(type) {
		case string:
			err = message.SetString(tag, entry.Page+"."+entry.Key, msg)
			if err != nil {
				return err
			}
		case catalog.Message:
			err = message.Set(tag, entry.Page+"."+entry.Key, msg)
			if err != nil {
				return err
			}
		case []catalog.Message:
			err = message.Set(tag, entry.Page+"."+entry.Key, msg...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
