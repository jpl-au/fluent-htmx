# HTMX Extension Guide for LLMs

## Methods That Do Not Exist

The following methods **have never existed** in this package. Do not use them:

- `Attr()` — does not exist. Use the typed `Hx*()` methods listed below.
- `SetAttr()` — does not exist.
- `Attribute()` — does not exist.
- `Data()` — does not exist. Fluent's `node.Element` has `SetData()`, but it is not available on the HTMX `*Wrapper`.
- `Aria()` — does not exist. Fluent's `node.Element` has `SetAria()`, but it is not available on the HTMX `*Wrapper`.
- `SetData()` — does not exist on `*Wrapper`. It exists on the underlying `node.Element`.
- `SetAria()` — does not exist on `*Wrapper`. It exists on the underlying `node.Element`.

If you need `SetData()` or `SetAria()`, call them on the Fluent element **before** wrapping it with `htmx.New()`.

## Embedded Assets (embed.go, handler.go, script.go)

The package embeds htmx 2.0.8 and supported extensions via `go:embed`. No CDN or manual downloads required.

### Serving

```go
mux.Handle("/_htmx/", htmx.Handler("/_htmx/"))
```

### Script Helpers

| Function | Output |
|----------|--------|
| `Script(prefix string)` | `<script src="{prefix}htmx.min.js"></script>` |
| `ExtScript(prefix, name string)` | `<script src="{prefix}ext/{name}.js"></script>` |

Both return `node.Node` so they compose directly in fluent trees.

### Raw Access

| Function | Returns |
|----------|---------|
| `Assets()` | `fs.FS` rooted at `dist/` |
| `Version` | `"2.0.8"` (constant) |

### Bundled Extensions

`ws`, `sse`, `preload`, `response-targets`, `head-support` — matching the extension helpers in this package.

## Architecture

`htmx.New(element)` wraps a Fluent `node.Element` and returns `*Wrapper`. The Wrapper delegates these `node.Element` methods to the underlying element: `Render`, `RenderBuilder`, `RenderOpen`, `RenderClose`, `Nodes`, `SetAttribute`. All other methods on `*Wrapper` are the HTMX-specific methods listed in this document.

`SetAttribute(key, value)` is exposed on Wrapper as a pass-through to the underlying element. **Never call `SetAttribute` directly for HTMX attributes** — always use the typed `Hx*()` methods instead.

## Constraints

- **NEVER use `.SetAttribute()` for HTMX attributes** — always use `htmx.New(element).HxX()` methods
- **NEVER use `r.Header.Get("HX-Request")`** — use `htmx.HxRequest(r)`
- **NEVER use `w.Header().Set("HX-*")`** — use `htmx.HxX(w, ...)` functions
- **NEVER add global JavaScript event listeners for HTMX** — use `HxOn()` for locality of behaviour
- If a method is not listed in this document, it does not exist

## Pattern

`htmx.New()` accepts `node.Element` — any HTML element created via Fluent's element packages. Text nodes, function components, and conditionals are `node.Node` only and cannot be wrapped.

```go
// Client-side: wrap element, chain HTMX methods
htmx.New(element).HxPost("/api/endpoint").HxTarget("#result").HxOn("after-swap", "console.log('done')")

// Server-side: use helper functions
if htmx.HxRequest(r) { /* partial */ } else { /* full page */ }
htmx.HxPushURL(w, "/new-url")
```

## Complete Client Method Reference

This is the **exhaustive** list of methods on `*Wrapper`. If a method is not listed here, it does not exist.

### HTTP Verbs (client.go)

| Method | Attribute |
|--------|-----------|
| `HxGet(url string)` | `hx-get` |
| `HxPost(url string)` | `hx-post` |
| `HxPut(url string)` | `hx-put` |
| `HxPatch(url string)` | `hx-patch` |
| `HxDelete(url string)` | `hx-delete` |

### Swap & Targeting (client.go)

| Method | Attribute |
|--------|-----------|
| `HxSwap(strategy string)` | `hx-swap` |
| `HxTarget(selector string)` | `hx-target` |
| `HxSwapOOB(value string)` | `hx-swap-oob` |
| `HxSelect(selector string)` | `hx-select` |
| `HxSelectOOB(selector string)` | `hx-select-oob` |

### Triggers & Events (client.go)

| Method | Attribute |
|--------|-----------|
| `HxTrigger(events string)` | `hx-trigger` |
| `HxOn(event string, handler string)` | `hx-on::event` |

Use constants from `constants.go` for event names: `EventAfterSwap`, `EventBeforeSwap`, `EventAfterSettle`, `EventBeforeRequest`, `EventAfterRequest`, `EventConfigRequest`, etc.

### Boolean Attributes (client.go)

| Method | Attribute |
|--------|-----------|
| `HxBoost(enabled bool)` | `hx-boost` |
| `HxPreserve(preserve bool)` | `hx-preserve` |
| `HxValidate(validate bool)` | `hx-validate` |

