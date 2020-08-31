---
title: register_item_event
searchTitle: Lua EQ register_item_event
description: Registers an item event.
weight: 1
hidden: true
menuTitle: register_item_event
---

#### register_item_event

Registers an item event.

**Syntax**

eq.{{% lua_type_functionname register_item_event %}}({{% lua_type_number %}} evt, {{% lua_type_number %}} item_id, {{% lua_type_class %}} func) -- {{% lua_type_nil %}}
eq.{{% lua_type_functionname register_item_event %}}({{% lua_type_string %}} name, {{% lua_type_number %}} evt, {{% lua_type_number %}} item_id, {{% lua_type_class %}} func) -- {{% lua_type_nil %}}

**Parameters**

- {{% lua_type_number %}} **evt**: Event ID, e.g. Event.Trade
- {{% lua_type_number %}} **item_id**: Item ID to tie event notification to
- {{% lua_type_class %}} **func**: function to invoke on finish

**Example**

```lua
function event_say(e)
    eq.register_item_event(Event.Trade, 1001, custom_event);
    my key = e.self:GetNPCTypeID() .. "-custom-key";
    eq.register_item_event(key, "custom data");
    my value = eq.get_data(key);
    e.self:Say(key .. " is set to " .. value);
end

function custom_event()
    e.self:Say("custom event triggered");
    eq.unregister_item_event(Event.Trade, 1001);
end
```