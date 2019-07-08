## movement
Only FullMap (0x64) and FieldData (0x69) map packets require a position to be written to the network message before the map data.

## stack position
The Tibia client is very, very strict about stack position. Items and creatures have different priorities on a tile which determine where they go in the stack.  
Priority is on a scale of 0-5:  
0 - ground items  
1 - clipping items (e.g., grass borders)  
2 - bottom items (e.g., windows)  
3 - top items (e.g., open doorways)  
4 - creatures  
5 - all other items (e.g., gold coins)  

Let's say you have a tile with a stone ground, open door, and 1 gold coin on it; the stack positions would be as followed:  
0 - stone ground  
1 - open door  
2 - gold coin  
Even though visually, in the client, the gold coin isn't on top of the open door, it has a higher place in the stack because it has a higher priority value. Now let's say a player logs in on this tile, that player would now occupy stack position 2 and the gold coin would be pushed up to stack position 3. Now let's say players can walk on the same tile; if a player were to then step on that tile it would occupy stack position 3, the gold coin would be pushed up to stack position 4, and the first player would stay at stack position 2.

One thing to keep in mind is, the client does this stack position placement on its own. For example, when you send a creature move packet to the client, you don't tell the client the stack position of the tile it's going to, the client determines that, so the server needs to use the same logic because if you try to move a creature at a certain stack position and the client can't find the creature at that stack position (as you've already found out) the client will panic and crash.

The Tibia.dat file, and items.otb if you plan to use that, contains all the information you need to determine an item's priority.

Also, the client will only store 10 objects per tile, so there's no point in sending more than 10 objects per tile to the client.
