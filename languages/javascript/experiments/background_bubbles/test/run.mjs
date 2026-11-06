#!/usr/bin/env node
/**
 * Runs test/regression.html in headless Chrome and exits non-zero on failure.
 *
 *   node test/run.mjs
 *
 * No dependencies: Chrome is driven over the DevTools protocol using Node's
 * built-in fetch and WebSocket. Set CHROME=/path/to/chrome if the browser is
 * not in a standard location.
 */

import { spawn } from "node:child_process";
import { existsSync } from "node:fs";
import { mkdtemp, readFile, rm } from "node:fs/promises";
import { tmpdir } from "node:os";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const here = path.dirname(fileURLToPath(import.meta.url));
const fixture = pathToFileURL(path.join(here, "regression.html")).href;

const CANDIDATES = [
    process.env.CHROME,
    "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
    "/Applications/Chromium.app/Contents/MacOS/Chromium",
    "/usr/bin/google-chrome",
    "/usr/bin/chromium",
    "/usr/bin/chromium-browser"
].filter(Boolean);

const chromePath = CANDIDATES.find((candidate) => existsSync(candidate));

if (!chromePath) {
    console.error("No Chrome or Chromium binary found.");
    console.error("Set CHROME=/path/to/chrome and run again.");
    process.exit(2);
}

function fail(message) {
    console.error("run.mjs: " + message);
    process.exit(1);
}

const profile = await mkdtemp(path.join(tmpdir(), "bubbles-test-"));
const browser = spawn(chromePath, [
    "--headless=new",
    "--disable-gpu",
    "--no-first-run",
    "--no-default-browser-check",
    "--allow-file-access-from-files",
    "--remote-debugging-port=0",
    "--user-data-dir=" + profile,
    "about:blank"
], { stdio: "ignore" });

let socket;
let exitCode = 1;

try {
    // Chrome writes the chosen port to DevToolsActivePort once it is listening.
    const portFile = path.join(profile, "DevToolsActivePort");
    const deadline = Date.now() + 15000;
    let port;
    while (Date.now() < deadline) {
        if (existsSync(portFile)) {
            const lines = (await readFile(portFile, "utf8")).split("\n");
            if (lines[0] && lines[0].trim()) {
                port = Number(lines[0].trim());
                break;
            }
        }
        await new Promise((resolve) => setTimeout(resolve, 100));
    }
    if (!port) {
        fail("Chrome did not open a devtools port within 15s.");
    }

    const target = await (await fetch("http://127.0.0.1:" + port + "/json/new?about:blank", {
        method: "PUT"
    })).json();

    socket = new WebSocket(target.webSocketDebuggerUrl);
    let nextId = 0;
    const pending = new Map();

    socket.addEventListener("message", (event) => {
        const message = JSON.parse(event.data);
        if (message.id && pending.has(message.id)) {
            pending.get(message.id)(message);
            pending.delete(message.id);
        }
    });

    await new Promise((resolve, reject) => {
        socket.addEventListener("open", resolve, { once: true });
        socket.addEventListener("error", () => reject(new Error("DevTools socket error")), { once: true });
    });

    function send(method, params) {
        return new Promise((resolve) => {
            const id = ++nextId;
            pending.set(id, resolve);
            socket.send(JSON.stringify({ id: id, method: method, params: params || {} }));
        });
    }

    async function evaluate(expression) {
        const reply = await send("Runtime.evaluate", {
            expression: expression,
            returnByValue: true,
            awaitPromise: true
        });
        if (reply.error || !reply.result) {
            fail("Runtime.evaluate failed: " + JSON.stringify(reply.error || {}));
        }
        const thrown = reply.result.exceptionDetails;
        if (thrown) {
            fail("Page threw: " + (thrown.exception && thrown.exception.description || thrown.text));
        }
        return reply.result.result.value;
    }

    await send("Runtime.enable");
    await send("Page.enable");
    // Headless Chrome can report "reduce" on its own; pin the preference so the
    // reduced-motion media block never changes what these assertions observe.
    await send("Emulation.setEmulatedMedia", {
        features: [{ name: "prefers-reduced-motion", value: "no-preference" }]
    });

    await send("Page.navigate", { url: fixture });

    const ready = await (async () => {
        const stop = Date.now() + 15000;
        while (Date.now() < stop) {
            if (await evaluate("window.__done === true")) {
                return true;
            }
            await new Promise((resolve) => setTimeout(resolve, 100));
        }
        return false;
    })();

    if (!ready) {
        fail("The test page never finished (window.__done stayed false).");
    }

    const pageError = await evaluate("window.__error");
    const results = await evaluate("JSON.stringify(window.__results)").then(JSON.parse);

    let failed = 0;
    for (const result of results) {
        if (result.pass) {
            console.log("  PASS  " + result.name);
        } else {
            failed++;
            console.log("  FAIL  " + result.name + (result.detail ? "  (" + result.detail + ")" : ""));
        }
    }

    if (pageError) {
        console.log("\nPage error:\n" + pageError);
    }

    console.log("\n" + (results.length - failed) + "/" + results.length + " passed");

    if (failed === 0 && !pageError && results.length > 0) {
        exitCode = 0;
    }
} catch (error) {
    console.error("run.mjs: " + (error && error.stack ? error.stack : error));
} finally {
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.close();
    }
    // Wait for the browser to actually exit before deleting its profile,
    // otherwise Chrome is still writing files into it when rm runs.
    await new Promise((resolve) => {
        if (browser.exitCode !== null || browser.signalCode !== null) {
            return resolve();
        }
        const force = setTimeout(() => {
            if (browser.exitCode === null) {
                browser.kill("SIGKILL");
            }
        }, 5000);
        force.unref();
        browser.once("exit", () => {
            clearTimeout(force);
            resolve();
        });
        browser.kill("SIGTERM");
    });
    await rm(profile, { recursive: true, force: true, maxRetries: 5, retryDelay: 100 }).catch(() => {});
}

process.exit(exitCode);
