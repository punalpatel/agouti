package selection

import (
	"fmt"
	"github.com/onsi/gomega/format"
)

type HaveTextMatcher struct {
	ExpectedText string
	actualText   string
}

type Texter interface {
	Text() (string, error)
}

func (m *HaveTextMatcher) Match(actual interface{}) (success bool, err error) {
	actualSelection, ok := actual.(Texter)
	if !ok {
		return false, fmt.Errorf("HaveText matcher requires a Selection.  Got:\n%s", format.Object(actual, 1))
	}

	m.actualText, err = actualSelection.Text()
	if err != nil {
		return false, err
	}

	return m.actualText == m.ExpectedText, nil
}

func (m *HaveTextMatcher) FailureMessage(actual interface{}) (message string) {
	return selectorMessage(actual, "to have text matching", m.ExpectedText, m.actualText)

}

func (m *HaveTextMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return selectorMessage(actual, "not to have text matching", m.ExpectedText, m.actualText)
}