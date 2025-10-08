package navigation

import "maragu.dev/gomponents"

type Side struct{}

func (s Side) Node() gomponents.Node {
	return gomponents.El("lens-main-navigation",
		gomponents.Map(entries, func(e Entry) gomponents.Node { return e.Node() }),
	)
}
