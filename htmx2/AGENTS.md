# HTMX Extension Guide for Agents

## Methods That Do Not Exist

The following methods **have never existed** in this package. Do not use them:

- `Attr()` - does not exist. Use the typed `Hx*()` methods listed below.
- `SetAttr()` - does not exist.
- `Attribute()` - does not exist.
- `Data()` - does not exist. Fluent's `node.Element` has `SetData()`, but it is not available on the HTMX `*Wrapper`.
- `Aria()` - does not exist. Fluent's `node.Element` has `SetAria()`, but it is not available on the HTMX `*Wrapper`.
- `SetData()` - does not exist on `*Wrapper`. It exists on the underlying `node.Element`.
- `SetAria()` - does not exist on `*Wrapper`. It exists on the underlying `node.Element`.

If you need `SetData()` or `SetAria()`, call them on the Fluent element **before** wrapping it with `htmx.New()`.

## Architecture

`htmx.New(element)` wraps a Fluent `node.Element` and returns `*Wrapper`. The Wrapper delegates these `node.Element` methods to the underlying element: `Render`, `WriteTo`, `RenderBytes`, `RenderBuilder`, `RenderOpen`, `RenderClose`, `Nodes`, `SetAttribute`, `SetAttributeRaw`. (`SetAttributeRaw` is the trusted-value raw hatch, mirroring `RawText`: it stores the value verbatim, whereas `SetAttribute` escapes it.) All other methods on `*Wrapper` are the HTMX-specific methods listed in this document.

`SetAttribute(key, value)` is exposed on Wrapper as a pass-through to the underlying element. **Never call `SetAttribute` directly for HTMX attributes** - always use the typed `Hx*()` methods instead.

## Constraints

- **NEVER use `.SetAttribute()` for HTMX attributes** - always use `htmx.New(element).HxX()` methods
- **NEVER use `r.Header.Get("HX-Request")`** - use `htmx.HxRequest(r)`
- **NEVER use `w.Header().Set("HX-*")`** - use `htmx.HxX(w, ...)` functions
- **NEVER add global JavaScript event listeners for HTMX** - use `HxOn()` for locality of behaviour
- If a method is not listed in this document, it does not exist

## Pattern

`htmx.New()` accepts `node.Element` - any HTML element created via Fluent's element packages. Text nodes, function components, and conditionals are `node.Node` only and cannot be wrapped.

```go
// Client-side: wrap element, chain HTMX methods
htmx.New(element).HxPost("/api/endpoint").HxTarget("#result").HxOn("after-swap", "console.log('done')")

// Server-side: use helper functions
if htmx.HxRequest(r) { /* partial */ } else { /* full page */ }
htmx.HxPushURL(w, "/new-url")
```

## Sub-Packages

Swap strategies, events, and CSS classes live in their own packages for cleaner call sites.

### Swap Strategies (`swap` package)

`swap.Strategy` is a typed string used by `HxSwap()`, `HxReswap()`, and `DefaultSwapStyle()`.

| Constant | Value |
|----------|-------|
| `swap.InnerHTML` | `"innerHTML"` |
| `swap.OuterHTML` | `"outerHTML"` |
| `swap.BeforeBegin` | `"beforebegin"` |
| `swap.AfterBegin` | `"afterbegin"` |
| `swap.BeforeEnd` | `"beforeend"` |
| `swap.AfterEnd` | `"afterend"` |
| `swap.Delete` | `"delete"` |
| `swap.None` | `"none"` |

`swap.Custom(strategy string) swap.Strategy` - creates a strategy with modifiers, e.g. `swap.Custom("innerHTML swap:1s")`.

### Sync Strategies (`sync` package)

`sync.Strategy` is a typed string used by `HxSync()`.

| Constant | Value |
|----------|-------|
| `sync.Drop` | `"drop"` |
| `sync.Abort` | `"abort"` |
| `sync.Replace` | `"replace"` |
| `sync.QueueFirst` | `"queue first"` |
| `sync.QueueLast` | `"queue last"` |
| `sync.QueueAll` | `"queue all"` |

`sync.Custom(strategy string) sync.Strategy` - creates a strategy with element selectors, e.g. `sync.Custom("closest form:abort")`.

### Events (`event` package)

Event constants for use with `HxOn()`. Examples:

| Constant | Value |
|----------|-------|
| `event.AfterSwap` | `"afterSwap"` |
| `event.BeforeRequest` | `"beforeRequest"` |
| `event.BeforeSwap` | `"beforeSwap"` |
| `event.ConfigRequest` | `"configRequest"` |
| `event.Load` | `"load"` |
| `event.ResponseError` | `"responseError"` |
| `event.HistoryCacheHit` | `"historyCacheHit"` |
| `event.SSEOpen` | `"sseOpen"` |
| `event.ValidationFailed` | `"validation:failed"` |
| `event.XHRProgress` | `"xhr:progress"` |

See `event/event.go` for the full list.

### CSS Classes (`class` package)

