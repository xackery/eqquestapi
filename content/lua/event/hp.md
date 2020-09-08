---
title: event_hp
searchTitle: Lua Event event_hp
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_hp
---

#### event_hp

Triggers when a NPC hits a certain amount of HP

**Syntax**

function {{% lua_type_functionname event_hp %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_hp(e)
    e.self:Say("I have arrived at a waypoint");
end
```

**Exported Variables**

These variables are available with event_hp
- {{% type_npc mob %}} **e.self**: mob emitting the event