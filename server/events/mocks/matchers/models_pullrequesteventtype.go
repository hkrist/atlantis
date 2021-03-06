package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsPullRequestEventType() models.PullRequestEventType {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.PullRequestEventType))(nil)).Elem()))
	var nullValue models.PullRequestEventType
	return nullValue
}

func EqModelsPullRequestEventType(value models.PullRequestEventType) models.PullRequestEventType {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.PullRequestEventType
	return nullValue
}
