# ngin2d

This is an attempt to make a small collision and physics engine in Golang. Mainly for learning purposes. This package is heavily influenced by SolarLune's [resolv](https://github.com/SolarLune/resolv)

## project status

Under development, not working

### concepts
The GameWorld is the main object, which has an array of Cells and a map of GameObjects. Right now everything in the world (and the GameWorld itself) has a Rect object (which represents its position and size in the world). These Rects can be Moved by Vectors and can be Collided agains eachother. A Collision object contains the two colliding GameObject and a possible Resolution Vector which can Resolve the Collision.
