// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

/*
encoding is a package of default encodings supplied with mimic.

Each of the encodings match the following func declaration:
```
func XXX(in interface{}) io.Reader {
```

Encodings are called when generating the desired output for a configuration.

Additional encoders should return `io.Reader` and can be imported directly into the configuration.
*/
package encoding
