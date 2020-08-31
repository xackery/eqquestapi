---
title: event_signal
searchTitle: Lua Event event_signal
description: Triggers when a NPC gets a signal
weight: 1
hidden: true
menuTitle: event_signal
---

#### event_signal

Triggers when a NPC gets a signal. A signal is typically invoked by another NPC.

**Syntax**

function {{% lua_type_functionname event_signal %}}({{% lua_type_class %}} e) -- {{% lua_type_nil %}}

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_signal(e)
    e.self:Say("Got signal " .. e.signal);
end
```

**Exported Variables**

These variables are available with event_signal
- {{% lua_type_mob %}} **e.self**: mob receiving the event
- {{% lua_type_number %}} **e.signal**: signal index
