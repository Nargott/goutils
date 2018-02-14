package goutils

import (
    "crypto/sha256"
    "encoding/base64"
)

/**
    Returns a Base64-encoded sha256 hash of given string
 */
func HashSHA256Base64(data string) string {
    h := sha256.New()
    h.Write([]byte(data))
    return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
