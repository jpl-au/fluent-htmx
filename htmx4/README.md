# Fluent HTMX

An HTMX extension for [Fluent](https://github.com/jpl-au/fluent), targeting htmx 4. Wrap any Fluent element to add HTMX attributes through method chaining. Server-side helpers for handling HTMX requests and responses are also included.

## Install

```bash
go get github.com/jpl-au/fluent-htmx/htmx4
```

## Documentation for Agents

[`AGENTS.md`](AGENTS.md) - guide for coding agents. It is a self-contained superset of this README and the rules an agent should follow.

---

# Client-Side Attributes

Wrap any Fluent element with `htmx.New()` to add HTMX attributes. `New()` accepts `node.Element` - any HTML element created via Fluent's element packages (e.g. `div.New()`, `button.Text()`). Text nodes, function components, and conditionals are not elements and cannot be wrapped:

```go
package main

import (
    "net/http"

    "github.com/jpl-au/fluent/html5/button"
    "github.com/jpl-au/fluent/html5/div"
    "github.com/jpl-au/fluent-htmx/htmx4"
    "github.com/jpl-au/fluent-htmx/htmx4/swap"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        btn := button.Text("Load More")
        htmx.New(btn).
            HxGet("/api/items").
            HxTarget("#results").
            HxSwap(swap.BeforeEnd)

        div.New(btn).ID("results").Render(w)
    })
    http.ListenAndServe(":8080", mux)
}
```

The package is imported from `.../htmx4` but is named `htmx`, so call sites read `htmx.New(...)`.

## Core Attributes

```go
elem := div.New()
htmx.New(elem).
    HxGet("/api/data").
    HxPost("/api/submit").
    HxTarget("#result").
    HxSwap(swap.InnerHTML).
    HxTrigger("click").
    HxIndicator("#spinner").
    HxConfirm("Are you sure?")
```

| Method | Attribute | Description |
|--------|-----------|-------------|
| `HxGet(url)` | `hx-get` | Issue GET request |
| `HxPost(url)` | `hx-post` | Issue POST request |
| `HxPut(url)` | `hx-put` | Issue PUT request |
| `HxPatch(url)` | `hx-patch` | Issue PATCH request |
| `HxDelete(url)` | `hx-delete` | Issue DELETE request (excludes form data; add `HxInclude("closest form")` if needed) |
| `HxAction(url)` | `hx-action` | Request URL, paired with `HxMethod` |
| `HxMethod(method)` | `hx-method` | HTTP method for an `HxAction` URL |
| `HxTarget(selector, ...Mod)` | `hx-target` | Element to swap content into |
| `HxSwap(strategy, ...Mod)` | `hx-swap` | How to swap content |
| `HxTrigger(events, ...Mod)` | `hx-trigger` | What triggers the request |
| `HxBoost(bool, ...Mod)` | `hx-boost` | Progressive enhancement |
| `HxPushURL(value, ...Mod)` | `hx-push-url` | Push URL to history |
| `HxReplaceURL(url, ...Mod)` | `hx-replace-url` | Replace URL without a history entry |
| `HxSelect(selector, ...Mod)` | `hx-select` | Select content from response |
| `HxSelectOOB(selector, ...Mod)` | `hx-select-oob` | Out-of-band selection |
| `HxSwapOOB(value)` | `hx-swap-oob` | Out-of-band swaps |
| `HxVals(json, ...Mod)` | `hx-vals` | Add values to request (`js:` prefix for computed values) |
| `HxHeaders(json, ...Mod)` | `hx-headers` | Add headers to request |
| `HxInclude(selector, ...Mod)` | `hx-include` | Include additional values |
| `HxIndicator(selector, ...Mod)` | `hx-indicator` | Loading indicator element |
| `HxConfirm(message, ...Mod)` | `hx-confirm` | Confirmation prompt |
| `HxSync(strategy, ...Mod)` | `hx-sync` | Request synchronisation (and queuing) |
| `HxEncoding(encoding, ...Mod)` | `hx-encoding` | Request encoding (e.g. `multipart/form-data`) |
| `HxValidate(bool, ...Mod)` | `hx-validate` | Enable native form validation |
| `HxPreserve()` | `hx-preserve` | Preserve element during swap (presence-based) |
| `HxHistoryElt()` | `hx-history-elt` | Use this element as the history snapshot source |
| `HxIgnore()` | `hx-ignore` | Skip HTMX processing on this element and its children |
| `HxDisable(selector, ...Mod)` | `hx-disable` | Disable form elements while a request is in flight |
| `HxConfig(json, ...Mod)` | `hx-config` | Per-element request configuration (JSON or `key:value`) |
| `HxStatus(code, spec, ...Mod)` | `hx-status:CODE` | Per-status-code swap behaviour |
| `HxOn(event, handler)` | `hx-on:event` | Inline event handlers |

`HxSwap` accepts a `swap.Strategy` type. Use the predefined constants `swap.InnerHTML`, `swap.OuterHTML`, `swap.BeforeBegin`, `swap.AfterBegin`, `swap.BeforeEnd`, `swap.AfterEnd`, `swap.Delete`, `swap.None`, `swap.InnerMorph`, `swap.OuterMorph`, `swap.TextContent`, or `swap.Custom("innerHTML show:top")` for strategies with modifiers.

`HxSync` accepts a `sync.Strategy` type. Use the predefined constants `sync.Drop`, `sync.Abort`, `sync.Replace`, `sync.QueueFirst`, `sync.QueueLast`, `sync.QueueAll`, or `sync.Custom("this:queue all")` for element-scoped strategies.

## Common use cases

### Form submission, partial or full page

Post a form and swap the response in. The handler returns a partial for htmx requests and a full page otherwise.

```go
htmx.New(form).HxPost("/save").HxTarget("#content").HxSwap(swap.InnerHTML)

func HandleSave(w http.ResponseWriter, r *http.Request) {
    if htmx.Handle(r, func() { Partial().Render(w) }) {
        return
    }
    FullPage().Render(w)
}
```

### Active search (type-ahead)

Search as the user types, after a short pause.

```go
htmx.New(input).
    HxGet("/search").
    HxTrigger("keyup changed delay:300ms").
    HxTarget("#results").
    HxSwap(swap.InnerHTML)
```

## Attribute Inheritance

By default an attribute applies only to the element it is set on. The behavioural setters take an optional modifier that opts the attribute into inheritance down the DOM tree:

```go
// hx-confirm:inherited (every descendant inherits the confirmation)
htmx.New(div).HxConfirm("Are you sure?", htmx.Inherited)

// hx-include:inherited:append (descendants append to the inherited value)
htmx.New(form).HxInclude("#global-fields", htmx.InheritedAppend)
```

| Modifier | Suffix | Effect |
|----------|--------|--------|
| `htmx.Inherited` | `:inherited` | Attribute applies to the element and inherits to all descendants |
| `htmx.InheritedAppend` | `:inherited:append` | Inherits, but a descendant that also sets the attribute appends to the inherited value instead of replacing it |

The modifier is accepted by `HxTarget`, `HxSwap`, `HxTrigger`, `HxBoost`, `HxConfirm`, `HxVals`, `HxHeaders`, `HxIndicator`, `HxPushURL`, `HxReplaceURL`, `HxSelect`, `HxSelectOOB`, `HxInclude`, `HxSync`, `HxEncoding`, `HxValidate`, `HxDisable`, `HxConfig`, and `HxStatus`. Only `HxPreserve`, `HxSwapOOB`, `HxHistoryElt` and `HxIgnore` take no modifier - htmx reads those by presence or by id on a single element. The request verbs, `HxAction` and `HxMethod` also omit it, since an inherited verb cannot by itself initiate a request.

## Status-Code Swaps

`HxStatus` controls swap behaviour per HTTP status code. The code may be an exact status (`"404"`), a single-digit wildcard (`"50x"`), or a range wildcard (`"5xx"`); the spec uses htmx's `key:value` syntax with the keys `swap:`, `target:`, `select:`, `push:`, `replace:` and `transition:`.

```go
htmx.New(form).
    HxPost("/save").
    HxStatus("422", "swap:innerHTML target:#errors").
    HxStatus("5xx", "swap:none push:false")
```

## Extensions

htmx 4 ships two builds. `htmax.js` bundles eight popular extensions (ws, sse, preload, optimistic, targets, live, browser-indicator, download), so their attributes work once it is loaded. The core `htmx.js` build includes none of them; with it you load each extension's `dist/ext/<name>.js` script yourself, after htmx. htmx 4 has no `hx-ext`: loading the script, or the bundle, registers the extension page-wide, and the `extensions` config key can restrict which ones register.

A few extensions are never in `htmax.js` and always need their own script regardless of build: `hx-head`, `hx-ptag`, `hx-history-cache`, `hx-csp`, and the `upsert` swap.

If an extension is not loaded, its method has no effect. The binding writes the attribute, htmx does not recognise it, and core htmx still works. Each method below notes the script it needs.

### WebSocket

```go
htmx.New(elem).
    WsConnect("/ws/chat").
    WsSend()
```

`WsConnect` sets `hx-ws:connect`; `WsSend` sets `hx-ws:send`.

### Server-Sent Events

```go
htmx.New(elem).
    SSEConnect("/events").
    SSEClose("done")
```

`SSEConnect` sets `hx-sse:connect`; `SSEClose` sets `hx-sse:close`. There is no `sse-swap`: unnamed SSE messages are swapped automatically, and named events are dispatched as DOM events you handle with `HxTrigger`.

### Preload

```go
htmx.New(elem).
    HxGet("/page").
    Preload("mousedown")
```

`Preload` sets `hx-preload` to a trigger spec (e.g. `"mousedown"`, `"mouseover"`).

### Optimistic UI

```go
htmx.New(btn).HxPost("/save").HxTarget("#list").HxOptimistic("#pending")
```

`HxOptimistic` takes a CSS selector for an element whose HTML is shown in the target while the request runs.

### Multiple targets

```go
htmx.New(btn).HxGet("/refresh").HxTargets(".card")
```

`HxTargets` takes a CSS selector and swaps the response into every element that matches.

### Reactive expressions

```go
htmx.New(span).HxLive("q('#qty').value * unitPrice")
```

`HxLive` takes a JavaScript expression that recomputes when its inputs change.

### Browser indicator

```go
htmx.New(form).HxPost("/go").HxBrowserIndicator(true)
```

`HxBrowserIndicator` takes a bool, and shows the browser's native loading indicator during the request.

### Downloads

Downloads have no client attribute. Trigger one with the `swap.Download` swap style, or from the server with `HxDownload`:

```go
htmx.New(link).HxGet("/files/report.pdf").HxSwap(swap.Download)  // client: stream as a file
htmx.HxDownload(w, "/files/report.pdf")                          // server: sets HX-Download
```

The bundled download extension streams the response to the browser as a file and fires `htmx:download:start`, `:progress` and `:complete` events.

### Head Support

```go
htmx.New(elem).HxHead("merge")  // merge, append, or re-eval
```

`HxHead` sets `hx-head`, read by the htmx 4 **`hx-head` extension** (a separate script, `dist/ext/hx-head.js`, not bundled in `htmax.js`). Put `"merge"` or `"append"` on the response `<head>`, or `"re-eval"` on an individual head element. Without that extension loaded, core htmx swaps in the response title only and ignores the attribute.

### Other extensions (separate scripts)

These htmx 4 extensions are **not** bundled in `htmax.js` - load the matching `dist/ext/*.js` script for each, or the attribute is inert:

```go
htmx.New(div).HxGet("/news").HxTrigger("every 3s").HxPtag("v42")  // hx-ptag: skip swap if unchanged
htmx.New(div).HxHistory(false)                                    // hx-history: keep page out of the history cache
htmx.New(btn).HxPost("/save").HxNonce(nonce)                      // hx-nonce: CSP nonce (hx-csp extension)
htmx.New(btn).HxGet("/items").HxSwap(swap.Upsert)                 // upsert swap: update or insert by id
```

The `htmx-2-compat` (restore htmx 2 defaults) and `hx-alpine-compat` (run alongside Alpine.js) extensions add no per-element attribute. Load their scripts to use them.

---

# Server-Side Helpers

Since you're already using HTMX on the client, the package includes server-side helpers for handling requests and responses. These work with any Go HTTP framework.

## Detecting HTMX Requests

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if htmx.HxRequest(r) {
        // HTMX request: return partial HTML
        w.Write([]byte("<div>Updated content</div>"))
        return
    }
    // Non-HTMX request: return full page
    w.Write([]byte("<html>...</html>"))
}
```

Or use `Handle()` which executes a closure for HTMX requests and returns `true`:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if htmx.Handle(r, func() {
        w.Write([]byte("<div>Updated content</div>"))
        htmx.HxPushURL(w, "/new-path")
    }) {
        return
    }
    w.Write([]byte("<html>...</html>"))
}
```

