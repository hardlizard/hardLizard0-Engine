# goActionAdventure

goActionAdventure is a prototype for an action adventure game written in golang using hajimehoshi/ebiten as a backend. Once working, the codebase is intended to be reworked into a generic 2d tile engine.

To run, execute `go run *.go` in the source directory.

## Todo

- Entities
  - [x] Player
  - [o] Mob
  - [ ] NPC
  - [ ] Cursor
  - [ ] Projectile
  - [ ] Obj (No AI)
  - [ ] Items
  - [ ] Map
- Components
  - [ ] Player Singleton
  - [x] Position
  - [x] Velocity
  - [x] Acceleration
  - [ ] Z Position
  - [x] Hitbox
  - [ ] Health
  - [ ] Collision Flag
  - [ ] Physics Flag
  - [ ] Inventory
  - [ ] Currency
  - [ ] VRAM
    - [ ] Tileset
    - [ ] Sprite Sheet
- Systems
  - [o] Player Input
    - [x] Movement
    - [ ] Action
    - [ ] Pause Menu
  - [x] Update Position
  - [x] Update Velocity
  - [x] Update Acceleration
  - [ ] Update Z Position
  - [o] AABB Collision Detection
  - [ ] Tile Floor Collision
  - [ ] Tile Wall Collision
  - [ ] Damage
  - [ ] Heal
  - [ ] Knockback
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
