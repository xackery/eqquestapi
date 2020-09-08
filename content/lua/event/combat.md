---
title: event_combat
searchTitle: Lua Event event_combat
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_combat
---

#### event_combat

Triggers when a NPC enters combat

**Syntax**

function {{% lua_type_functionname event_combat %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_combat(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");
end
```

**Exported Variables**

These variables are available with event_combat
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in