## Reading Request Headers

```go
htmx.HxRequest(r)                   // Is this an HTMX request?
htmx.HxBoosted(r)                   // Was hx-boost used?
htmx.HxCurrentURL(r)                // URL the request was sent from
htmx.HxTarget(r)                    // Target element ("tagName#id")
htmx.HxSource(r)                    // Triggering element ("tagName#id")
htmx.HxRequestType(r)               // "full" or "partial"
htmx.HxPrompt(r)                    // HX-Prompt header; empty unless a prompt extension sets it
htmx.HxHistoryRestoreRequest(r)     // Is this a history restore?
```

## Setting Response Headers

```go
// Navigation
htmx.HxRedirect(w, r, "/login", http.StatusSeeOther)
htmx.HxPushURL(w, "/dashboard")
htmx.HxReplaceURL(w, "/dashboard")
htmx.HxLocation(w, "/page")
htmx.HxRefresh(w)

// Swap control
htmx.HxRetarget(w, "#other-element")
htmx.HxReswap(w, swap.OuterHTML)
htmx.HxReselect(w, ".content")
```

## Triggering Client Events

Trigger events fire as soon as the response is received, via the `HX-Trigger` header.

```go
// Simple event
htmx.NewTrigger(w).
    AddTrigger("itemAdded", nil).
    Write(div.Text("Item added"), http.StatusOK)

// Event with details
htmx.NewTrigger(w).
    AddTrigger("showMessage", map[string]string{
        "level": "success",
        "text":  "Item saved",
    }).
    Write(text.RawText(""), http.StatusOK)

// Multiple events
htmx.NewTrigger(w).
    AddTrigger("formReset", nil).
    AddTrigger("scrollToTop", nil).
    Write(formNode, http.StatusOK)
```

