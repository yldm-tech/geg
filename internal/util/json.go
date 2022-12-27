/**
 * @Author: Evan<suzukaze.hazuki2020@gmail.com>
 * @Description:
 * @File:  json_util
 * @Version: 1.0.0
 * @Date: 2021/11/29 17:15
 */

package util

import (
	"bytes"
	"encoding/json"
)

// Marshal is a UTF-8 friendly marshaller.  Go's json.Marshal is not UTF-8
// friendly because it replaces the valid UTF-8 and JSON characters "&". "<",
// ">" with the "slash u" unicode escaped forms (e.g. \u0026).  It preemptively
// escapes for HTML friendliness.  Where text may include any of these
// characters, json.Marshal should not be used. Playground of Go breaking a
// title: https://play.golang.org/p/o2hiX0c62oN
func Marshal(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "	")
	err := encoder.Encode(i)
	return bytes.TrimRight(buffer.Bytes(), "\n"), err
}
