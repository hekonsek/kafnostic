package randomstrings

import (
	"crypto/sha1"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"time"
)

// ForHuman returns "sillyname" (like "Pumacloud Lightning"), but without space and in lowercase
// i.e. "Pumacloud Lightning" becomes ("pumacloudlightning"). This function is useful for generating human-friendly
// random strings that should be used as semi-unique identifiers (as lowercase alphabetical-only string with
// reasonable size almost always can be used as valid identifier).
func ForHuman() string {
	return fmt.Sprintf("%s%s", randomdata.Adjective(), randomdata.Noun())
}

func ForHumanWithDash() string {
	return fmt.Sprintf("%s-%s", randomdata.Adjective(), randomdata.Noun())
}

func ForHumanWithHash() string {
	return ForHuman() + timestampHash()[:6]
}

func ForHumanWithDashAndHash() string {
	return ForHumanWithDash() + "-" + timestampHash()[:6]
}

// Private helpers

func timestampHash() string {
	hash := sha1.New()
	hash.Write([]byte(time.Now().Format(time.RFC850)))
	return fmt.Sprintf("%x\n", hash.Sum(nil))
}
