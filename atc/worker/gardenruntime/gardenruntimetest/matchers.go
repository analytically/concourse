package gardenruntimetest

import (
	"errors"

	"github.com/concourse/concourse/atc/runtime/runtimetest"
	"github.com/concourse/concourse/worker/baggageclaim"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

func HaveStrategy(strategy baggageclaim.Strategy) types.GomegaMatcher {
	return haveStrategyMatcher{strategy}
}

type haveStrategyMatcher struct {
	expected baggageclaim.Strategy
}

func (m haveStrategyMatcher) Match(actual any) (bool, error) {
	volume, ok := actual.(*Volume)
	if !ok {
		return false, errors.New("expecting a *grt.Volume")
	}

	return StrategyEq(m.expected)(volume), nil
}
func (m haveStrategyMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to have strategy", m.expected)
}
func (m haveStrategyMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to have strategy", m.expected)
}

func HaveContent(content runtimetest.VolumeContent) types.GomegaMatcher {
	return haveContentMatcher{content}
}

type haveContentMatcher struct {
	expected runtimetest.VolumeContent
}

func (m haveContentMatcher) Match(actual any) (bool, error) {
	volume, ok := actual.(*Volume)
	if !ok {
		return false, errors.New("expecting a *grt.Volume")
	}

	return ContentEq(m.expected)(volume), nil
}
func (m haveContentMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to have content", m.expected)
}
func (m haveContentMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to have content", m.expected)
}

func BePrivileged() types.GomegaMatcher {
	return bePrivilegedMatcher{true}
}

type bePrivilegedMatcher struct {
	expected bool
}

func (m bePrivilegedMatcher) Match(actual any) (bool, error) {
	volume, ok := actual.(*Volume)
	if !ok {
		return false, errors.New("expecting a *grt.Volume")
	}

	return PrivilegedEq(m.expected)(volume), nil
}
func (m bePrivilegedMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to be "+m.expectation())
}
func (m bePrivilegedMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to be "+m.expectation())
}

func (m bePrivilegedMatcher) expectation() string {
	if m.expected {
		return "privileged"
	}
	return "unprivileged"
}

func HaveHandle(handle string) types.GomegaMatcher {
	return haveHandleMatcher{handle}
}

type haveHandleMatcher struct {
	expected string
}

func (m haveHandleMatcher) Match(actual any) (bool, error) {
	volume, ok := actual.(*Volume)
	if !ok {
		return false, errors.New("expecting a *grt.Volume")
	}

	return HandleEq(m.expected)(volume), nil
}
func (m haveHandleMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to have handle "+m.expected)
}
func (m haveHandleMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to have handle "+m.expected)
}
