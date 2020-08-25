---
date: 2020-08-24T16:50:16+02:00
title: Object
menuTitle: Object
weight: 25
---

## Object Methods (Lua)
- object:[ClearUser](clearuser)(); -- void
- object:[Close](close)(); -- void
- object:[Delete](delete)(); -- void
- object:[Delete](delete)(bool reset_state); -- void
- object:[DeleteItem](deleteitem)(int index); -- void
- object:[Depop](depop)(); -- void
- object:[EntityVariableExists](entityvariableexists)(const char *name); -- bool
- object:[GetDBID](getdbid)(); -- uint32
- object:[GetHeading](getheading)(); -- float
- object:[GetID](getid)(); -- int
- object:[GetIcon](geticon)(); -- uint32
- object:[GetItemID](getitemid)(); -- uint32
- object:[GetType](gettype)(); -- uint32
- object:[GetX](getx)(); -- float
- object:[GetY](gety)(); -- float
- object:[GetZ](getz)(); -- float
- object:[IsGroundSpawn](isgroundspawn)(); -- bool
- object:[Repop](repop)(); -- void
- object:[Save](save)(); -- bool
- object:[SetEntityVariable](setentityvariable)(const char *name, const char *value); -- void
- object:[SetHeading](setheading)(float h); -- void
- object:[SetID](setid)(int user); -- void
- object:[SetIcon](seticon)(uint32 icon); -- void
- object:[SetItemID](setitemid)(uint32 item_id); -- void
- object:[SetLocation](setlocation)(float x, float y, float z); -- void
- object:[SetModelName](setmodelname)(const char *name); -- void
- object:[SetType](settype)(uint32 type); -- void
- object:[SetX](setx)(float x); -- void
- object:[SetY](sety)(float y); -- void
- object:[SetZ](setz)(float z); -- void
- object:[StartDecay](startdecay)(); -- void
- object:[VarSave](varsave)(); -- uint32