---
title: event_killed_merit
searchTitle: Lua Event event_killed_merit
searchDescription: Triggers when a player talks to a NPC
weight: 1
hidden: true
menuTitle: event_killed_merit
---

#### event_killed_merit

Triggers when a player kills a NPC

**Syntax**

function {{% lua_type_functionname event_killed_merit %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_killed_merit(e)
    e.self:killed_merit("Hail " .. e.other:GetName() .. "!");        
end
```

**Exported Variables**

These variables are available with event_killed_merit
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in