## Server-Sent Events

The package includes a server-side SSE writer that pairs with the client-side extension. It handles the SSE protocol, multi-line data, and response flushing:

```go
func eventsHandler(w http.ResponseWriter, r *http.Request) {
    sse, err := htmx.NewSSE(w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    sse.Send("message", "<div>New content</div>")
    sse.Send("done", "")  // closes the stream on the client's hx-sse:close event
}
```

`NewSSE` sets `Content-Type: text/event-stream`, `Cache-Control: no-cache`, and `Connection: keep-alive`. It returns an error if the ResponseWriter does not support flushing. Each `Send` call writes a named SSE event and flushes immediately.

---

# Configuration

Generate HTMX configuration meta tags:

```go
cfg := htmx.Config().
    DefaultSwapStyle(swap.OuterHTML).
    DefaultTimeout(5000).
    Transitions(true)

metaTag, err := cfg.ToMetaTag()
// <meta name="htmx-config" content='{"defaultSwap":"outerHTML","defaultTimeout":5000,...}'>
```

| Method | Default | Description |
|--------|---------|-------------|
| `DefaultSwapStyle(swap.Strategy)` | `"innerHTML"` | Default swap method (key `defaultSwap`) |
| `DefaultSettleDelay(ms)` | `1` | Delay before settle |
| `DefaultTimeout(ms)` | `60000` | Request timeout (0 disables) |
| `Transitions(bool)` | `false` | Use the View Transitions API |
| `HistoryEnabled(bool)` | `true` | Track history (key `history`) |
| `ImplicitInheritance(bool)` | `false` | Inherit attributes without the `:inherited` modifier |
| `NoSwap([]string)` | `["204","304"]` | Status codes that must not be swapped |
| `Mode(string)` | `"same-origin"` | Fetch mode (cross-origin behaviour) |
| `Extensions(string)` | `""` | Allow list of extension names |
| `MetaCharacter(string)` | `":"` | Separator in attribute/event names |
| `IncludeIndicatorCSS(bool)` | `true` | Inject default indicator CSS |
| `IndicatorClass(string)` | `"htmx-indicator"` | Loading indicator class |
| `RequestClass(string)` | `"htmx-request"` | Request in progress class |
| `InlineScriptNonce(string)` | `""` | CSP nonce for inline scripts |
| `DefaultFocusScroll(bool)` | `false` | Scroll focused element into view |
| `MorphIgnore([]string)` | `["data-htmx-powered"]` | Attribute prefixes preserved during morph |
| `MorphScanLimit(int)` | `10` | Max elements scanned during morph |
| `MorphSkip(string)` | `""` | Selector for elements to skip during morph |
| `MorphSkipChildren(string)` | `""` | Selector whose children to skip during morph |

