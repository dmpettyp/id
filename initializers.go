package id

import "fmt"

func Inititalizers[T any](
	createWrapper func(ID) T,
) (
	newID func() (T, error),
	mustNewID func() T,
	parseID func(s string) (T, error),
) {
	newID = func() (T, error) {
		id, err := New()

		wrappedID := createWrapper(id)

		if err != nil {
			return wrappedID, fmt.Errorf(
				"could not generate a %T: %w", wrappedID, err,
			)
		}

		return wrappedID, nil
	}

	mustNewID = func() T {
		return createWrapper(MustNew())
	}

	parseID = func(s string) (T, error) {
		id, err := Parse(s)

		wrappedID := createWrapper(id)

		if err != nil {
			return wrappedID, fmt.Errorf("could not parse a %T: %w", wrappedID, err)
		}

		return wrappedID, nil
	}

	return newID, mustNewID, parseID
}
