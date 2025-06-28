# thai2karaoke

Package `thai2karaoke` provides a simple converter from Thai text to Latin karaoke using the Royal Thai General System (RTGS). The public API exposes a single `Translate` function.

```
import "github.com/comdevx/thai2karaoke"

latin := thai2karaoke.Translate("สวัสดีครับ")
// latin == "Sawatdi Khrap"
```
