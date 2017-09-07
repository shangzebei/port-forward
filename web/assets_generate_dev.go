// +build !vfs

package web
import "net/http"

// Assets contains project assets.
var assets http.FileSystem = http.Dir("./asset")
