---
title: event_attack
searchTitle: Lua Event event_attack
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_attack
---

#### event_attack

Triggers when a NPC attacks

**Syntax**

function {{% lua_type_functionname event_attack %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_attack(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");
end
```

**Exported Variables**

These variables are available with event_attack
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in