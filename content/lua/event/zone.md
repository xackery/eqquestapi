---
title: event_zone
searchTitle: Lua Event event_zone
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_zone
---

#### event_zone

Triggers when a player enters a zone to a NPC

**Syntax**

function {{% lua_type_functionname event_zone %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_zone(e)
    e.self:zone("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_zone
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in