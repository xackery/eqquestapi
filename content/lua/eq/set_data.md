---
title: set_data
searchTitle: Lua EQ set_data
description: Sets a data bucket to provided data. By default, expires_at is never (0)
weight: 1
hidden: true
menuTitle: set_data
---

#### set_data

Sets a data bucket to provided data. By default, expires_at is never (0)

**Syntax**

eq.{{% lua_type_functionname set_data %}}({{% lua_type_string %}} bucket_key, {{% lua_type_string %}} bucket_value) -- {{% lua_type_nil %}}

eq.{{% lua_type_functionname set_data %}}({{% lua_type_string %}} bucket_key, {{% lua_type_string %}} bucket_value. {{% lua_type_string %}} expires_at) -- {{% lua_type_nil %}}

**Parameters**

- {{% lua_type_string %}} **bucket_key**: Global scoped key stored in data_buckets table
- {{% lua_type_string %}} **bucket_value**: Global scoped value stored in data_buckets table

**Example**

```lua
function event_say(e)
    my key = e.self:GetNPCTypeID() .. "-custom-key";
    eq.set_data(key, "custom data");
    my value = eq.get_data(key);
    e.self:Say(key .. " is set to " .. value);
end
```