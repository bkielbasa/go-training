# JSONs

## Complex data

Task: Fix TestComplexData test by writing a custom `UnmarshalJSON(data []byte) error` implementation for `PC` struct.

Based on the `type` the unmarshal should fill the `CPU` or `GPU` accordingly.