| Constant | Value |
|----------|-------|
| `class.Added` | `"htmx-added"` |
| `class.Indicator` | `"htmx-indicator"` |
| `class.Request` | `"htmx-request"` |
| `class.Settling` | `"htmx-settling"` |
| `class.Swapping` | `"htmx-swapping"` |

## Complete Client Method Reference

This is the **exhaustive** list of methods on `*Wrapper`. If a method is not listed here, it does not exist.

### HTTP Verbs (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxGet(url string)` | `hx-get` |
| `HxPost(url string)` | `hx-post` |
| `HxPut(url string)` | `hx-put` |
| `HxPatch(url string)` | `hx-patch` |
| `HxDelete(url string)` | `hx-delete` |

### Swap & Targeting (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxSwap(strategy swap.Strategy)` | `hx-swap` |
| `HxTarget(selector string)` | `hx-target` |
| `HxSwapOOB(value string)` | `hx-swap-oob` |
| `HxSelect(selector string)` | `hx-select` |
| `HxSelectOOB(selector string)` | `hx-select-oob` |

### Triggers & Events (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxTrigger(events string)` | `hx-trigger` |
| `HxOn(event string, handler string)` | `hx-on::event` |

Use constants from the `event` package for event names: `event.AfterSwap`, `event.BeforeSwap`, `event.AfterSettle`, `event.BeforeRequest`, `event.AfterRequest`, `event.ConfigRequest`, etc.

### Boolean Attributes (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxBoost(enabled bool)` | `hx-boost` |
| `HxPreserve(preserve bool)` | `hx-preserve` |
| `HxValidate(validate bool)` | `hx-validate` |

### URL Management (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxPushURL(value string)` | `hx-push-url` |
| `HxReplaceURL(url string)` | `hx-replace-url` |

### Form & Request Parameters (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxVals(json string)` | `hx-vals` |
| `HxHeaders(json string)` | `hx-headers` |
| `HxParams(params string)` | `hx-params` |
| `HxInclude(selector string)` | `hx-include` |
| `HxEncoding(encoding string)` | `hx-encoding` |
| `HxConfirm(message string)` | `hx-confirm` |
| `HxPrompt(message string)` | `hx-prompt` |

### Control Flow (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxIndicator(selector string)` | `hx-indicator` |
| `HxSync(strategy sync.Strategy)` | `hx-sync` |
| `HxDisabledElt(selector string)` | `hx-disabled-elt` |
| `HxDisable()` | `hx-disable` |
| `HxHistoryElt()` | `hx-history-elt` |

### Inheritance & History (htmx.go)

| Method | Attribute |
|--------|-----------|
| `HxExt(extensions string)` | `hx-ext` |
| `HxHistory(value string)` | `hx-history` |
| `HxDisinherit(attributes string)` | `hx-disinherit` |
| `HxInherit(attributes string)` | `hx-inherit` |
| `HxRequest(config string)` | `hx-request` |

To use an extension, load its script in the page `<head>` (each is its own `htmx-ext-<name>` package) and switch it on with `hx-ext`. `hx-ext` is inherited, so call `HxExt` once on the `<body>` and every element below can use the extension's methods without repeating `HxExt`. The tables below list each extension's setters.

```go
// Scripts in <head>: core htmx, then each htmx-ext-<name> package.
// Enable the extensions once on the body (inherited by all descendants):
htmx.New(body).HxExt("ws,sse")

// Then use the methods on any element below; no HxExt needed:
htmx.New(chat).WsConnect("/ws/chat")
htmx.New(form).WsSend()
htmx.New(feed).SSEConnect("/events").SSESwap("message")
```

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

### Deprecated (htmx.go)

| Method | Attribute | Note |
|--------|-----------|------|
| `HxVars(variables string)` | `hx-vars` | Use `HxVals` instead |

## Server Functions (server.go)

### Request Detection

| Function | Returns |
|----------|---------|
| `HxRequest(r *http.Request)` | `bool` - true if HTMX request |
| `Handle(r *http.Request, fn func())` | `bool` - executes fn if HTMX request, returns true |

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
trigger.Write(node, statusCode)
```

### Simple Response

```go
htmx.Response(w, div.Text("content"), http.StatusOK)
```

### SSE Server Writer (sse_server.go)

| Function | Returns |
|----------|---------|
| `NewSSE(w http.ResponseWriter)` | `(*SSEWriter, error)` - initialises SSE stream, sets headers |
| `(*SSEWriter).Send(event, data string)` | `error` - sends a named event, handles multi-line data, flushes |

```go
sse, err := htmx.NewSSE(w)
sse.Send("message", "<div>Updated</div>")
sse.Send("done", "")  // triggers sse-close on client
```

### Header Constants

Prefer the reader and writer functions above. These exported constants hold the raw header names, for cases where no helper exists:

```go
// Request headers (sent by the client)
htmx.HXRequestHeader               // "HX-Request"
htmx.HXBoostedHeader               // "HX-Boosted"
htmx.HXCurrentURLHeader            // "HX-Current-URL"
htmx.HXHistoryRestoreRequestHeader // "HX-History-Restore-Request"
htmx.HXPromptHeader                // "HX-Prompt"
htmx.HXTargetHeader                // "HX-Target"
htmx.HXTriggerNameHeader           // "HX-Trigger-Name"
htmx.HXTriggerHeader               // "HX-Trigger"

