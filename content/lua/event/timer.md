---
title: event_timer
searchTitle: Lua Event event_timer
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_timer
---

#### event_timer

Triggers when a NPC timer happens

**Syntax**

function {{% lua_type_functionname event_timer %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_timer(e)
    e.self:Say("I have arrived at a waypoint");
end
```

**Exported Variables**

These variables are available with event_timer
- {{% type_npc mob %}} **e.self**: mob emitting the event