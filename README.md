# 🌀 Lucide icons for gomponents

This module provides the set of [Lucide Icons](https://lucide.dev/) for [gomponents](https://www.gomponents.com/).

## Features

- More than 1450 beautiful icons
- Easy to use
- Easy to customize
- Zero dependencies
- No client-side JavaScript

## Installation

```bash
go get github.com/eduardolat/gomponents-lucide
```

## Usage

You can [find your icons here](https://lucide.dev/icons/), convert the name to UpperCamelCase, and use it as a function that receives optional attributes to customize the SVG.

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
			html.Class("size-6 text-blue-500"),
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

## Customization

You can customize the SVG icons in two ways: individually and globally.

### Individually

You can use the `html.Class` or `html.Style` attributes to customize the SVG icons individually.

```go
lucide.Cherry(
	html.Class("size-6 text-blue-500"),
)
```

### Globally

**Every SVG in this module** includes the `data-glucide="icon"` attribute. You can use this attribute to customize all the icons globally using CSS:

```css
svg[data-glucide="icon"] {
  stroke-width: 4;
  stroke: red;
}
```

This approach ensures that any SVG icon with the `data-glucide="icon"` attribute will inherit the styles defined in your CSS, making it easy to maintain a consistent appearance across all icons in your project.

### Resolving conflicts with Tailwind CSS or custom styles

When using Tailwind CSS or custom styles for individual icon styling, you might encounter conflicts with global styles. To resolve this, you can modify your global styles to allow for Tailwind or custom styles overrides:

```css
svg[data-glucide="icon"]:not([class*="size-"]) {
	width: 16px;
	height: 16px;
}
```

This CSS rule applies a default size to all icons that don't have a specific `size-` class, allowing you to easily override the size using Tailwind or custom classes when needed:

```go
// This will use the default size (16px) defined in the global CSS
lucide.Cherry()

// This will override the default size with Tailwind's size class
lucide.Cherry(html.Class("size-32"))
```

This approach provides a flexible way to maintain consistent sizing across your project while allowing for easy customization of individual icons when necessary.

## Star and follow

If you like this project, please consider giving it a ⭐ on GitHub and [following me on X (Twitter)](https://twitter.com/eduardoolat).

## Versioning

This project increments its version independently of the Lucide Icons version. However, in each release, the Lucide Icons version is updated to the latest available.

You can see the Lucide Icons version in the notes of every release of this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

The Lucide icons are licensed under another open license. You can check it out [here](https://github.com/lucide-icons/lucide/blob/main/LICENSE).