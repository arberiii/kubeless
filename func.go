package kubeless

import (
	"fmt"

	"github.com/kubeless/kubeless/pkg/functions"
)

// SayHello greets the caller of this function!
func SayHello(event functions.Event, context functions.Context) (string, error) {
	return fmt.Sprintf("Hello %s", event.Data), nil
}
