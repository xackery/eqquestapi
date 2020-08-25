---
date: 2020-08-24T16:50:16+02:00
title: Door
menuTitle: Door
weight: 25
---

## Door Methods (Lua)
- door:[CreateDatabaseEntry](createdatabaseentry)(); -- void
- door:[ForceClose](forceclose)(Lua_Mob sender); -- void
- door:[ForceClose](forceclose)(Lua_Mob sender, bool alt_mode); -- void
- door:[ForceOpen](forceopen)(Lua_Mob sender); -- void
- door:[ForceOpen](forceopen)(Lua_Mob sender, bool alt_mode); -- void
- door:[GetDisableTimer](getdisabletimer)(); -- bool
- door:[GetDoorDBID](getdoordbid)(); -- uint32
- door:[GetDoorID](getdoorid)(); -- uint32
- door:[GetHeading](getheading)(); -- float
- door:[GetIncline](getincline)(); -- uint32
- door:[GetKeyItem](getkeyitem)(); -- uint32
- door:[GetLockPick](getlockpick)(); -- uint32
- door:[GetNoKeyring](getnokeyring)(); -- int
- door:[GetOpenType](getopentype)(); -- uint32
- door:[GetSize](getsize)(); -- uint32
- door:[GetX](getx)(); -- float
- door:[GetY](gety)(); -- float
- door:[GetZ](getz)(); -- float
- door:[SetDisableTimer](setdisabletimer)(bool flag); -- void
- door:[SetDoorName](setdoorname)(const char *name); -- void
- door:[SetHeading](setheading)(float h); -- void
- door:[SetIncline](setincline)(uint32 incline); -- void
- door:[SetKeyItem](setkeyitem)(uint32 key); -- void
- door:[SetLocation](setlocation)(float x, float y, float z); -- void
- door:[SetLockPick](setlockpick)(uint32 pick); -- void
- door:[SetNoKeyring](setnokeyring)(int type); -- void
- door:[SetOpenType](setopentype)(uint32 type); -- void
- door:[SetSize](setsize)(uint32 sz); -- void
- door:[SetX](setx)(float x); -- void
- door:[SetY](sety)(float y); -- void
- door:[SetZ](setz)(float z); -- void