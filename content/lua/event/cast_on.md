---
title: event_cast_on
searchTitle: Lua Event event_cast_on
searchDescription: Triggers when a NPC or player casts a spell on a target
weight: 1
hidden: true
menuTitle: event_cast_on
---

#### event_cast_on

Triggers when a NPC or player casts a spell on a target

**Syntax**

function {{% lua_type_functionname event_cast_on %}}({{% lua_type_class %}} e)


**Parameters**
- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_cast_on(e)
    e.self:Say("spell " .. e.spell:GetName() .. " has been casted on " .. e.other:GetName() .. "!");
end
```

**Exported Variables**

These variables are available with event_cast_on
- {{% type_npc mob %}} **e.other**: mob the spell is being casted on
- {{% type_npc mob %}} **e.self**: mob emitting the event
- {{% lua_type_spell %}} **e.spell**: spell that triggered the event