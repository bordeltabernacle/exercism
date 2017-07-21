package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var r Resource

	for {
		r, err = o()
		if err != nil {
			if _, ok := err.(TransientError); !ok {
				return err
			}
		} else {
			break
		}
	}

	defer r.Close()

	defer func() {
		if x := recover(); x != nil {
			err = x.(error)

			if frobError, ok := x.(FrobError); ok {
				r.Defrob(frobError.defrobTag)
			}
		}
	}()

	r.Frob(input)
	return err
}
