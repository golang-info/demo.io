package main

import "errors"

type survey struct {
	id int
	name string
	feedback string
}

type formHandler struct {
	surveys []survey
}

func (f *formHandler) addFeedback(name, feedback string) int {
	total := len(f.surveys)
	id := total + 1
	newSurvey := survey{
		id: id,
		name: name,
		feedback: feedback,
	}

	f.surveys = append(f.surveys, newSurvey)
	return id
}

func (f *formHandler) getFeedback(id int) (survey, error) {
	for _, s := range f.surveys {
		if s.id == id {
			return s, nil
		}
	}
	return survey{}, errors.New("Survey not found")
}
