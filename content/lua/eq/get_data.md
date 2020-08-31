---
title: get_data
searchTitle: Lua EQ get_data
description: Gets data from a data bucket.
weight: 1
hidden: true
menuTitle: get_data
---

#### get_data

Gets data from a data bucket.

**Syntax**

function {{% lua_type_functionname get_data %}}({{% lua_type_string %}} bucket_key) -- {{% lua_type_string %}}

**Parameters**

- {{% lua_type_string %}} **bucket_key**: Global scoped key stored in data_buckets table

**Example**

```lua
function event_say(e)
    my key = e.self:GetNPCTypeID() .. "-custom-key";
    eq.get_data(key, "custom data");
    my value = eq.get_data(key);
    e.self:Say(key .. " is set to " .. value);
end
```