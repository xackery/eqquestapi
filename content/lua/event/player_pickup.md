---
title: event_player_pickup
searchTitle: Lua Event event_player_pickup
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_player_pickup
---

#### event_player_pickup

Triggers when a player picks up an item

**Syntax**

function {{% lua_type_functionname event_player_pickup %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_player_pickup(e)
    e.self:player_pickup("got something!");
end
```

**Exported Variables**

These variables are available with event_player_pickup
- {{% type_client lua %}} **e.self**: player that triggered the event