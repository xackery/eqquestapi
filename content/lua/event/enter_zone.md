---
title: event_enter_zone
searchTitle: Lua Event event_enter_zone
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_enter_zone
---

#### event_enter_zone

Triggers when a player enter a specified zone of a NPC

**Syntax**

function {{% lua_type_functionname event_enter_zone %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_enter_zone(e)
    e.self:enter_zone("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_enter_zone
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in