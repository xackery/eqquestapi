---
title: Packet Class
menuTitle: Packet Class
description: Lua Packet Class
searchTitle: Lua Packet Class
weight: 25
---

## Packet Methods
- [GetSize](getsize)(); -- number
- [GetOpcode](getopcode)(); -- number
- [SetOpcode](setopcode)(number op); -- void
- [GetRawOpcode](getrawopcode)(); -- number
- [SetRawOpcode](setrawopcode)(number op); -- void
- [WriteFloat](writefloat)(number offset, float value); -- void
- [WriteDouble](writedouble)(number offset, double value); -- void
- [WriteString](writestring)(number offset, std::string value); -- void
- [WriteFixedLengthString](writefixedlengthstring)(number offset, std::string value, number string_length); -- void
- [ReadFloat](readfloat)(number offset); -- number
- [ReadDouble](readdouble)(number offset); -- number
- [ReadString](readstring)(number offset); -- string
- [ReadFixedLengthString](readfixedlengthstring)(number offset, number string_length); -- string
