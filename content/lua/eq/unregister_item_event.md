---
title: unregister_item_event
searchTitle: Lua EQ unregister_item_event
searchDescription: Unregisters an item event.
weight: 1
hidden: true
menuTitle: unregister_item_event
---

#### unregister_item_event

Unregisters an item event.

**Syntax**

eq.{{% lua_type_functionname unregister_item_event %}}({{% lua_type_number %}} evt, {{% lua_type_number %}} item_id) -- {{% lua_type_nil %}}
eq.{{% lua_type_functionname unregister_item_event %}}({{% lua_type_string %}} name, {{% lua_type_number %}} evt, {{% lua_type_number %}} item_id) -- {{% lua_type_nil %}}

**Parameters**

- {{% lua_type_number %}} **evt**: Event ID, e.g. Event.Trade
- {{% lua_type_number %}} **item_id**: Item ID to tie event notification to

**Example**

```lua
function event_say(e)
    eq.unregister_item_event(Event.Trade, 1001, custom_event);
    my key = e.self:GetNPCTypeID() .. "-custom-key";
    eq.unregister_item_event(key, "custom data");
    my value = eq.get_data(key);
    e.self:Say(key .. " is set to " .. value);
end

function custom_event()
    e.self:Say("custom event triggered");
    eq.ununregister_item_event(Event.Trade, 1001);
end
```