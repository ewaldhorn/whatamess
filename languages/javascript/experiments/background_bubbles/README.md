# bubbles.js

A tiny, zero-dependency animated bubble background for static sites. One file, no build step: the script injects its own CSS, so there is nothing else to include or keep in sync.

## Demo

Open `dist/index.html` in a browser. It shows JS init with a custom config, a zero-JS `data-` attribute section, and a live playground.

## Quick start

1. Copy `bubbles.js` next to your page.
2. Add a container and initialize it:

```html
<section id="hero"></section>
<script src="bubbles.js"></script>
<script>
  Bubbles.init("#hero", {
    count: 80,
    colors: ["#4bffbb", "#f6bf00", "#2a9dff"],
    speed: 1.2,
    background: "#001f23"
  });
</script>
```

Or skip the JS entirely — elements with `data-bubbles` are initialized automatically on page load:

```html
<section data-bubbles data-count="40" data-speed="0.6"
         data-background="#241031" data-colors="#ff6ff0,#8e1cff"></section>
```

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `count` | number | `60` | How many bubbles to render. A non-negative integer, capped at 2000. |
| `colors` | string array | 7-color palette | Bubble colors, cycled in order. As a data attribute, comma-separated; functional syntax like `rgb(0,0,0)` or `color-mix(in srgb, red, blue)` works. |
| `speed` | number | `1` | Speed multiplier: `2` is twice as fast, `0.5` is half speed. Must be above `0`, capped at `10`. |
| `background` | string | `"#001f23"` | Container background. Any CSS color works; gradients work too. |

Precedence: options passed to `init()` beat `data-` attributes, which beat the defaults. Values are validated strictly — `count` must be a non-negative integer, `speed` a finite number above `0`. Anything that cannot be parsed logs a console warning and is ignored, so the next source down (or the default) applies; unknown option names warn as well, which catches typos like `colour`.

## API

- `Bubbles.init(target?, options?)` — `target` is a selector string, an element, or a list of elements. Omit it (or pass only an options object) to initialize every `[data-bubbles]` element. Safe to call again on the same container: old bubbles are replaced. Returns the initialized elements.
- `Bubbles.version` — the library version string.

Content placed inside the container renders above the bubbles automatically, and bubbles never intercept clicks, so no extra layering markup is needed.

## Tests

`test/regression.html` holds the assertions; `test/run.mjs` runs them in headless Chrome with no dependencies:

```sh
node test/run.mjs
```

It checks that bubble travel follows the container height, that it tracks container resizes, and that auto-init does not overwrite options passed to `init()`. Set `CHROME=/path/to/chrome` if the browser is not in a standard location.

## Notes

- Containers need their own height (page content or CSS) — bubbles are absolutely positioned and don't create layout. Initializing a zero-height container logs a warning.
- Respects `prefers-reduced-motion`: the animation switches off and the static background remains.
- Bubble travel distance follows the container height and tracks resizes, so it works in short sections as well as full-viewport heroes.
- No dependencies; works from any static host or straight from `file://`.
- Needs a modern browser (CSS custom properties). Browsers without `dvh` fall back to `vh`, and without `ResizeObserver` the travel distance is measured once at init.
