---
title: event_popup_response
searchTitle: Lua Event event_popup_response
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_popup_response
---

#### event_popup_response

Triggers when a player responds to a popup

**Syntax**

function {{% lua_type_functionname event_popup_response %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_popup_response(e)
    e.self:popup_response("got something!");
end
```

**Exported Variables**

These variables are available with event_popup_response
- {{% type_client lua %}} **e.self**: player that triggered the event