### URL Management (client.go)

| Method | Attribute |
|--------|-----------|
| `HxPushURL(value string)` | `hx-push-url` |
| `HxReplaceURL(url string)` | `hx-replace-url` |

### Form & Request Parameters (client.go)

| Method | Attribute |
|--------|-----------|
| `HxVals(json string)` | `hx-vals` |
| `HxHeaders(json string)` | `hx-headers` |
| `HxParams(params string)` | `hx-params` |
| `HxInclude(selector string)` | `hx-include` |
| `HxEncoding(encoding string)` | `hx-encoding` |
| `HxConfirm(message string)` | `hx-confirm` |
| `HxPrompt(message string)` | `hx-prompt` |

### Control Flow (client.go)

| Method | Attribute |
|--------|-----------|
| `HxIndicator(selector string)` | `hx-indicator` |
| `HxSync(strategy string)` | `hx-sync` |
| `HxDisabledElt(selector string)` | `hx-disabled-elt` |
| `HxDisable()` | `hx-disable` |
| `HxHistoryElt()` | `hx-history-elt` |

### Inheritance & History (client.go)

| Method | Attribute |
|--------|-----------|
| `HxExt(extensions string)` | `hx-ext` |
| `HxHistory(value string)` | `hx-history` |
| `HxDisinherit(attributes string)` | `hx-disinherit` |
| `HxInherit(attributes string)` | `hx-inherit` |
| `HxRequest(config string)` | `hx-request` |

### WebSocket Extension (ws.go)

| Method | Attribute |
|--------|-----------|
| `WsConnect(url string)` | `ws-connect` |
| `WsSend()` | `ws-send` |

### Server-Sent Events Extension (sse.go)

| Method | Attribute |
|--------|-----------|
| `SSEConnect(url string)` | `sse-connect` |
| `SSESwap(eventName string)` | `sse-swap` |
| `SSEClose(eventName string)` | `sse-close` |

### Preload Extension (preload.go)

| Method | Attribute |
|--------|-----------|
| `Preload(trigger string)` | `preload` |
| `PreloadImages(enabled bool)` | `preload-images` |

### Response Targets Extension (response_targets.go)

| Method | Attribute |
|--------|-----------|
| `HxTargetError(selector string)` | `hx-target-error` |
| `HxTargetCode(code int, selector string)` | `hx-target-{code}` |
| `HxTargetCodePattern(pattern string, selector string)` | `hx-target-{pattern}` |

### Head Support Extension (head_support.go)

| Method | Attribute |
|--------|-----------|
| `HxHead(mode string)` | `hx-head` |

### Deprecated (client.go)

| Method | Attribute | Note |
|--------|-----------|------|
| `HxVars(variables string)` | `hx-vars` | Use `HxVals` instead |

## Server Functions (server.go)

### Request Detection

| Function | Returns |
|----------|---------|
| `HxRequest(r *http.Request)` | `bool` — true if HTMX request |
| `Handle(r *http.Request, fn func())` | `bool` — executes fn if HTMX request, returns true |

### Read Request Headers

| Function | Returns |
|----------|---------|
| `HxBoosted(r)` | `bool` |
| `HxCurrentURL(r)` | `string` |
| `HxHistoryRestoreRequest(r)` | `bool` |
| `HxTarget(r)` | `string` |
| `HxTrigger(r)` | `string` |
| `HxTriggerName(r)` | `string` |
| `HxPrompt(r)` | `string` |

### Write Response Headers

| Function | Parameters |
|----------|------------|
| `HxPushURL(w, url)` | Push URL to browser history |
| `HxReplaceURL(w, url)` | Replace URL without history entry |
| `HxRedirect(w, r, url, code)` | Client-side redirect (HTMX) or HTTP redirect (standard) |
| `HxLocation(w, url)` | Client-side redirect without full reload |
| `HxRefresh(w)` | Full page refresh |
| `HxRetarget(w, selector)` | Override swap target |
| `HxReswap(w, strategy)` | Override swap strategy |
| `HxReselect(w, selector)` | Override response selection |

### Trigger Events

```go
trigger := htmx.NewTrigger(w)
trigger.AddTrigger(eventName, detailMap)           // immediate
trigger.AddTriggerAfterSwap(eventName, detailMap)  // after swap
trigger.AddTriggerAfterSettle(eventName, detailMap) // after settle
trigger.Write(content, statusCode)
```

### Simple Response

```go
htmx.Response(w, "<div>content</div>", http.StatusOK)
```

## Configuration (config.go)

`htmx.Config()` creates a builder for generating HTMX `<meta>` configuration tags. All methods return `*config` for chaining.

```go
cfg := htmx.Config().
    DefaultSwapStyle("outerHTML").
    Timeout(5000).
    GlobalViewTransitions(true)

metaTag, err := cfg.ToMetaTag()
// <meta name="htmx-config" content='{"defaultSwapStyle":"outerHTML","timeout":5000,...}'>

jsonStr, err := cfg.ToJSON()
```

