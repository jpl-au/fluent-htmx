# Fluent HTMX

An HTMX extension for [Fluent](https://github.com/jpl-au/fluent). Wrap any Fluent element to add HTMX attributes through method chaining. Server-side helpers for handling HTMX requests and responses are also included.

## Install

```bash
go get github.com/jpl-au/fluent-htmx
```

## Documentation for LLMs

- `AGENTS.md` - Comprehensive guide to help LLMs work with Fluent HTMX

---

# Client-Side Attributes

Wrap any Fluent element with `htmx.New()` to add HTMX attributes. `New()` accepts `node.Element` — any HTML element created via Fluent's element packages (e.g. `div.New()`, `button.Text()`). Text nodes, function components, and conditionals are not elements and cannot be wrapped:

```go
package main

import (
    "net/http"

    "github.com/jpl-au/fluent/html5/button"
    "github.com/jpl-au/fluent/html5/div"
    "github.com/jpl-au/fluent-htmx"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        btn := button.Text("Load More")
        htmx.New(btn).
            HxGet("/api/items").
            HxTarget("#results").
            HxSwap(htmx.SwapBeforeEnd)

        div.New(btn).ID("results").Render(w)
    })
    http.ListenAndServe(":8080", mux)
}
```

## Core Attributes

```go
elem := div.New()
htmx.New(elem).
    HxGet("/api/data").
    HxPost("/api/submit").
    HxPut("/api/update").
    HxPatch("/api/patch").
    HxDelete("/api/remove").
    HxTarget("#result").
    HxSwap(htmx.SwapInnerHTML).
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
| `HxDelete(url)` | `hx-delete` | Issue DELETE request |
| `HxTarget(selector)` | `hx-target` | Element to swap content into |
| `HxSwap(strategy)` | `hx-swap` | How to swap content |
| `HxTrigger(event)` | `hx-trigger` | What triggers the request |
| `HxBoost(bool)` | `hx-boost` | Progressive enhancement |
| `HxPushURL(value)` | `hx-push-url` | Push URL to history |
| `HxSelect(selector)` | `hx-select` | Select content from response |
| `HxSelectOOB(selector)` | `hx-select-oob` | Out-of-band selection |
| `HxSwapOOB(value)` | `hx-swap-oob` | Out-of-band swaps |
| `HxVals(json)` | `hx-vals` | Add values to request |
| `HxHeaders(json)` | `hx-headers` | Add headers to request |
| `HxParams(filter)` | `hx-params` | Filter request parameters |
| `HxInclude(selector)` | `hx-include` | Include additional values |
| `HxIndicator(selector)` | `hx-indicator` | Loading indicator element |
| `HxConfirm(message)` | `hx-confirm` | Confirmation prompt |
| `HxPrompt(message)` | `hx-prompt` | User input prompt |
| `HxValidate(bool)` | `hx-validate` | Enable validation |
| `HxSync(strategy)` | `hx-sync` | Request synchronisation |
| `HxPreserve(bool)` | `hx-preserve` | Preserve element during swap |
| `HxDisable()` | `hx-disable` | Disable HTMX processing |
| `HxDisabledElt(selector)` | `hx-disabled-elt` | Disable elements during request |
| `HxExt(extensions)` | `hx-ext` | Enable extensions |
| `HxOn(event, handler)` | `hx-on::event` | Inline event handlers |

`HxSwap` accepts a `SwapStrategy` type. Use the predefined constants `SwapInnerHTML`, `SwapOuterHTML`, `SwapBeforeBegin`, `SwapAfterBegin`, `SwapBeforeEnd`, `SwapAfterEnd`, `SwapDelete`, `SwapNone`, or `CustomSwap("innerHTML swap:1s")` for strategies with modifiers.

## Extensions

### WebSocket

```go
htmx.New(elem).
    HxExt("ws").
    WsConnect("/ws/chat").
    WsSend()
```

### Server-Sent Events

```go
htmx.New(elem).
    HxExt("sse").
    SSEConnect("/events").
    SSESwap("message").
    SSEClose("done")
```

### Preload

```go
htmx.New(elem).
    HxExt("preload").
    Preload("mousedown").
    PreloadImages(true)
```

### Response Targets

```go
htmx.New(elem).
    HxExt("response-targets").
    HxTargetError("#error-container").
    HxTargetCode(404, "#not-found").
    HxTargetCodePattern("5*", "#server-error")
```

### Head Support

```go
htmx.New(elem).
    HxExt("head-support").
    HxHead("merge")  // merge, append, or re-eval
```

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
htmx.HxTarget(r)                    // Target element ID
htmx.HxTrigger(r)                   // Triggering element ID
htmx.HxTriggerName(r)               // Triggering element name
htmx.HxPrompt(r)                    // User's response to hx-prompt
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
htmx.HxReswap(w, htmx.SwapOuterHTML)
htmx.HxReselect(w, ".content")
```

## Triggering Client Events

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

// Multiple events at different phases
htmx.NewTrigger(w).
    AddTrigger("formReset", nil).
    AddTriggerAfterSwap("focusInput", nil).
    AddTriggerAfterSettle("scrollToTop", nil).
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
    sse.Send("done", "")  // triggers sse-close on client
}
```

`NewSSE` sets `Content-Type: text/event-stream`, `Cache-Control: no-cache`, and `Connection: keep-alive`. It returns an error if the ResponseWriter does not support flushing. Each `Send` call writes a named SSE event and flushes immediately.

---

# Configuration

Generate HTMX configuration meta tags:

```go
cfg := htmx.Config().
    DefaultSwapStyle("outerHTML").
    Timeout(5000).
    GlobalViewTransitions(true).
    HistoryCacheSize(20)

