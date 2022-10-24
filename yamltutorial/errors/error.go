package errors

const (
	RequiredFailCode = "required_field"
	PatternFailCode  = "invalid_pattern"
)

const (
	ErrManifestInvalid = "invalid_manifest"
	ErrManifestLoad    = "load_manifest"
	ErrManifestOpen    = "open_manifest"
	ErrManifestRead    = "read_manifest"
	ErrFailedMarshal   = "failed_marshal"
	ErrInvalidFile     = "invalid_file"
)

type Code struct {
	code  int
	field string
}
