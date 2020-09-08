---
title: event_proximity_say
searchTitle: Lua Event event_proximity_say
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_proximity_say
---

#### event_proximity_say

Triggers when a player says something in a proximity of a NPC

**Syntax**

function {{% lua_type_functionname event_proximity_say %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_proximity_say(e)
    e.self:proximity_say("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_proximity_say
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in