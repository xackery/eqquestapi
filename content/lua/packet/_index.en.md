---
date: 2016-04-09T16:50:16+02:00
title: Packet
menuTitle: Packet
weight: 25
---

# Packet Methods \(Lua\)

- packet:[GetOpcode()](getopcode); -- int
- packet:[GetRawOpcode()](getrawopcode); -- int
- packet:[GetSize()](getsize); -- int
- packet:[Lua_Packet(const Lua_Packet& o)](lua_packet); -- Lua_Packet::Lua_Packet(const
- packet:[ReadDouble(int offset)](readdouble); -- double
- packet:[ReadFixedLengthString(int offset, int string_length)](readfixedlengthstring); -- std::string
- packet:[ReadFloat(int offset)](readfloat); -- float
- packet:[ReadInt16(int offset)](readint16); -- int
- packet:[ReadInt32(int offset)](readint32); -- int
- packet:[ReadInt64(int offset)](readint64); -- int64
- packet:[ReadInt8(int offset)](readint8); -- int
- packet:[ReadString(int offset)](readstring); -- std::string
- packet:[SetOpcode(int op)](setopcode); -- void
- packet:[SetRawOpcode(int op)](setrawopcode); -- void
- packet:[WriteDouble(int offset, double value)](writedouble); -- void
- packet:[WriteFixedLengthString(int offset, std::string value, int string_length)](writefixedlengthstring); -- void
- packet:[WriteFloat(int offset, float value)](writefloat); -- void
- packet:[WriteInt16(int offset, int value)](writeint16); -- void
- packet:[WriteInt32(int offset, int value)](writeint32); -- void
- packet:[WriteInt64(int offset, int64 value)](writeint64); -- void
- packet:[WriteInt8(int offset, int value)](writeint8); -- void
- packet:[WriteString(int offset, std::string value)](writestring); -- void
- packet:[operator=(const Lua_Packet& o)](operator); -- Lua_Packet&