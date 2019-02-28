# Usage

BTagChecker allows you to check if the tags used within the provided string has any unclosed tags. Great if someone modified some html and broke the header tag without your knowledge

# Example

```go
import "github.com/exts/BTagChecker"

func main() {
	if BTagChecker.HasValidClosingTags("<b>aye") {
		println("up")
	} else {
		println("oh no someone broke the html in your code")
	}
}
```