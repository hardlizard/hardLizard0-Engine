# hardLizard0 Engine

The hardLizard0 engine is a tile based game engine written in golang using hajimehoshi/ebiten as a backend and an Entity-Components-Systems (ECS) framework.

This is a <a href = "https://hardlizard.games"> hardLizard Studios</a> project.
To run, execute `go run *.go` in the source directory.

## Todo

- Entities
  - [x] Player - Being user controllable mob
  - [x] Mob - Being damageable and damaging NPC.
  - [ ] NPC - Being movable, interactable sprite.
  - [ ] Cursor - For use in menus.
  - [ ] Obj (No AI) - Boxes, Chests, Question Blocks, etc.
  - [ ] Items - Usable items that are held in inventory and selected for use.
  - [ ] Map - World Map, Dungeon, etc Maps.
  - [x] Health
  - [x] Death

- Components
  - [ ] Player Singleton - The player singleton will contain one off information for the player entity.  
  - [x] PVA
    - [x] Position
    - [x] Velocity
    - [x] Acceleration
  - [x] Z Position - Game world height, used to determine render order and collision information
  - [x] Hitbox - Circular Hitboxes.
  - [x] Health - Player/Mob/NPC health. Allows negative.
  - [x] Collision Flag - Ture if collidable object.
  - [x] Physics Flag - True if object is movable.
  - [x] Projectile Flag - Moves using AI, damages on collision, despawns on collision.
  - [ ] Inventory - Flags and counters for items. Limit to player or allow NPCs to have inventory depending on game type.
  - [x] Currency - Player only, for now. Possibly all entities depending on game type.
  - [ ] VRAM
    - [ ] Tileset
    - [ ] Sprite Sheet
- Systems
  - [ ] Player Input
    - [x] Movement
    - [ ] Action
    - [ ] Pause Menu
  - [x] Update Position
  - [x] Update Velocity
  - [x] Update Acceleration
  - [ ] Update Z Position
  - [ ] Circle Collision Detection
    - [x] Detection
    - [ ] Push
    - [ ] Knockback
  - [ ] Tile Floor Collision
  - [ ] Tile Wall Collision
  - [x] Damage
  - [x] Heal
  - [x] Death
  - [ ] Render Loop
    - [ ] Render Map (On Grid Single Entity)
    - [x] Render Sprites (Off Grid Entities)
  - [ ] Update Inventory
  - [ ] AI
    - [ ] A*
    - [ ] Dijkstra Map
  - [ ] Animation
    - [ ] Tiles
    - [ ] Sprites
    - [ ] Items
  - [ ] Camera
    - [ ] Follow Player
    - [ ] Scroll on Screen Boundary
    - [ ] Scroll to New Room
- Utility
  - [ ] JSON Map Loader
  - [ ] JSON Sprite Loader
  - [ ] Randomizer
  - [ ] Timer
