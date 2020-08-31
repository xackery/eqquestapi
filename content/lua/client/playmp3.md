---
title: PlayMP3
searchTitle: Lua EQ PlayMP3
description: Clears compass's current mark.
weight: 1
hidden: true
menuTitle: PlayMP3
---

#### PlayMP3

Play an mp3 in the client's eq root directory.

**Syntax**

client:{{% lua_type_functionname PlayMP3 %}}({{% lua_type_string %}} name) -- {{% lua_type_nil %}}

**Parameters**

- **name**: mp3 file name

**Example**

```lua
function event_say(e)
    e.other:PlayMP3("abysmal.mp3");
    e.self:Say("playing abysmal sea theme song");
end
```