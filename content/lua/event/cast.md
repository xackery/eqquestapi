---
title: event_cast
searchTitle: Lua Event event_cast
searchDescription: Triggers when a NPC casts anything
weight: 1
hidden: true
menuTitle: event_cast
---

#### event_cast

Triggers when a NPC casts

**Syntax**

function {{% lua_type_functionname event_cast %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_cast(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");
end
```

**Exported Variables**

These variables are available with event_cast
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in