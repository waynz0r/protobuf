package typedeclimport

import subpkg "github.com/waynz0r/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
