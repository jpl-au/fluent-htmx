# HTMX Extension Guide for LLMs

## Constraints

- **NEVER use `.SetAttribute()` for HTMX attributes** - always use dedicated `htmx.New(element).HxX()` methods
- **NEVER use `r.Header.Get("HX-Request")`** - use `htmx.HxRequest(r)`
- **NEVER use `w.Header().Set("HX-*")`** - use `htmx.HxX(w, ...)` functions
- **NEVER add global JavaScript event listeners for HTMX** - use `HxOn()` method for locality of behaviour
- All HTMX attributes have typed methods in `client.go`
- All server request/response helpers are in `server.go`

## Pattern

`htmx.New()` accepts `node.Element` — any HTML element created via Fluent's element packages. Text nodes, function components, and conditionals are `node.Node` only and cannot be wrapped.

```go
// Client-side: wrap element, chain HTMX methods
htmx.New(element).HxPost("/api/endpoint").HxTarget("#result").HxOn("after-swap", "console.log('done')")

// Server-side: use helper functions
if htmx.HxRequest(r) { /* partial */ } else { /* full page */ }
htmx.HxPushURL(w, "/new-url")
```

## Client Methods (client.go)

### HTTP
`HxGet(url)` `HxPost(url)` `HxPut(url)` `HxPatch(url)` `HxDelete(url)`

### Swap
`HxSwap(strategy)` `HxTarget(selector)` `HxSwapOOB(value)` `HxSelect(selector)` `HxSelectOOB(selector)`

### Triggers
`HxTrigger(events)` - e.g. `"keyup changed delay:500ms"`

### Events (Locality of Behaviour)
`HxOn(event, jsCode)` - attach inline event handlers directly to elements

Use constants from `constants.go` for type-safe event names: `EventAfterSwap`, `EventBeforeSwap`, `EventAfterSettle`, `EventBeforeRequest`, `EventAfterRequest`, `EventConfigRequest`, etc.

```go
htmx.New(div).HxOn(htmx.EventAfterSwap, `/* JavaScript */`)
htmx.New(div).HxOn("click", `/* also accepts standard DOM events */`)
```

### Boolean
`HxBoost(bool)` `HxPreserve(bool)` `HxValidate(bool)` - all accept bool, convert to "true"/"false" strings

### URL Management
`HxPushURL(value)` `HxReplaceURL(url)` - accepts "true", "false", or custom URL string

### Form & Validation
`HxVals(json)` `HxHeaders(json)` `HxParams(params)` `HxInclude(selector)` `HxEncoding(encoding)` `HxConfirm(message)` `HxPrompt(message)`

### Control Flow
`HxIndicator(selector)` `HxSync(strategy)` `HxDisabledElt(selector)` `HxDisable()` `HxHistoryElt()`

### Advanced
`HxExt(extensions)` `HxHistory(value)` `HxDisinherit(attributes)` `HxInherit(attributes)` `HxRequest(config)`

## Server Functions (server.go)

### Request Detection
`HxRequest(r)` - returns bool if HTMX request
`Handle(r, fn)` - executes fn if HTMX request, returns bool

### Read Request Headers
`HxBoosted(r)` `HxCurrentURL(r)` `HxHistoryRestoreRequest(r)` `HxTarget(r)` `HxTrigger(r)` `HxTriggerName(r)` `HxPrompt(r)`

### Write Response Headers
**URL:** `HxPushURL(w, url)` `HxReplaceURL(w, url)` `HxRedirect(w, r, url, code)` `HxLocation(w, url)`
**Swap:** `HxRetarget(w, selector)` `HxReswap(w, strategy)` `HxReselect(w, selector)`
**Other:** `HxRefresh(w)`

### Trigger Events
```go
trigger := htmx.NewTrigger(w)
trigger.AddTrigger(eventName, detailMap)           // immediate
trigger.AddTriggerAfterSwap(eventName, detailMap)  // after swap
trigger.AddTriggerAfterSettle(eventName, detailMap) // after settle
trigger.Write(content, statusCode)
```

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

### Active State (Inline Event Handler)
```go
handler := `document.querySelectorAll('.item').forEach(el => el.classList.remove('active'));
event.target.closest('.item').classList.add('active');`
htmx.New(div).HxOn("after-swap", handler)
```

### Delete with Confirmation
```go
htmx.New(btn).HxDelete("/items/"+id).HxConfirm("Sure?").HxTarget("closest .item").HxSwap("outerHTML")
```
