---
title: event_level_up
searchTitle: Lua Event event_level_up
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_level_up
---

#### event_level_up

Triggers when a player levels up

**Syntax**

function {{% lua_type_functionname event_level_up %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_level_up(e)
    e.self:level_up("Ding!");
end
```

**Exported Variables**

These variables are available with event_level_up
- {{% type_client lua %}} **e.self**: player that triggered the event