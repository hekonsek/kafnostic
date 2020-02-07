# Random strings

This library generates random strings.

## Usage

To generate human-friendly strings use `ForHuman*` functions.

```
import "github.com/hekonsek/random-strings"

var name string
name = randomstrings.ForHuman()                 // foobar
name = randomstrings.ForHumanWithHash()         // foobar6a3b21
name = randomstrings.ForHumanWithDash()         // foo-bar
name = randomstrings.ForHumanWithDashAndHash()  // foo-bar-6a3b21
```