// Response headers (set by the server)
htmx.HXLocationHeader           // "HX-Location"
htmx.HXPushURLHeader            // "HX-Push-Url"
htmx.HXRedirectHeader           // "HX-Redirect"
htmx.HXRefreshHeader            // "HX-Refresh"
htmx.HXReplaceURLHeader         // "HX-Replace-Url"
htmx.HXReswapHeader             // "HX-Reswap"
htmx.HXRetargetHeader           // "HX-Retarget"
htmx.HXReselectHeader           // "HX-Reselect"
htmx.HXTriggerAfterSettleHeader // "HX-Trigger-After-Settle"
htmx.HXTriggerAfterSwapHeader   // "HX-Trigger-After-Swap"
```

## Configuration (config.go)

`htmx.Config()` creates a builder for generating HTMX `<meta>` configuration tags. All methods return `*config` for chaining.

```go
cfg := htmx.Config().
    DefaultSwapStyle(swap.OuterHTML).
    Timeout(5000).
    GlobalViewTransitions(true)

metaTag, err := cfg.ToMetaTag()
// <meta name="htmx-config" content='{"defaultSwapStyle":"outerHTML","timeout":5000,...}'>

jsonStr, err := cfg.ToJSON()
```

### Config Methods

| Method | Default | Description |
|--------|---------|-------------|
| `DefaultSwapStyle(swap.Strategy)` | `swap.InnerHTML` | Default swap method |
| `DefaultSwapDelay(int)` | `0` | Delay in ms before swap |
| `DefaultSettleDelay(int)` | `20` | Delay in ms before settle |
| `Timeout(int)` | `0` | Request timeout in ms |
| `HistoryEnabled(bool)` | `true` | Enable history snapshots |
| `HistoryCacheSize(int)` | `10` | Max cached history pages |
| `RefreshOnHistoryMiss(bool)` | `false` | Full refresh on cache miss |
| `HistoryRestoreAsHxRequest(bool)` | `true` | Send HX-Request on history restore |
| `GlobalViewTransitions(bool)` | `false` | Use View Transitions API |
| `ScrollBehaviour(string)` | `"instant"` | Scroll animation style |
| `ScrollBehavior(string)` | - | American spelling alias |
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
| `TriggerSpecsCache(interface{})` | - | Pre-populated trigger spec cache |
| `ResponseHandling(interface{})` | - | Custom response handling rules |
| `ToMetaTag()` | - | Returns `(string, error)` |
| `ToJSON()` | - | Returns `(string, error)` |

## Usage Patterns

### Form Submission

```go
htmx.New(form).HxPost("/save").HxTarget("#content").HxSwap(swap.InnerHTML)

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
htmx.New(form).HxPost("/add").HxTarget("#list").HxSwap(swap.AfterBegin)
// Swap strategies: swap.InnerHTML, swap.OuterHTML, swap.BeforeBegin, swap.AfterBegin, swap.BeforeEnd, swap.AfterEnd
```

### Inline Event Handler (Locality of Behaviour)

```go
handler := `document.querySelectorAll('.item').forEach(el => el.classList.remove('active'));
event.target.closest('.item').classList.add('active');`
htmx.New(div).HxOn("after-swap", handler)
```

### Delete with Confirmation

```go
htmx.New(btn).HxDelete("/items/"+id).HxConfirm("Sure?").HxTarget("closest .item").HxSwap(swap.OuterHTML)
```

### Active search (type-ahead)

```go
htmx.New(input).
    HxGet("/search").
    HxTrigger("keyup changed delay:300ms").
    HxTarget("#results").
    HxSwap(swap.InnerHTML)
```

### Load more (append to a list)

```go
htmx.New(btn).HxGet("/items?page=2").HxTarget("#list").HxSwap(swap.BeforeEnd)
```

### Polling

```go
htmx.New(div).HxGet("/status").HxTrigger("every 2s").HxSwap(swap.InnerHTML)
```

### Lazy load on reveal

```go
htmx.New(div).HxGet("/widget").HxTrigger("load").HxSwap(swap.OuterHTML)
```

### Loading indicator

```go
htmx.New(btn).HxPost("/save").HxTarget("#out").HxIndicator("#spinner")
```

### Boosted navigation

```go
// Links and forms inside become AJAX requests, and the URL is pushed to history.
htmx.New(nav).HxBoost(true)
```

### Server: trigger a client event after the response

```go
func HandleSave(w http.ResponseWriter, r *http.Request) {
    t := htmx.NewTrigger(w)
    t.AddTrigger("itemSaved", map[string]any{"id": id})
    t.Write(SavedPartial(), http.StatusOK)
}
```

### Server: client-side redirect

```go
// For an htmx request this sets HX-Redirect; otherwise it is a normal HTTP redirect.
htmx.HxRedirect(w, r, "/login", http.StatusSeeOther)
```
