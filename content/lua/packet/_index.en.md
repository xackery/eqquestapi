---
date: 2020-08-24T16:50:16+02:00
title: Packet
menuTitle: Packet
weight: 25
---

## Packet Methods (Lua)
- Packet:[GetSize](getsize)(); -- number
- Packet:[GetOpcode](getopcode)(); -- number
- Packet:[SetOpcode](setopcode)(int op); -- void
- Packet:[GetRawOpcode](getrawopcode)(); -- number
- Packet:[SetRawOpcode](setrawopcode)(int op); -- void
- Packet:[WriteFloat](writefloat)(int offset, float value); -- void
- Packet:[WriteDouble](writedouble)(int offset, double value); -- void
- Packet:[WriteString](writestring)(int offset, std::string value); -- void
- Packet:[WriteFixedLengthString](writefixedlengthstring)(int offset, std::string value, int string_length); -- void
- Packet:[ReadFloat](readfloat)(int offset); -- number
- Packet:[ReadDouble](readdouble)(int offset); -- number
- Packet:[ReadString](readstring)(int offset); -- string
- Packet:[ReadFixedLengthString](readfixedlengthstring)(int offset, int string_length); -- string
