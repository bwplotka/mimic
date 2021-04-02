// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

type errReader struct{ err error }

func (r errReader) Read(_ []byte) (int, error) { return 0, r.err }
