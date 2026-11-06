/**
 * bubbles.js — a tiny, zero-dependency animated bubble background.
 *
 * Drop-in usage for static sites (no build step):
 *
 *   <section id="hero"></section>
 *   <script src="bubbles.js"></script>
 *   <script>
 *     Bubbles.init("#hero", {
 *       count: 80,
 *       colors: ["#4bffbb", "#f6bf00", "#2a9dff"],
 *       speed: 1.2,
 *       background: "#001f23"
 *     });
 *   </script>
 *
 * Or with zero JS (auto-initialized on page load):
 *
 *   <section data-bubbles data-count="40" data-speed="0.6"></section>
 */
(function (global, factory) {
    if (typeof module === "object" && module !== null && module.exports) {
        module.exports = factory();
    } else {
        global.Bubbles = factory();
    }
}(typeof globalThis !== "undefined" ? globalThis : typeof window !== "undefined" ? window : this, function () {
    "use strict";

    const version = "1.0.0";

    const MAX_COUNT = 2000;
    const MAX_SPEED = 10;

    const defaults = {
        count: 60,
        colors: ["#6bff6b", "#f6bf00", "#f64f40", "#8e1cff", "#ff6ff0", "#2a9dff", "#4bffbb"],
        speed: 1,
        background: "#001f23"
    };

    const KNOWN_OPTIONS = ["count", "colors", "speed", "background"];

    const CONTAINER_CLASS = "bubble-bg";
    const BUBBLE_CLASS = "bubble-bg-bubble";
    const PAUSED_CLASS = "bubble-bg-paused";

    // Containers already rendered by an explicit init() call. Auto-init skips
    // these so it cannot overwrite the options the caller passed in.
    const initialized = typeof WeakSet === "function" ? new WeakSet() : null;

    // Containers with a live ResizeObserver. Kept here rather than as a property
    // on the element, which would leak library internals into the page's DOM.
    const observed = typeof WeakSet === "function" ? new WeakSet() : null;

    // Containers handed to the shared IntersectionObserver.
    const offscreen = typeof WeakSet === "function" ? new WeakSet() : null;

    // One observer for the whole document rather than one per container, and
    // held in this closure so it is never collected while containers still
    // point at it.
    let visibilityObserver = null;

    const cssText = `
.bubble-bg {
    --travel: 100vh;
    --travel: 100dvh;
    overflow: hidden;
    isolation: isolate;
}
.bubble-bg-bubble {
    position: absolute;
    z-index: -1;
    border-radius: 50%;
    width: var(--size);
    height: var(--size);
    left: var(--position-x);
    bottom: calc(-1 * var(--size) - 20px);
    opacity: var(--visibility);
    background-color: var(--color);
    pointer-events: none;
    animation: bubble-bg-float-up var(--anim-duration) linear infinite;
    animation-delay: var(--anim-delay);
}
@keyframes bubble-bg-float-up {
    0% {
        transform: translate3d(0, 0, 0);
    }
    25% {
        transform: translate3d(var(--distance), calc(var(--travel) * -0.28), 0);
    }
    50% {
        transform: translate3d(0, calc(var(--travel) * -0.56), 0);
    }
    75% {
        transform: translate3d(calc(-1 * var(--distance)), calc(var(--travel) * -0.84), 0);
    }
    100% {
        transform: translate3d(0, calc(var(--travel) * -1 - 80px), 0);
    }
}
@media (prefers-reduced-motion: reduce) {
    .bubble-bg-bubble {
        animation: none;
    }
}
.bubble-bg-paused .bubble-bg-bubble {
    animation-play-state: paused;
}
`.trim();

    function warn(message) {
        if (typeof console !== "undefined" && console.warn) {
            console.warn("[bubbles] " + message);
        }
    }

    function isSupported(property, value) {
        if (typeof CSS === "undefined" || typeof CSS.supports !== "function") {
            return true;
        }
        try {
            return CSS.supports(property, value);
        } catch (e) {
            return true;
        }
    }

    function ensureCss() {
        if (typeof document === "undefined") {
            return;
        }
        const existing = document.querySelector("style[data-bubble-bg]");
        if (existing && existing.getAttribute("data-bubble-bg") === version) {
            return;
        }
        // A sheet from a different version would style these class names with
        // stale rules; replace it so the attribute stays truthful.
        if (existing && existing.parentNode) {
            existing.parentNode.removeChild(existing);
        }
        const style = document.createElement("style");
        style.setAttribute("data-bubble-bg", version);
        style.textContent = cssText;
        (document.head || document.documentElement).appendChild(style);
    }

    function readData(element, name) {
        if (element.getAttribute) {
            return element.getAttribute("data-" + name);
        }
        return null;
    }

    // Returns undefined when the value is absent, and NaN when it is present but
    // not a number, so callers can tell "not specified" from "specified badly."
    // Number() rather than parseInt()/parseFloat(), which silently truncate:
    // parseInt("2.9") is 2 and parseInt("4e3") is 4.
    function parseNumber(value) {
        if (value === undefined || value === null) {
            return undefined;
        }
        if (typeof value === "number") {
            return value;
        }
        if (typeof value !== "string") {
            return NaN;
        }
        const trimmed = value.trim();
        if (!trimmed) {
            return undefined;
        }
        return Number(trimmed);
    }

    function parseColors(value) {
        if (Array.isArray(value)) {
            return value;
        }
        if (typeof value !== "string") {
            return undefined;
        }
        const list = splitColorList(value);
        return list.length ? list : undefined;
    }

    // Splits on top-level commas only, so functional color syntax survives:
    // "rgb(0,0,0),color-mix(in srgb, red, blue)" is two colors, not six.
    function splitColorList(value) {
        const parts = [];
        let depth = 0;
        let current = "";
        for (let i = 0; i < value.length; i++) {
            const character = value[i];
            if (character === "(") {
                depth++;
            } else if (character === ")") {
                depth = Math.max(0, depth - 1);
            }
            if (character === "," && depth === 0) {
                parts.push(current);
                current = "";
            } else {
                current += character;
            }
        }
        parts.push(current);
        return parts.map((part) => part.trim()).filter(Boolean);
    }

    // Applies one layer of configuration. A source that is present but
    // unparseable is ignored with a warning, leaving whatever the lower
    // precedence layer already set in place.
    function applySource(config, source, describe) {
        if (!source) {
            return;
        }
        KNOWN_OPTIONS.forEach((key) => {
            const value = source[key];
            if (value === undefined) {
                return;
            }
            if (typeof value === "number" && Number.isNaN(value)) {
                warn("Ignoring invalid " + describe(key) + " value.");
                return;
            }
            config[key] = value;
        });
    }

    function checkUnknownOptions(options) {
        if (!options || typeof options !== "object") {
            return;
        }
        Object.keys(options).forEach((key) => {
            if (!KNOWN_OPTIONS.includes(key)) {
                warn('Unknown option "' + key + '" ignored.');
            }
        });
    }

    function resolveConfig(element, options) {
        const config = {
            count: defaults.count,
            colors: defaults.colors.slice(),
            speed: defaults.speed,
            background: defaults.background
        };

        applySource(config, {
            count: parseNumber(readData(element, "count")),
            colors: parseColors(readData(element, "colors")),
            speed: parseNumber(readData(element, "speed")),
            background: readData(element, "background") || undefined
        }, (key) => "data-" + key);

        if (options) {
            checkUnknownOptions(options);
            applySource(config, {
                count: parseNumber(options.count),
                colors: parseColors(options.colors),
                speed: parseNumber(options.speed),
                background: options.background || undefined
            }, (key) => key + " option");
        }

        if (!Number.isInteger(config.count) || config.count < 0) {
            warn("Invalid count, using default " + defaults.count + ".");
            config.count = defaults.count;
        } else if (config.count > MAX_COUNT) {
            warn("Count capped at " + MAX_COUNT + ".");
            config.count = MAX_COUNT;
        } else if (config.count > 500) {
            warn("Large count (" + config.count + ") may hurt performance.");
        }

        if (!Array.isArray(config.colors) || !config.colors.length ||
                !config.colors.every((c) => typeof c === "string" && c)) {
            warn("Invalid colors, using the default palette.");
            config.colors = defaults.colors.slice();
        } else {
            const rejected = config.colors.filter((c) => !isSupported("color", c));
            if (rejected.length) {
                warn("Ignoring unsupported color(s): " + rejected.join(", ") + ".");
                config.colors = config.colors.filter((c) => isSupported("color", c));
            }
            if (!config.colors.length) {
                warn("No usable colors, using the default palette.");
                config.colors = defaults.colors.slice();
            }
        }

        if (typeof config.speed !== "number" || !Number.isFinite(config.speed) || config.speed <= 0) {
            warn("Invalid speed, using default " + defaults.speed + ".");
            config.speed = defaults.speed;
        } else if (config.speed > MAX_SPEED) {
            warn("Speed capped at " + MAX_SPEED + ".");
            config.speed = MAX_SPEED;
        }

        if (typeof config.background !== "string" || !config.background) {
            warn("Invalid background, using default.");
            config.background = defaults.background;
        } else if (!isSupported("background", config.background)) {
            warn('Unsupported background "' + config.background + '", using default.');
            config.background = defaults.background;
        }

        return config;
    }

    function isElement(value) {
        return !!value && value.nodeType === 1;
    }

    function resolveTargets(target) {
        if (typeof document === "undefined") {
            return [];
        }
        if (typeof target === "string") {
            return Array.from(document.querySelectorAll(target));
        }
        if (isElement(target)) {
            return [target];
        }
        if (target && typeof target.length === "number") {
            return Array.from(target).filter(isElement);
        }
        warn("Invalid target, expected a selector, an element, or a list of elements.");
        return [];
    }

    function clearBubbles(container) {
        container.querySelectorAll("." + BUBBLE_CLASS).forEach((node) => {
            node.remove();
        });
    }

    function updateTravel(container) {
        const height = container.clientHeight || 0;
        if (height > 0) {
            container.style.setProperty("--travel", height + "px");
        }
    }

    function watchSize(container) {
        if (typeof ResizeObserver === "undefined" || (observed && observed.has(container))) {
            return;
        }
        try {
            const observer = new ResizeObserver(() => updateTravel(container));
            observer.observe(container);
            if (observed) {
                observed.add(container);
            }
        } catch (e) {
            // ResizeObserver unavailable or blocked; travel stays at its init value.
        }
    }

    // Stops compositing work for containers scrolled out of view. A page with
    // many bubbles otherwise keeps every one of them animating offscreen.
    function watchVisibility(container) {
        if (typeof IntersectionObserver === "undefined" || (offscreen && offscreen.has(container))) {
            return;
        }
        try {
            if (!visibilityObserver) {
                visibilityObserver = new IntersectionObserver((entries) => {
                    entries.forEach((entry) => {
                        if (entry.target.classList) {
                            entry.target.classList.toggle(PAUSED_CLASS, !entry.isIntersecting);
                        }
                    });
                });
            }
            visibilityObserver.observe(container);
            if (offscreen) {
                offscreen.add(container);
            }
        } catch (e) {
            // IntersectionObserver unavailable or blocked; bubbles keep animating.
        }
    }

    function renderBubbles(container, config) {
        if (initialized) {
            initialized.add(container);
        }
        if (container.classList) {
            container.classList.add(CONTAINER_CLASS);
        }
        if (typeof getComputedStyle === "function" &&
                getComputedStyle(container).position === "static") {
            container.style.position = "relative";
        }
        container.style.background = config.background;
        clearBubbles(container);
        updateTravel(container);
        watchSize(container);
        watchVisibility(container);
        if (!container.clientHeight) {
            warn("Container has no height; give it content or a height to see bubbles.");
        }

        const fragment = document.createDocumentFragment();
        for (let i = 0; i < config.count; i++) {
            const bubble = document.createElement("span");
            if (bubble.classList) {
                bubble.classList.add(BUBBLE_CLASS);
            }
            bubble.setAttribute("aria-hidden", "true");

            const size = 4 + 14 * Math.random();
            const positionX = 100 * Math.random();
            const visibility = 0.08 + 0.22 * Math.random();
            const animDuration = (16 + 18 * Math.random()) / config.speed;
            const animDelay = -1 * (Math.random() * animDuration);
            const swayDirection = Math.random() > 0.5 ? 1 : -1;
            const distance = swayDirection * (16 + 28 * Math.random());

            bubble.style.setProperty("--size", size + "px");
            bubble.style.setProperty("--position-x", positionX + "%");
            bubble.style.setProperty("--visibility", String(visibility));
            bubble.style.setProperty("--anim-duration", animDuration + "s");
            bubble.style.setProperty("--anim-delay", animDelay + "s");
            bubble.style.setProperty("--distance", distance + "px");
            bubble.style.setProperty("--color", config.colors[i % config.colors.length]);

            fragment.appendChild(bubble);
        }
        container.appendChild(fragment);
    }

    function init(target, options) {
        if (typeof document === "undefined") {
            return [];
        }
        if (target === undefined || target === null) {
            target = "[data-bubbles]";
        } else if (!isElement(target) && typeof target !== "string" &&
                !(target && typeof target.length === "number")) {
            options = target;
            target = "[data-bubbles]";
        }
        const elements = resolveTargets(target);
        if (!elements.length) {
            if (typeof target === "string") {
                warn('No elements found for target "' + target + '".');
            } else {
                warn("No elements to initialize.");
            }
            return [];
        }
        ensureCss();
        elements.forEach((element) => {
            renderBubbles(element, resolveConfig(element, options));
        });
        return elements;
    }

    function autoInit() {
        if (typeof document === "undefined") {
            return [];
        }
        const elements = Array.from(document.querySelectorAll("[data-bubbles]"))
            .filter((element) => !initialized || !initialized.has(element));
        if (!elements.length) {
            return [];
        }
        return init(elements);
    }

    if (typeof document !== "undefined") {
        if (document.readyState === "loading") {
            document.addEventListener("DOMContentLoaded", autoInit);
        } else {
            autoInit();
        }
    }

    return {
        init: init,
        version: version
    };
}));
