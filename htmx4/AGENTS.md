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
- `AddTriggerAfterSwap()`, `AddTriggerAfterSettle()` - do not exist. htmx 4 fires HX-Trigger events after the swap; use `AddTrigger()`.

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

Inheritable setters: `HxTarget`, `HxSwap`, `HxTrigger`, `HxBoost`, `HxBoostConfig`, `HxConfirm`, `HxVals`, `HxHeaders`, `HxIndicator`, `HxPushURL`, `HxReplaceURL`, `HxSelect`, `HxSelectOOB`, `HxInclude`, `HxSync`, `HxEncoding`, `HxValidate`, `HxDisable`, `HxConfig`, `HxStatus`. Only `HxPreserve`, `HxSwapOOB`, `HxHistoryElt` and `HxIgnore` take no modifier - htmx reads those by presence or by id on a single element. The request verbs, `HxAction` and `HxMethod` also omit the modifier (an inherited verb cannot by itself fire a request). `HxMorphSkip`, `HxMorphSkipChildren` and `HxOn` take none.

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
| `swap.OuterSync` | `"outerSync"` |
| `swap.Download` | `"download"` (download extension, bundled) |
| `swap.Upsert` | `"upsert"` (upsert extension, bundled) |

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
| `class.Pending` | `"hx-pending"` |

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
| `HxQuery(url string)` | `hx-query` |
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
| `HxOnExtended(spec string)` | `hx-on` extended form `"event -> code; ..."` with the hx-trigger grammar; the only way to bind `htmx:before:viewTransition` and `htmx:after:viewTransition` |

### Boolean Attributes

| Method | Attribute |
|--------|-----------|
| `HxBoost(enabled bool, mods ...Mod)` | `hx-boost` |
| `HxBoostConfig(config string, mods ...Mod)` | `hx-boost` with request overrides, e.g. `"target:#main swap:innerHTML"` |
| `HxPreserve()` | `hx-preserve` |
| `HxMorphSkip()` | `hx-morph-skip` |
| `HxMorphSkipChildren()` | `hx-morph-skip-children` |
| `HxValidate(validate bool, mods ...Mod)` | `hx-validate` |

### URL Management

| Method | Attribute |
|--------|-----------|
| `HxPushURL(value string, mods ...Mod)` | `hx-push-url` |
| `HxReplaceURL(url string, mods ...Mod)` | `hx-replace-url` |

### Form & Request Parameters

| Method | Attribute |
|--------|-----------|
| `HxVals(values string, mods ...Mod)` | `hx-vals` (JSON or `key:value`) |
| `HxHeaders(headers string, mods ...Mod)` | `hx-headers` (JSON or `key:value`) |
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
| `WSConnect(url string)` | `hx-ws:connect` |
| `WSSend()` | `hx-ws:send` |
| `SSEConnect(url string)` | `hx-sse:connect` |
| `SSEClose(eventName string)` | `hx-sse:close` |
| `Preload(trigger string, mods ...Mod)` | `hx-preload` |
| `HxPending(selector string)` | `hx-pending` |
| `HxTargets(selector string, mods ...Mod)` | `hx-targets` |
| `HxLive(expression string)` | `hx-live` (run for effect) |
| `HxLiveText(expression string)` | `hx-live:text` |
| `HxLiveHTML(expression string)` | `hx-live:html` |
| `HxLiveClass(expression string)` | `hx-live:class` |
| `HxLiveClassToggle(class, expression string)` | `hx-live:.<class>` |
| `HxLiveStyle(expression string)` | `hx-live:style` |
| `HxLiveAttr(name, expression string)` | `hx-live:<name>` |
| `HxBrowserIndicator(enabled bool)` | `hx-browser-indicator` |
| `HxHead(mode string)` | `hx-head` (separate script dist/ext/hx-head.js) |
| `HxPtag(tag string)` | `hx-ptag` (separate script dist/ext/hx-ptag.js) |
| `HxHistoryExclude()` | `hx-history="false"` (history-cache is bundled; off until `HistoryCacheEnabled(true)`) |
| `HxNonce(nonce string)` | `hx-nonce` (separate script dist/ext/hx-csp.js) |
| `HxPrompt(question string, mods ...Mod)` | `hx-prompt` (separate script dist/ext/hx-prompt.js) |

htmx 4 ships two builds. `htmax.js` bundles 11 extensions (registered as `sse`, `ws`, `preload`, `hx-pending`, `hx-targets`, `hx-live`, `browser-indicator`, `download`, `history-cache`, `upsert`, `alpine-compat`), so their methods work once it is loaded. The core `htmx.js` build includes none of them; load each extension's `dist/ext/<name>.js` script yourself. `HxHead`, `HxPtag`, `HxNonce`, `HxPrompt` and `MultipartConnect` are never bundled and always need their own script. htmx 4 has no `hx-ext`: loading the script, or the bundle, activates the extension. A method whose extension is absent is an inert no-op, so the attribute is written but ignored. Downloads have no client attribute; use `swap.Download` or the `HxDownload(w, url)` server helper. The `htmx-2-compat` and `hx-alpine-compat` extensions add no per-element attribute.

