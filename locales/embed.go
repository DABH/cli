// Package locales holds translation catalogs maintained by Localizer (https://github.com/DABH/localizer).
package locales

import "embed"

// FS contains one "<language>.json" catalog per language.
//
//go:embed *.json
var FS embed.FS
