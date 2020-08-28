---
title: Packet
menuTitle: Packet
weight: 25
---

## Packet Methods (Lua)
- Packet:[GetSize](getsize)(); -- number
- Packet:[GetOpcode](getopcode)(); -- number
- Packet:[SetOpcode](setopcode)(number op); -- void
- Packet:[GetRawOpcode](getrawopcode)(); -- number
- Packet:[SetRawOpcode](setrawopcode)(number op); -- void
- Packet:[WriteFloat](writefloat)(number offset, float value); -- void
- Packet:[WriteDouble](writedouble)(number offset, double value); -- void
- Packet:[WriteString](writestring)(number offset, std::string value); -- void
- Packet:[WriteFixedLengthString](writefixedlengthstring)(number offset, std::string value, number string_length); -- void
- Packet:[ReadFloat](readfloat)(number offset); -- number
- Packet:[ReadDouble](readdouble)(number offset); -- number
- Packet:[ReadString](readstring)(number offset); -- string
- Packet:[ReadFixedLengthString](readfixedlengthstring)(number offset, number string_length); -- string
