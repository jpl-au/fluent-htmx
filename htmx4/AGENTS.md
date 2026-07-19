# HTMX Extension Guide for Agents

## Methods That Do Not Exist

The following methods **do not exist** in this package. Do not use them - use the listed replacement:

- `Attr()`, `SetAttr()`, `Attribute()` - do not exist. Use the typed `Hx*()` methods.
- `Data()`, `Aria()`, `SetData()`, `SetAria()` - not available on `*Wrapper`. Call them on the Fluent element **before** wrapping it with `htmx.New()`.
- `HxRequest()` (client setter) - does not exist. Use `HxConfig()` (sets `hx-config`).
- `HxVars()` - does not exist. Use `HxVals()` with a `js:` prefix for computed values.
- `HxExt()` - does not exist. Extensions load by including their script; restrict them with the `Extensions` config key.
- `HxParams()`, `HxDisinherit()`, `HxInherit()` - do not exist. For inheritance, use the `htmx.Inherited` / `htmx.InheritedAppend` modifier.
- `HxDisabledElt()` - does not exist. Use `HxDisable(selector)`.
- `HxDisable()` with no arguments - does not exist. `HxDisable` takes a selector (disable form elements during a request). To skip HTMX processing, use `HxIgnore()`.
- `HxTargetError()`, `HxTargetCode()`, `HxTargetCodePattern()` - do not exist. Use `HxStatus(code, spec)`.
- `SSESwap()` - does not exist. Unnamed SSE messages swap automatically; handle named events with `HxTrigger`.
- `PreloadImages()` - does not exist. `Preload(trigger)` is the only preload method.
- `HxTriggerName(r)` (server) - does not exist. Use `HxSource(r)`.
- `AddTriggerAfterSwap()`, `AddTriggerAfterSettle()` - do not exist. Trigger events fire immediately; use `AddTrigger()`.

If a method is not listed in this document, it does not exist.

## Architecture

`htmx.New(element)` wraps a Fluent `node.Element` and returns `*Wrapper`. The Wrapper delegates these `node.Element` methods to the underlying element: `Render`, `WriteTo`, `RenderBytes`, `RenderBuilder`, `RenderOpen`, `RenderClose`, `Nodes`, `SetAttribute`, `SetAttributeRaw`. (`SetAttributeRaw` is the trusted-value raw hatch, mirroring `RawText`: it stores the value verbatim, whereas `SetAttribute` escapes it.) All other methods on `*Wrapper` are the HTMX-specific methods listed in this document.

The package is imported from `github.com/jpl-au/fluent-htmx/htmx4` but is named `htmx`; call sites read `htmx.New(...)`.

`SetAttribute(key, value)` is exposed on Wrapper as a pass-through to the underlying element. **Never call `SetAttribute` directly for HTMX attributes** - always use the typed `Hx*()` methods instead.

## Constraints

- **NEVER use `.SetAttribute()` for HTMX attributes** - always use `htmx.New(element).HxX()` methods
- **NEVER use `r.Header.Get("HX-Request")`** - use `htmx.HxRequest(r)`
- **NEVER use `w.Header().Set("HX-*")`** - use `htmx.HxX(w, ...)` functions
- **NEVER add global JavaScript event listeners for HTMX** - use `HxOn()` for locality of behaviour
- If a method is not listed in this document, it does not exist

## Pattern

```go
// Client-side: wrap element, chain HTMX methods
htmx.New(element).HxPost("/api/endpoint").HxTarget("#result").HxOn("htmx:after:swap", "console.log('done')")

// Server-side: use helper functions
if htmx.HxRequest(r) { /* partial */ } else { /* full page */ }
htmx.HxPushURL(w, "/new-url")
```

## Attribute Inheritance

