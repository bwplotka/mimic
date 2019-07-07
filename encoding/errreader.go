package encoding

type errReader struct{ err error }

func (r errReader) Read(_ []byte) (int, error) { return 0, r.err }