### Config Methods

| Method | Default | Description |
|--------|---------|-------------|
| `DefaultSwapStyle(string)` | `"innerHTML"` | Default swap method |
| `DefaultSwapDelay(int)` | `0` | Delay in ms before swap |
| `DefaultSettleDelay(int)` | `20` | Delay in ms before settle |
| `Timeout(int)` | `0` | Request timeout in ms |
| `HistoryEnabled(bool)` | `true` | Enable history snapshots |
| `HistoryCacheSize(int)` | `10` | Max cached history pages |
| `RefreshOnHistoryMiss(bool)` | `false` | Full refresh on cache miss |
| `HistoryRestoreAsHxRequest(bool)` | `true` | Send HX-Request on history restore |
| `GlobalViewTransitions(bool)` | `false` | Use View Transitions API |
| `ScrollBehaviour(string)` | `"instant"` | Scroll animation style |
| `ScrollBehavior(string)` | — | American spelling alias |
| `DefaultFocusScroll(bool)` | `false` | Scroll focused element into view |
| `ScrollIntoViewOnBoost(bool)` | `true` | Scroll on boosted navigation |
| `IndicatorClass(string)` | `"htmx-indicator"` | Loading indicator CSS class |
| `RequestClass(string)` | `"htmx-request"` | Request-in-progress CSS class |
| `AddedClass(string)` | `"htmx-added"` | Newly added content CSS class |
| `SettlingClass(string)` | `"htmx-settling"` | Settling phase CSS class |
| `SwappingClass(string)` | `"htmx-swapping"` | Swapping phase CSS class |
| `IncludeIndicatorStyles(bool)` | `true` | Inject default indicator CSS |
| `AllowEval(bool)` | `true` | Allow eval() |
| `AllowScriptTags(bool)` | `true` | Execute scripts in swapped content |
| `InlineScriptNonce(string)` | `""` | CSP nonce for inline scripts |
| `InlineStyleNonce(string)` | `""` | CSP nonce for inline styles |
| `AttributesToSettle([]string)` | `["class","style"]` | Attributes updated during settle |
| `SelfRequestsOnly(bool)` | `true` | Restrict to same-domain requests |
| `WithCredentials(bool)` | `false` | Cross-origin credentials |
| `GetCacheBusterParam(bool)` | `false` | Append cache-buster to GET |
| `IgnoreTitle(bool)` | `false` | Prevent title updates from swaps |
| `DisableSelector(string)` | `"[hx-disable]..."` | Selector for disabled elements |
| `DisableInheritance(bool)` | `false` | Prevent attribute inheritance |
| `WsReconnectDelay(string)` | `"full-jitter"` | WebSocket reconnect strategy |
| `WsBinaryType(string)` | `"blob"` | WebSocket binary data type |
| `MethodsThatUseURLParams([]string)` | `["get"]` | Methods using URL query params |
| `ReportValidityOfForms(bool)` | `false` | Call reportValidity() before submit |
| `AllowNestedOobSwaps(bool)` | `true` | Process nested OOB swaps |
| `TriggerSpecsCache(interface{})` | — | Pre-populated trigger spec cache |
| `ResponseHandling(interface{})` | — | Custom response handling rules |
| `ToMetaTag()` | — | Returns `(string, error)` |
| `ToJSON()` | — | Returns `(string, error)` |

## Usage Patterns

### Form Submission

```go
htmx.New(form).HxPost("/save").HxTarget("#content").HxSwap("innerHTML")

func HandleSave(w http.ResponseWriter, r *http.Request) {
    if htmx.Handle(r, func() {
        ViewPartial().Render(w)
        htmx.HxPushURL(w, "/new-path")
    }) { return }
    http.Redirect(w, r, "/new-path", http.StatusSeeOther)
}
```

### List Updates

```go
htmx.New(form).HxPost("/add").HxTarget("#list").HxSwap("afterbegin")
// Swap strategies: innerHTML, outerHTML, beforebegin, afterbegin, beforeend, afterend
```

### Inline Event Handler (Locality of Behaviour)

```go
handler := `document.querySelectorAll('.item').forEach(el => el.classList.remove('active'));
event.target.closest('.item').classList.add('active');`
htmx.New(div).HxOn("after-swap", handler)
```

### Delete with Confirmation

```go
htmx.New(btn).HxDelete("/items/"+id).HxConfirm("Sure?").HxTarget("closest .item").HxSwap("outerHTML")
```

## Profile-Guided Optimization (PGO)

Applications using Fluent HTMX benefit from [PGO](https://go.dev/doc/pgo) (Go 1.21+). Collect a CPU profile from production, place it as `default.pgo` in the main package, and `go build` applies it automatically. Expect 10-20% speed improvements with no code changes. Allocations are unaffected — PGO improves inlining decisions only.
