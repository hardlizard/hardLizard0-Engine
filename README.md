# goActionAdventure

goActionAdventure is a prototype for an action adventure game written in golang using hajimehoshi/ebiten as a backend. Once working, the codebase is intended to be reworked into a generic 2D tile engine.

To run, execute `go run *.go` in the source directory.

## Todo

- Entities
  - [x] Player - Being user controllable mob
  - [x] Mob - Being damageable and damaging NPC.
  - [ ] NPC - Being movable, interactable sprite.
  - [ ] Cursor - For use in menus.
  - [ ] Projectile - Moves using AI, damages on collision, despawns on collision.
  - [ ] Obj (No AI) - Boxes, Chests, Question Blocks, etc.
  - [ ] Items - Usable items that are held in inventory and selected for use.
  - [ ] Map - World Map, Dungeon, etc Maps.
  - [x] Health
  - [x] Death

- Components
  - [ ] Player Singleton - The player singleton will contain one off information for the player entity.  
  -[x] PVA
    - [x] Position
    - [x] Velocity
    - [x] Acceleration
  - [ ] Z Position - Game world height, used to determine render order and collision information
  - [x] Hitbox - Currently AABB hitboxes are implemented. Consider adding circles or polygons.
  - [ ] Health - Player/Mob/NPC health. Allows negative.
  - [ ] Collision Flag - Ture if collidable object.
  - [ ] Physics Flag - True if object is movable.
  - [ ] Inventory - Flags and counters for items. Limit to player or allow NPCs to have inventory depending on game type.
  - [ ] Currency - Player only, for now. Possibly all entities depending on game type.
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
  - [ ] Heal
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
