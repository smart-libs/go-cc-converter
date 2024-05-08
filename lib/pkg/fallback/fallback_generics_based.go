package fallback

// FromFtoF creates a new conversion as `func(F, *F) error`
func FromFtoF[F any]() func(F, *F) error {
	return func(f F, pF *F) error {
		*pF = f
		return nil
	}
}

// FromFtoPointerToF creates a new conversion as `func(F, **F) error`
func FromFtoPointerToF[F any]() func(F, **F) error {
	return func(f F, pF **F) error {
		*pF = &f
		return nil
	}
}

// FromPointerToFtoF creates a new conversion as `func(F, **F) error`
func FromPointerToFtoF[F any]() func(*F, *F) error {
	return func(f *F, pF *F) error {
		var fZero F
		*pF = fZero
		if f != nil {
			*pF = *f
		}
		return nil
	}
}

// FromFtoArrayOfF creates a new conversion as `func(F, *[]F) error`
func FromFtoArrayOfF[F any]() func(F, *[]F) error {
	return func(f F, pF *[]F) error {
		*pF = []F{f}
		return nil
	}
}

// FromPointerToFtoArrayOfF creates a new conversion as `func(*F, *[]F) error`
func FromPointerToFtoArrayOfF[F any]() func(*F, *[]F) error {
	return func(f *F, pF *[]F) error {
		*pF = nil
		if f != nil {
			*pF = []F{*f}
		}
		return nil
	}
}

// PointerToPointerToT creates a new conversion as `func(F, **T) error` using `func(F, *T) error` as input
func PointerToPointerToT[F any, T any](baseConverter func(F, *T) error) func(F, **T) error {
	return func(f F, pT **T) error {
		*pT = nil
		var t T
		if err := baseConverter(f, &t); err != nil {
			return err
		}
		*pT = &t
		return nil
	}
}

// FromPointerToFToT creates a new conversion as `func(*F, *T) error` using `func(F, *T) error` as input
func FromPointerToFToT[F any, T any](baseConverter func(F, *T) error) func(*F, *T) error {
	return func(f *F, pT *T) error {
		var t T
		if f != nil {
			if err := baseConverter(*f, &t); err != nil {
				return err
			}
		}
		*pT = t
		return nil
	}
}

// FromPointerToFToPointerT creates a new conversion as `func(*F, *T) error` using `func(F, **T) error` as input
func FromPointerToFToPointerT[F any, T any](baseConverter func(F, *T) error) func(*F, **T) error {
	return func(f *F, pT **T) error {
		var t T
		*pT = nil
		if f != nil {
			if err := baseConverter(*f, &t); err != nil {
				return err
			}
			*pT = &t
		}
		return nil
	}
}

// FtoArrayOfT creates a new conversion as `func(F, *[]T) error` using `func(F, *T) error` as input
func FtoArrayOfT[F any, T any](baseConverter func(F, *T) error) func(F, *[]T) error {
	return func(f F, pT *[]T) error {
		*pT = nil
		var t T
		if err := baseConverter(f, &t); err != nil {
			return err
		}
		*pT = []T{t}
		return nil
	}
}

// PointerToFtoArrayOfT creates a new conversion as `func(F, *[]T) error` using `func(F, *T) error` as input
func PointerToFtoArrayOfT[F any, T any](baseConverter func(F, *T) error) func(*F, *[]T) error {
	fToArray := FtoArrayOfT(baseConverter)
	return func(f *F, pT *[]T) error {
		*pT = nil
		if f == nil {
			return nil
		}
		return fToArray(*f, pT)
	}
}

// ArrayOfFtoArrayOfT creates a new conversion as `func([]F, *[]T) error` using `func(F, *T) error` as input
func ArrayOfFtoArrayOfT[F any, T any](baseConverter func(F, *T) error) func([]F, *[]T) error {
	return func(arrayOfF []F, pT *[]T) error {
		*pT = nil
		for _, f := range arrayOfF {
			var t T
			if err := baseConverter(f, &t); err != nil {
				return err
			}
			*pT = append(*pT, t)
		}
		return nil
	}
}