```go
// Bundled in htmax.js, work with no extra setup:
htmx.New(div).SSEConnect("/events")
htmx.New(btn).HxPost("/save").HxTarget("#list").HxPending("#saving")
htmx.New(link).HxGet("/file.pdf").HxSwap(swap.Download)

// Separate-script extensions (load dist/ext/<name>.js):
htmx.New(div).HxGet("/news").HxTrigger("every 3s").HxPtag("v42")
htmx.New(btn).HxPost("/save").HxNonce(nonce)

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
| `HxPTag(r)` | `string` (the tag the hx-ptag extension stored; `""` on a first poll) |
| `HxLastEventID(r)` | `string` (last SSE event id handled, sent on reconnect) |
| `HxLastPartID(r)` | `string` (last multipart part id swapped, sent on reconnect) |
| `HxPTagUnchanged(w, r, tag)` | `bool` - answers 304 and returns true when the client's tag matches; otherwise sets HX-PTag to `tag` and returns false |

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
| `HxDownload(w, url)` | Fetch a separate URL as a file download while the response body is swapped as usual |
| `HxLocationWith(w, Location)` | HX-Location as an object: path plus source, target, swap, select, values, headers, and `Push` / `Replace` taking `"true"`, `"false"` or a URL; returns `error` |
| `HxPTagResponse(w, tag)` | Set HX-PTag for the hx-ptag extension to store |

### Partials and WebSocket Messages

| Function | Returns |
|----------|---------|
| `Partial(target string, nodes ...node.Node)` | `*Wrapper` - an `<hx-partial hx-target="...">` block; chain `HxSwap` for the swap style |
| `ParseWSMessage(data []byte)` | `(*WSMessage, error)` - splits an `hx-ws:send` frame into `Values` and `Headers` |
| `WSResponse{Content, Target, Swap, Select}.JSON()` | `([]byte, error)` - the JSON frame that swaps content with overrides; plain HTML frames swap with the element's own attributes |

### Trigger Events

Events in the `HX-Trigger` header fire on the requesting element after the swap completes, so a handler can read the new content. There are no after-swap or after-settle variants; htmx 4 has one header and one timing.

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
| `(*SSEWriter).Swap(n node.Node)` | `error` - sends an unnamed message, which the client swaps into the connecting element's target |
| `(*SSEWriter).SwapBytes(data []byte)` | `error` - Swap for a non-fluent payload |
| `(*SSEWriter).Send(event string, n node.Node)` | `error` - sends a named event, which the client dispatches as a DOM event and does not swap; a nil node sends the event line alone |
| `(*SSEWriter).SendBytes(event string, data []byte)` | `error` - escape hatch for non-fluent payloads; sends raw bytes as the event data, one data line per physical line, flushes |
| `(*SSEWriter).SendEvent(e Event)` | `error` - sends an event with name, id, retry and data; the client returns the id as `Last-Event-ID` on reconnect (`htmx.LastEventIDHeader`); an id-only event moves that cursor without dispatching |
| `(*SSEWriter).Release()` | `error` - sends `hx:release`, completing the request that opened the stream while the stream stays open |

```go
sse, err := htmx.NewSSE(w)
sse.Swap(div.Text("Updated"))          // swapped into the element's target
sse.SwapBytes(buf.Bytes())             // non-fluent payload
sse.Send("saved", nil)                 // DOM event "saved" on the element, not swapped
sse.Send("done", nil)  // closes the stream on the client's hx-sse:close event
sse.SendEvent(htmx.Event{Name: "tick", ID: "42", Data: buf.Bytes()})  // resumable
sse.Release()  // complete the opening request, keep streaming
```

### Multipart Server Writer (multipart_server.go)

| Function | Returns |
|----------|---------|
| `NewMultipart(w http.ResponseWriter, t multipart.Type)` | `(*MultipartWriter, error)` - starts a `multipart.Mixed` or `multipart.Parallel` response |
| `(*MultipartWriter).WritePart(n node.Node, opts ...PartOption)` | `error` - renders the node as one part and flushes |
| `(*MultipartWriter).Close()` | `error` - writes the closing boundary |
| `PartTarget(selector)`, `PartSwap(strategy)`, `PartSelect(selector)`, `PartTrigger(events)` | `PartOption` - the part's HX-Retarget, HX-Reswap, HX-Reselect and HX-Trigger headers |
| `PartID(id)` | `PartOption` - the part's HX-Part-ID; the client returns the last swapped id as `HX-Last-Part-ID` (`htmx.HXLastPartIDHeader`) on reconnect |
| `PartRefresh()`, `PartRedirect(url)`, `PartLocation(url)` | `PartOption` - the part's HX-Refresh, HX-Redirect and HX-Location; such a part is never swapped, so send it with a nil node. A `PartTrigger` fires before the part's content swaps |
| `PartPushURL(url)`, `PartReplaceURL(url)` | `PartOption` - the part's HX-Push-Url and HX-Replace-Url, applied when the part swaps; `"true"`, `"false"` or a URL, as for `HxPushURL` |

The client needs `dist/ext/hx-multipart.js`. A one-shot response needs no attribute; `MultipartConnect(url)` opens a reconnecting stream and `MultipartClose(trigger)` closes it.

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
| `NoSwap([]string)` | `[204, 304]` | `noSwap` |
| `Mode(string)` | `"same-origin"` | `mode` |
| `Extensions(string)` | `""` | `extensions` (registration names: `sse`, `ws`, `preload`, `download`, `upsert`, `ptag`, `browser-indicator`, `history-cache`, `alpine-compat`, `compat`, `hx-pending`, `hx-targets`, `hx-live`, `hx-head`, `hx-csp`, `hx-prompt`, `hx-multipart`) |
| `HistoryReload()` | - | `history: "reload"` |
| `SafeEval(bool)` | `false` | `safeEval` (hx-csp) |
| `BoostBrowserIndicator(bool)` | `false` | `boostBrowserIndicator` (browser-indicator) |
| `MetaCharacter(string)` | `":"` | `metaCharacter` |
| `IncludeIndicatorCSS(bool)` | `true` | `includeIndicatorCSS` |
| `IndicatorClass(string)` | `"htmx-indicator"` | `indicatorClass` |
| `RequestClass(string)` | `"htmx-request"` | `requestClass` |
| `InlineScriptNonce(string)` | `""` | `inlineScriptNonce` |
| `DefaultFocusScroll(bool)` | `false` | `defaultFocusScroll` |
| `MorphIgnore([]string)` | `["data-htmx-powered"]` | `morphIgnore` |
| `MorphScanLimit(int)` | `10` | `morphScanLimit` |
| `MorphSkip(string)` | `"[hx-morph-skip]"` | `morphSkip`; a new value replaces the default |
| `MorphSkipChildren(string)` | `"[hx-morph-skip-children]"` | `morphSkipChildren`; a new value replaces the default |
| `AllowEmptySwapAfterOOB(bool)` | `false` | `allowEmptySwapAfterOOB` |
| `LogAll(bool)` | `false` | `logAll` |
| `Prefix(string)` | `"data-hx-"` | `prefix` |
| `SSEReconnect(bool)`, `SSEReconnectDelay(ms)`, `SSEReconnectMaxDelay(ms)`, `SSEReconnectMaxAttempts(int)`, `SSEReconnectJitter(float64)`, `SSEPauseOnBackground(bool)`, `SSEReleaseOn(SSERelease)` | extension defaults | `sse.*` |
| `WSReconnect(bool)`, `WSReconnectCodes([]int)`, `WSReconnectDelay(ms)`, `WSReconnectMaxDelay(ms)`, `WSReconnectMaxAttempts(int)`, `WSReconnectJitter(float64)`, `WSPauseOnBackground(bool)`, `WSMaxOutgoingMessagesQueueSize(int)`, `WSProtocols([]string)` | extension defaults | `ws.*` |
| `MultipartReconnect(bool)`, `MultipartReconnectDelay(ms)`, `MultipartReconnectMaxDelay(ms)`, `MultipartReconnectMaxAttempts(int)`, `MultipartReconnectJitter(float64)`, `MultipartPauseOnBackground(bool)` | extension defaults | `multipart.*` |
| `LiveInputDebounce(ms)`, `LiveBindPrefix(string)`, `LiveUseDollar(bool)` | `100`, `":"`, `false` | `live.*` |
| `PreloadAutoBoost(bool)`, `PreloadBoostEvent(string)`, `PreloadBoostTimeout(ms)` | `true`, `"mousedown"`, `5000` | `preload.*` |
| `HistoryCacheEnabled(bool)`, `HistoryCacheSize(int)`, `HistoryCacheRefreshOnMiss(bool)`, `HistoryCacheSwapStyle(swap.Strategy)` | off in `htmax.js`, `10`, `false`, `outerSync` | `historyCache.*` |
| `CompatDoNotTriggerOldEvents(bool)`, `CompatUseExplicitInheritance(bool)`, `CompatSwapErrorResponseCodes(bool)`, `CompatSuppressInheritanceLogs(bool)` | `false` | `compat.*` |
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
