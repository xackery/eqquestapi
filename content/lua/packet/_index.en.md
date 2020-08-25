---
date: 2020-08-24T16:50:16+02:00
title: Packet
menuTitle: Packet
weight: 25
---

## Packet Methods (Lua)
- packet:[GetOpcode](getopcode)(); -- int
- packet:[GetRawOpcode](getrawopcode)(); -- int
- packet:[GetSize](getsize)(); -- int
- packet:[Lua_Packet](lua_packet)(const Lua_Packet& o); -- Lua_Packet::Lua_Packet(const
- packet:[ReadDouble](readdouble)(int offset); -- double
- packet:[ReadFixedLengthString](readfixedlengthstring)(int offset, int string_length); -- std::string
- packet:[ReadFloat](readfloat)(int offset); -- float
- packet:[ReadInt16](readint16)(int offset); -- int
- packet:[ReadInt32](readint32)(int offset); -- int
- packet:[ReadInt64](readint64)(int offset); -- int64
- packet:[ReadInt8](readint8)(int offset); -- int
- packet:[ReadString](readstring)(int offset); -- std::string
- packet:[SetOpcode](setopcode)(int op); -- void
- packet:[SetRawOpcode](setrawopcode)(int op); -- void
- packet:[WriteDouble](writedouble)(int offset, double value); -- void
- packet:[WriteFixedLengthString](writefixedlengthstring)(int offset, std::string value, int string_length); -- void
- packet:[WriteFloat](writefloat)(int offset, float value); -- void
- packet:[WriteInt16](writeint16)(int offset, int value); -- void
- packet:[WriteInt32](writeint32)(int offset, int value); -- void
- packet:[WriteInt64](writeint64)(int offset, int64 value); -- void
- packet:[WriteInt8](writeint8)(int offset, int value); -- void
- packet:[WriteString](writestring)(int offset, std::string value); -- void
- packet:[operator=](operator=)(const Lua_Packet& o); -- Lua_Packet&