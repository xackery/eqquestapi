---
title: Packet Class
menuTitle: Packet Class
searchDescription: Lua Packet Class
searchTitle: Lua Packet Class
weight: 25
---

## Packet Methods
- [GetSize](getsize)(); -- {{% lua_type_number %}}
- [GetOpcode](getopcode)(); -- {{% lua_type_number %}}
- [SetOpcode](setopcode)({{% lua_type_number %}} op); -- void
- [GetRawOpcode](getrawopcode)(); -- {{% lua_type_number %}}
- [SetRawOpcode](setrawopcode)({{% lua_type_number %}} op); -- void
- [WriteFloat](writefloat)({{% lua_type_number %}} offset, float value); -- void
- [WriteDouble](writedouble)({{% lua_type_number %}} offset, double value); -- void
- [WriteString](writestring)({{% lua_type_number %}} offset, std::string value); -- void
- [WriteFixedLengthString](writefixedlengthstring)({{% lua_type_number %}} offset, std::string value, {{% lua_type_number %}} string_length); -- void
- [ReadFloat](readfloat)({{% lua_type_number %}} offset); -- {{% lua_type_number %}}
- [ReadDouble](readdouble)({{% lua_type_number %}} offset); -- {{% lua_type_number %}}
- [ReadString](readstring)({{% lua_type_number %}} offset); -- string
- [ReadFixedLengthString](readfixedlengthstring)({{% lua_type_number %}} offset, {{% lua_type_number %}} string_length); -- string