`ToMetaTag()` and `ToJSON()` render the configuration.

---

# Constants

The package exports constants for HTMX headers. Swap strategies, CSS classes, and events are in their own sub-packages for cleaner call sites.

## Swap Strategies (`swap` package)

```go
swap.InnerHTML    // "innerHTML"
swap.OuterHTML    // "outerHTML"
swap.BeforeBegin  // "beforebegin"
swap.AfterBegin   // "afterbegin"
swap.BeforeEnd    // "beforeend"
swap.AfterEnd     // "afterend"
swap.Delete       // "delete"
swap.None         // "none"
swap.InnerMorph   // "innerMorph"
swap.OuterMorph   // "outerMorph"
swap.TextContent  // "textContent"
swap.Custom("innerHTML show:top showTarget:#other")
```

## Events (`event` package)

Event names follow the `htmx:phase:action` scheme, for use with `HxOn()` and JavaScript listeners.

```go
event.AfterSwap       // "htmx:after:swap"
event.BeforeRequest   // "htmx:before:request"
event.AfterRequest    // "htmx:after:request"
event.ConfigRequest   // "htmx:config:request"
event.ResponseError   // "htmx:response:error"
event.Error           // "htmx:error"
// ... and more, see event/event.go
```

## CSS Classes (`class` package)

