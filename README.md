# goActionAdventure

goActionAdventure is a prototype for an action adventure game written in golang using hajimehoshi/ebiten as a backend. Once working, the codebase is intended to be reworked into a generic 2d tile engine.

To run, execute `go run *.go` in the source directory.

## Todo

* Entities
    - [x] Player
    - [ ] Mob
    - [ ] NPC
    - [ ] Cursor
    - [ ] Projectile
    - [ ] Obj (No AI)
* Components
    - [x] Position
    - [x] Velocity
    - [x] Acceleration
    - [ ] Z Position
    - [ ] Health
    - [ ] Collision Flag
    - [ ] Physics Flag
    - [ ] Free List
* Systems
    - [ ] Update Position
    - [ ] Update Velocity
    - [ ] Update Acceleration
    - [ ] Update Z Position
    - [ ] AABB Collision Detection
    - [ ] Tile Floor Collision
    - [ ] Tile Wall Collision
    - [ ] Damage
    - [ ] Heal
    - [ ] Knockback
#### Alpha
* Map Engine
    - [x] JSON Map Loader
    - [ ] Chunk Renderer
    - [ ] Tile Loading
    - [ ] Use path library to open files
* Camera
    - [ ] Follow Player Sprite
    - [ ] Scroll on screen boundary
    - [ ] Move To New Room/Chunk
    - [x] Render Screen
* Sprites
    - [ ] Actors (PC/NPC)
    - [x] Player Movement
    - [ ] Collisions
    - [ ] Enemy AI
    - [ ] Player Inventory
* Animation Engine
    - [ ] Tiles
    - [x] Sprites
    - [ ] Items
#### Beta
* HUD
    - [ ] Health
    - [ ] Currency
    - [ ] Menus

* Player Character Features 
    - [ ] Health
    - [ ] Item Holding
* Art Assets
    - [ ] Sprites
    - [ ] Tiles 
    - [ ] Maps
#### Release
* Features
    - [ ] Randomizer
    - [ ] Timer