metaTag, err := cfg.ToMetaTag()
// <meta name="htmx-config" content='{"defaultSwapStyle":"outerHTML","timeout":5000,...}'>
```

| Method | Default | Description |
|--------|---------|-------------|
| `DefaultSwapStyle(style)` | `"innerHTML"` | Default swap method |
| `DefaultSwapDelay(ms)` | `0` | Delay before swap |
| `DefaultSettleDelay(ms)` | `20` | Delay before settle |
| `Timeout(ms)` | `0` | Request timeout |
| `HistoryEnabled(bool)` | `true` | Enable history |
| `HistoryCacheSize(n)` | `10` | History cache entries |
| `RefreshOnHistoryMiss(bool)` | `false` | Full refresh on miss |
| `GlobalViewTransitions(bool)` | `false` | Use View Transitions API |
| `ScrollBehaviour(style)` | `"instant"` | Scroll animation |
| `IndicatorClass(class)` | `"htmx-indicator"` | Loading indicator class |
| `RequestClass(class)` | `"htmx-request"` | Request in progress class |
| `AllowEval(bool)` | `true` | Allow eval() |
| `AllowScriptTags(bool)` | `true` | Process script tags |
| `SelfRequestsOnly(bool)` | `true` | Same-domain requests only |
| `WithCredentials(bool)` | `false` | Cross-site credentials |

---

# Constants

The package exports constants for HTMX headers, CSS classes, and events.

## Request Headers

```go
htmx.HXRequestHeader         // "HX-Request"
htmx.HXBoostedHeader         // "HX-Boosted"
htmx.HXCurrentURLHeader      // "HX-Current-URL"
htmx.HXTargetHeader          // "HX-Target"
htmx.HXTriggerHeader         // "HX-Trigger"
htmx.HXTriggerNameHeader     // "HX-Trigger-Name"
htmx.HXPromptHeader          // "HX-Prompt"
```

## Response Headers

```go
htmx.HXLocationHeader        // "HX-Location"
htmx.HXPushURLHeader         // "HX-Push-Url"
htmx.HXRedirectHeader        // "HX-Redirect"
htmx.HXRefreshHeader         // "HX-Refresh"
htmx.HXReplaceURLHeader      // "HX-Replace-Url"
htmx.HXReswapHeader          // "HX-Reswap"
htmx.HXRetargetHeader        // "HX-Retarget"
htmx.HXReselectHeader        // "HX-Reselect"
```

## CSS Classes

```go
htmx.HXClassRequest    // "htmx-request"
htmx.HXClassIndicator  // "htmx-indicator"
htmx.HXClassAdded      // "htmx-added"
htmx.HXClassSettling   // "htmx-settling"
htmx.HXClassSwapping   // "htmx-swapping"
```

## Events (for HxOn)

```go
htmx.EventAfterSwap       // "afterSwap"
htmx.EventBeforeRequest   // "beforeRequest"
htmx.EventBeforeSwap      // "beforeSwap"
htmx.EventConfigRequest   // "configRequest"
htmx.EventLoad            // "load"
htmx.EventResponseError   // "responseError"
// ... and many more
```

---

## Profile-Guided Optimization (PGO)

Applications using Fluent HTMX benefit from [Profile-Guided Optimization](https://go.dev/doc/pgo) (Go 1.21+). PGO uses a CPU profile from your running application to make more aggressive inlining decisions at compile time. Expect **10-20% speed improvements** with no code changes.

1. Collect a CPU profile under realistic load:
   ```bash
   curl -o default.pgo http://localhost:8080/debug/pprof/profile?seconds=30
   ```
2. Place `default.pgo` in your main package directory
3. `go build` — PGO is applied automatically

Allocations are unaffected; PGO improves speed only.

## Licence

MIT
