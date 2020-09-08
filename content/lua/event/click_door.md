---
title: event_click_door
searchTitle: Lua Event event_click_door
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_click_door
---

#### event_click_door

Triggers when a player clicks a door

**Syntax**

function {{% lua_type_functionname event_click_door %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_click_door(e)
    e.self:click_door("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_click_door
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in