---
title: event_player_environmental_damage
searchTitle: Lua Event event_player_environmental_damage
searchDescription: Triggers when a player takes damage
weight: 1
hidden: true
menuTitle: event_player_environmental_damage
---

#### event_player_environmental_damage

Triggers when a player takes damage

**Syntax**

function {{% lua_type_functionname event_player_environmental_damage %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_player_environmental_damage(e)
    e.self:player_environmental_damage("got something!");
end
```

**Exported Variables**

These variables are available with event_player_environmental_damage
- {{% type_client lua %}} **e.self**: player that triggered the event