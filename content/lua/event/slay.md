---
title: event_slay
searchTitle: Lua Event event_slay
searchDescription: Triggers when a NPC aggros anything
weight: 1
hidden: true
menuTitle: event_slay
---

#### event_slay

Triggers when a npc is slain

**Syntax**

function {{% lua_type_functionname event_slay %}}({{% lua_type_class %}} e)

**Parameters**

- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_slay(e)
    e.self:Say("Hail " .. e.other:GetName() .. "!");
end
```

**Exported Variables**

These variables are available with event_slay
- {{% type_client lua %}} **e.other**: player that triggered the event
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_string %}} **e.message**: player message that triggered the event
- {{% lua_type_number %}} **e.language**: language id the message was in