`htmx.Mod` is an optional modifier for the inheritable setters (those whose signature ends in `...htmx.Mod`). Pass `htmx.Inherited` to emit `:inherited` (the attribute inherits to descendants), `htmx.InheritedAppend` to emit `:inherited:append` (a descendant appends to the inherited value instead of replacing it), or `htmx.Append` for a bare `:append` (a terminal append that does not re-propagate to the descendant's own children). These modifier values assume htmx's default `:` meta character.

```go
htmx.New(div).HxConfirm("Sure?", htmx.Inherited)          // hx-confirm:inherited
htmx.New(form).HxInclude("#fields", htmx.InheritedAppend) // hx-include:inherited:append
```

Inheritable setters: `HxTarget`, `HxSwap`, `HxTrigger`, `HxBoost`, `HxConfirm`, `HxVals`, `HxHeaders`, `HxIndicator`, `HxPushURL`, `HxReplaceURL`, `HxSelect`, `HxSelectOOB`, `HxInclude`, `HxSync`, `HxEncoding`, `HxValidate`, `HxDisable`, `HxConfig`, `HxStatus`. Only `HxPreserve`, `HxSwapOOB`, `HxHistoryElt` and `HxIgnore` take no modifier - htmx reads those by presence or by id on a single element. The request verbs, `HxAction` and `HxMethod` also omit the modifier (an inherited verb cannot by itself fire a request).

## Sub-Packages

Swap strategies, sync strategies, events, and CSS classes live in their own packages for cleaner call sites.

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
| `swap.InnerMorph` | `"innerMorph"` |
| `swap.OuterMorph` | `"outerMorph"` |
| `swap.TextContent` | `"textContent"` |
| `swap.Download` | `"download"` (download extension, bundled) |
| `swap.Upsert` | `"upsert"` (upsert extension, separate script) |

`swap.Custom(strategy string) swap.Strategy` - creates a strategy with modifiers, e.g. `swap.Custom("innerHTML show:top showTarget:#other")`.

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

`sync.Custom(strategy string) sync.Strategy` - e.g. `sync.Custom("this:queue all")`.

### Events (`event` package)

Event constants follow the `htmx:phase:action` scheme, for use with `HxOn()`. Examples:

| Constant | Value |
|----------|-------|
| `event.BeforeRequest` | `"htmx:before:request"` |
| `event.AfterRequest` | `"htmx:after:request"` |
| `event.BeforeSwap` | `"htmx:before:swap"` |
| `event.AfterSwap` | `"htmx:after:swap"` |
| `event.ConfigRequest` | `"htmx:config:request"` |
| `event.ResponseError` | `"htmx:response:error"` |
| `event.Error` | `"htmx:error"` |

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

This is the **exhaustive** list of methods on `*Wrapper`. If a method is not listed here, it does not exist. Setters whose signature ends in `...Mod` accept the inheritance modifier.

### HTTP Verbs

| Method | Attribute |
|--------|-----------|
| `HxGet(url string)` | `hx-get` |
| `HxPost(url string)` | `hx-post` |
| `HxPut(url string)` | `hx-put` |
| `HxPatch(url string)` | `hx-patch` |
| `HxDelete(url string)` | `hx-delete` |
| `HxAction(url string)` | `hx-action` |
| `HxMethod(method string)` | `hx-method` |

### Swap & Targeting

| Method | Attribute |
|--------|-----------|
| `HxSwap(strategy swap.Strategy, mods ...Mod)` | `hx-swap` |
| `HxTarget(selector string, mods ...Mod)` | `hx-target` |
| `HxSelect(selector string, mods ...Mod)` | `hx-select` |
| `HxSelectOOB(selector string, mods ...Mod)` | `hx-select-oob` |
| `HxSwapOOB(value string)` | `hx-swap-oob` |
| `HxStatus(code string, spec string, mods ...Mod)` | `hx-status:CODE` |

### Triggers & Events

| Method | Attribute |
|--------|-----------|
| `HxTrigger(events string, mods ...Mod)` | `hx-trigger` |
| `HxOn(event string, handler string)` | `hx-on:event` (single colon) |

### Boolean Attributes

| Method | Attribute |
|--------|-----------|
| `HxBoost(enabled bool, mods ...Mod)` | `hx-boost` |
| `HxPreserve()` | `hx-preserve` |
| `HxValidate(validate bool, mods ...Mod)` | `hx-validate` |

### URL Management

| Method | Attribute |
|--------|-----------|
| `HxPushURL(value string, mods ...Mod)` | `hx-push-url` |
| `HxReplaceURL(url string, mods ...Mod)` | `hx-replace-url` |

### Form & Request Parameters

| Method | Attribute |
|--------|-----------|
| `HxVals(json string, mods ...Mod)` | `hx-vals` |
| `HxHeaders(json string, mods ...Mod)` | `hx-headers` |
| `HxInclude(selector string, mods ...Mod)` | `hx-include` |
| `HxEncoding(encoding string, mods ...Mod)` | `hx-encoding` |
| `HxConfirm(message string, mods ...Mod)` | `hx-confirm` |
| `HxConfig(json string, mods ...Mod)` | `hx-config` |

### Control Flow

| Method | Attribute |
|--------|-----------|
| `HxIndicator(selector string, mods ...Mod)` | `hx-indicator` |
| `HxSync(strategy sync.Strategy, mods ...Mod)` | `hx-sync` |
| `HxDisable(selector string, mods ...Mod)` | `hx-disable` |
| `HxIgnore()` | `hx-ignore` |
| `HxHistoryElt()` | `hx-history-elt` |

### Extensions

| Method | Attribute |
|--------|-----------|
| `WsConnect(url string)` | `hx-ws:connect` |
| `WsSend()` | `hx-ws:send` |
| `SSEConnect(url string)` | `hx-sse:connect` |
| `SSEClose(eventName string)` | `hx-sse:close` |
| `Preload(trigger string)` | `hx-preload` |
| `HxOptimistic(selector string)` | `hx-optimistic` |
| `HxTargets(selector string)` | `hx-targets` |
| `HxLive(expression string)` | `hx-live` |
| `HxBrowserIndicator(enabled bool)` | `hx-browser-indicator` |
| `HxHead(mode string)` | `hx-head` (separate script dist/ext/hx-head.js) |
| `HxPtag(tag string)` | `hx-ptag` (separate script dist/ext/hx-ptag.js) |
| `HxHistory(enabled bool)` | `hx-history` (separate script dist/ext/hx-history-cache.js) |
| `HxNonce(nonce string)` | `hx-nonce` (separate script dist/ext/hx-csp.js) |
| `HxPrompt(question string)` | `hx-prompt` (separate script dist/ext/hx-prompt.js) |

htmx 4 ships two builds. `htmax.js` bundles 8 extensions (ws, sse, preload, optimistic, targets, live, browser-indicator, download), so their methods work once it is loaded. The core `htmx.js` build includes none of them; load each extension's `dist/ext/<name>.js` script yourself. `HxHead`, `HxPtag`, `HxHistory`, `HxNonce`, `HxPrompt` and `swap.Upsert` are never bundled and always need their own script. htmx 4 has no `hx-ext`: loading the script, or the bundle, activates the extension. A method whose extension is absent is an inert no-op, so the attribute is written but ignored. Downloads have no client attribute; use `swap.Download` or the `HxDownload(w, url)` server helper. The `htmx-2-compat` and `hx-alpine-compat` extensions add no per-element attribute.

```go
// Bundled in htmax.js, work with no extra setup:
htmx.New(div).SSEConnect("/events")
htmx.New(btn).HxPost("/save").HxTarget("#list").HxOptimistic("#pending")
htmx.New(link).HxGet("/file.pdf").HxSwap(swap.Download)

// Separate-script extensions (load dist/ext/<name>.js):
htmx.New(div).HxGet("/news").HxTrigger("every 3s").HxPtag("v42")
htmx.New(btn).HxGet("/items").HxSwap(swap.Upsert)

// Server-side download (sets the HX-Download header):
htmx.HxDownload(w, "/files/report.pdf")
```

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
| `HxTarget(r)` | `string` (`tagName#id`) |
| `HxSource(r)` | `string` (`tagName#id`) |
| `HxRequestType(r)` | `string` (`"full"` or `"partial"`) |
| `HxPrompt(r)` | `string` (empty unless a prompt extension sets HX-Prompt; htmx 4 core never sends it) |

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
| `HxDownload(w, url)` | Stream a file download (download extension reads HX-Download) |

### Trigger Events

Events fire immediately when the response is received (the `HX-Trigger` header). There are no after-swap or after-settle phases.

```go
trigger := htmx.NewTrigger(w)
trigger.AddTrigger(eventName, detailMap)
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
| `(*SSEWriter).Send(event string, n node.Node)` | `error` - renders the node and sends a named event, one data line per physical line, flushes; a nil node sends an empty data line |
| `(*SSEWriter).SendBytes(event string, data []byte)` | `error` - escape hatch for non-fluent payloads; sends raw bytes as the event data, one data line per physical line, flushes |

```go
sse, err := htmx.NewSSE(w)
sse.Send("message", div.Text("Updated"))
sse.SendBytes("message", buf.Bytes())  // non-fluent payload
sse.Send("done", nil)  // closes the stream on the client's hx-sse:close event
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
htmx.HXSourceHeader                // "HX-Source"
htmx.HXRequestTypeHeader           // "HX-Request-Type"

// Response headers (set by the server)
htmx.HXTriggerHeader     // "HX-Trigger"
htmx.HXLocationHeader    // "HX-Location"
htmx.HXPushURLHeader     // "HX-Push-Url"
htmx.HXRedirectHeader    // "HX-Redirect"
htmx.HXRefreshHeader     // "HX-Refresh"
htmx.HXReplaceURLHeader  // "HX-Replace-Url"
htmx.HXReswapHeader      // "HX-Reswap"
htmx.HXRetargetHeader    // "HX-Retarget"
htmx.HXReselectHeader    // "HX-Reselect"
htmx.HXDownloadHeader    // "HX-Download"
```

## Configuration (config.go)

`htmx.Config()` creates a builder for generating HTMX `<meta>` configuration tags. All methods return `*config` for chaining.

```go
cfg := htmx.Config().
    DefaultSwapStyle(swap.OuterHTML).
    DefaultTimeout(5000).
    Transitions(true)

metaTag, err := cfg.ToMetaTag()
jsonStr, err := cfg.ToJSON()
```

### Config Methods

| Method | Default | Key |
|--------|---------|-----|
| `DefaultSwapStyle(swap.Strategy)` | `swap.InnerHTML` | `defaultSwap` |
| `DefaultSettleDelay(int)` | `1` | `defaultSettleDelay` |
| `DefaultTimeout(int)` | `60000` | `defaultTimeout` |
| `Transitions(bool)` | `false` | `transitions` |
| `HistoryEnabled(bool)` | `true` | `history` |
| `ImplicitInheritance(bool)` | `false` | `implicitInheritance` |
| `NoSwap([]string)` | `["204","304"]` | `noSwap` |
| `Mode(string)` | `"same-origin"` | `mode` |
| `Extensions(string)` | `""` | `extensions` |
| `MetaCharacter(string)` | `":"` | `metaCharacter` |
| `IncludeIndicatorCSS(bool)` | `true` | `includeIndicatorCSS` |
| `IndicatorClass(string)` | `"htmx-indicator"` | `indicatorClass` |
| `RequestClass(string)` | `"htmx-request"` | `requestClass` |
| `InlineScriptNonce(string)` | `""` | `inlineScriptNonce` |
| `DefaultFocusScroll(bool)` | `false` | `defaultFocusScroll` |
| `MorphIgnore([]string)` | `["data-htmx-powered"]` | `morphIgnore` |
| `MorphScanLimit(int)` | `10` | `morphScanLimit` |
| `MorphSkip(string)` | `""` | `morphSkip` |
| `MorphSkipChildren(string)` | `""` | `morphSkipChildren` |
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

### Error Handling with Status Swaps

The code may be an exact status (`"404"`), a single-digit wildcard (`"50x"`), or a range wildcard (`"5xx"`). The spec uses htmx's `key:value` syntax with the keys `swap:`, `target:`, `select:`, `push:`, `replace:` and `transition:`.

```go
htmx.New(form).
    HxPost("/save").
    HxStatus("422", "swap:innerHTML target:#errors").
    HxStatus("5xx", "swap:none")
```

### Inherited Defaults for a Subtree

```go
// Every descendant inherits the confirmation and target.
htmx.New(container).
    HxConfirm("Are you sure?", htmx.Inherited).
    HxTarget("#main", htmx.Inherited)
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
