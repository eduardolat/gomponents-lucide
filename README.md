# ✏️ Lucide icons for gomponents

This module provides the set of [Lucide Icons](https://lucide.dev/) for [gomponents](https://www.gomponents.com/).

## Features

- More than 1400 beautiful icons
- Easy to use
- Easy to customize
- Zero dependencies
- No client-side JavaScript

## Installation

```bash
go get github.com/eduardolat/gomponents-lucide
```

## Usage

You can [find your icons here](https://lucide.dev/icons/) convert the name to UpperCamelCase and use it as a function that receives optional attributes to customize the SVG.

Your code editor should help you with the autocompletion of the name of the functions.

Here is an example:

```go
package main

import (
	"os"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	"github.com/eduardolat/gomponents-lucide"
)

func myPage() gomponents.Node {
	return html.Div(
		lucide.CircleUser(),
		lucide.ChevronUp(),
		lucide.Power(),
		lucide.Star(),
		lucide.Languages(),
		lucide.Usb(),
		//...
		lucide.Cherry(
			// You can add any attributes you want
			// to customize the SVG
			html.Class("w-6 h-6 text-blue-500"),
		),
	)
}

func main() {
	// This prints the HTML to stdout but you can
	// write it to whatever io.Writer you want
	page := myPage()
	page.Render(os.Stdout)
}
```

## Star and follow

If you like this project, please consider giving it a ⭐ on GitHub and [following me on X (Twitter)](https://twitter.com/eduardoolat).

## Versioning

This project increments its version independently of the Lucide Icons version. However, in each release, the Lucide Icons version is updated to the latest available.

You can see the Lucide Icons version in the notes of every release of this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