```go
class.Request    // "htmx-request"
class.Indicator  // "htmx-indicator"
class.Added      // "htmx-added"
class.Settling   // "htmx-settling"
class.Swapping   // "htmx-swapping"
```

## Request Headers

```go
htmx.HXRequestHeader               // "HX-Request"
htmx.HXBoostedHeader               // "HX-Boosted"
htmx.HXCurrentURLHeader            // "HX-Current-URL"
htmx.HXHistoryRestoreRequestHeader // "HX-History-Restore-Request"
htmx.HXPromptHeader                // "HX-Prompt"
htmx.HXTargetHeader                // "HX-Target"
htmx.HXSourceHeader                // "HX-Source"
htmx.HXRequestTypeHeader           // "HX-Request-Type"
```

## Response Headers

```go
htmx.HXTriggerHeader         // "HX-Trigger"
htmx.HXLocationHeader        // "HX-Location"
htmx.HXPushURLHeader         // "HX-Push-Url"
htmx.HXRedirectHeader        // "HX-Redirect"
htmx.HXRefreshHeader         // "HX-Refresh"
htmx.HXReplaceURLHeader      // "HX-Replace-Url"
htmx.HXReswapHeader          // "HX-Reswap"
htmx.HXRetargetHeader        // "HX-Retarget"
htmx.HXReselectHeader        // "HX-Reselect"
htmx.HXDownloadHeader        // "HX-Download"
```

---

## Licence

MIT
