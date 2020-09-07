---
title: event_cast_begin
searchTitle: Lua Event event_cast_begin
searchDescription: Triggers when a NPC begins to cast
weight: 1
hidden: true
menuTitle: event_cast_begin
---

#### event_cast_begin

Triggers when a NPC begins to cast

**Syntax**

function {{% lua_type_functionname event_cast_begin %}}({{% lua_type_class %}} e)


**Parameters**
- {{% lua_type_class %}} **e**: Short for event, extended to contain any data relevant to the event

**Example**

```lua
function event_cast_begin(e)
    e.self:Say("spell " .. e.spell:GetName() .. " has an augment!");
end
```

**Exported Variables**

These variables are available with event_cast_begin
- {{% lua_type_spell %}} **e.spell**: spell that triggered the event